package conf

import (
	"github.com/coolestowl/go-base/json"
	"github.com/fsnotify/fsnotify"
	"github.com/jeremywohl/flatten"
	"github.com/spf13/viper"
)

type Hook func()

type Watcher interface {
	RegisterHook(string, Hook)
	SetFileName(string)
	Watch() error
}

func NewWatcher() Watcher {
	v := viper.New()
	return &viperImpl{
		v:     v,
		hooks: make(map[string][]Hook),
	}
}

type viperImpl struct {
	v     *viper.Viper
	hooks map[string][]Hook
	bak   map[string]interface{}
}

func (v *viperImpl) RegisterHook(key string, f Hook) {
	if _, ok := v.hooks[key]; !ok {
		v.hooks[key] = make([]Hook, 0, 1)
	}
	v.hooks[key] = append(v.hooks[key], f)
}

func (v *viperImpl) SetFileName(name string) {
	v.v.SetConfigFile(name)
}

func (v *viperImpl) Watch() error {
	if err := v.v.ReadInConfig(); err != nil {
		return err
	}

	mp, err := flatten.Flatten(v.v.AllSettings(), "", flatten.DotStyle)
	if err != nil {
		return err
	}
	v.bak = mp

	v.v.OnConfigChange(func(in fsnotify.Event) {
		mp, err := flatten.Flatten(v.v.AllSettings(), "", flatten.DotStyle)
		if err != nil {
			// FIXME: err handle
			panic(err)
		}

		changed, err := json.DiffMap(v.bak, mp)
		if err != nil {
			panic(err)
		}

		v.bak = mp
		if hooks, ok := v.hooks[changed]; ok {
			for _, f := range hooks {
				f()
			}
		}
	})

	v.v.WatchConfig()
	return nil
}
