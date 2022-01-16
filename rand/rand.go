package rand

import (
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/spf13/cast"
)

type Rand struct {
	inner *rand.Rand

	strSeed    string
	candidates []interface{}
}

func New() *Rand {
	defaultSeed := time.Now().Unix()

	return &Rand{
		inner:   rand.New(rand.NewSource(defaultSeed)),
		strSeed: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
	}
}

func (r *Rand) Seed(seed int64) {
	r.inner = rand.New(rand.NewSource(seed))
}

func (r *Rand) Candidates(args ...interface{}) *Rand {
	r.candidates = args
	return r
}

func (r *Rand) RandCandidate() interface{} {
	rdx := r.inner.Int()

	if r.candidates != nil && len(r.candidates) > 0 {
		idx := rdx % len(r.candidates)
		return r.candidates[idx]
	}

	return rdx
}

func (r *Rand) StrCandidate(str string) *Rand {
	r.strSeed = str
	return r
}

func (r *Rand) RandStr(length int) string {
	var (
		seedLength    = len(r.strSeed)
		resultBuilder = new(strings.Builder)
	)

	for range make([]struct{}, length) {
		idx := r.inner.Int() % seedLength
		resultBuilder.Write([]byte{r.strSeed[idx]})
	}

	return resultBuilder.String()
}

func (r *Rand) RangeInt(min, max int) int {
	length := max - min
	if length < 1 {
		return 0
	}

	return min + r.inner.Int()%length
}

func (r *Rand) RandStruct(ptr interface{}) {
	structType := reflect.TypeOf(ptr).Elem()

	structValue := reflect.ValueOf(ptr).Elem()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		tag := field.Tag.Get("rand")

		seps := strings.Split(tag, ":")
		if len(seps) != 2 {
			continue
		}

		switch seps[0] {
		case "int":
			rangeStrs := strings.Split(seps[1], ",")
			min, max := 0, 0
			if len(rangeStrs) == 1 {
				max = cast.ToInt(rangeStrs[0])
			} else if len(rangeStrs) == 2 {
				min = cast.ToInt(rangeStrs[0])
				max = cast.ToInt(rangeStrs[1])
			}

			structValue.Field(i).SetInt(int64(r.RangeInt(min, max)))
		case "str":
			length := cast.ToInt(seps[1])
			structValue.Field(i).SetString(r.RandStr(length))
		case "candidate":
			strs := strings.Split(seps[1], ",")

			candidates := make([]interface{}, 0, len(strs))
			for _, str := range strs {
				candidates = append(candidates, str)
			}

			randOne := r.Candidates(candidates...).RandCandidate()

			switch field.Type.Kind() {
			case reflect.Int:
				structValue.Field(i).SetInt(cast.ToInt64(randOne))
			case reflect.String:
				structValue.Field(i).SetString(cast.ToString(randOne))
			}
		default:
			panic(seps[0])
		}
	}
}
