package json

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jeremywohl/flatten"
)

func DiffString(this, other string) (string, error) {
	var thisObj interface{}
	if err := Unmarshal([]byte(this), &thisObj); err != nil {
		return "", err
	}

	var otherObj interface{}
	if err := Unmarshal([]byte(other), &otherObj); err != nil {
		return "", err
	}

	return Diff(thisObj, otherObj)
}

func Diff(this, other interface{}) (string, error) {
	thisMap, ok := this.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("bad params %v", this)
	}

	otherMap, ok := other.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("bad params %v", other)
	}

	return DiffMap(thisMap, otherMap)
}

func DiffMap(this, other map[string]interface{}) (string, error) {
	thisFlatten, err := flatten.Flatten(this, "", flatten.DotStyle)
	if err != nil {
		return "", err
	}

	otherFlatten, err := flatten.Flatten(other, "", flatten.DotStyle)
	if err != nil {
		return "", err
	}

	return diff(thisFlatten, otherFlatten), nil
}

func diff(this, other map[string]interface{}) string {
	keyMap := make(map[string]struct{})
	for k := range this {
		keyMap[k] = struct{}{}
	}
	for k := range other {
		keyMap[k] = struct{}{}
	}

	changedKey := make([]string, 0, 1)
	for k := range keyMap {
		thisVal, thisOk := this[k]
		otherVal, otherOk := other[k]

		if !(thisOk == otherOk && reflect.DeepEqual(thisVal, otherVal)) {
			changedKey = append(changedKey, k)
		}
	}

	commonPrefix := longestCommonPrefix(changedKey)
	if len(commonPrefix) == 0 {
		return "."
	}
	if strings.HasSuffix(commonPrefix, ".") {
		return strings.TrimSuffix(commonPrefix, ".")
	}
	return commonPrefix
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	result := ""
	for _, str := range strs {
		if len(str) == 0 {
			return ""
		}

		result = strs[0]
		for !strings.HasPrefix(str, result) {
			result = result[:len(result)-1]
		}
		if len(result) == 0 {
			return ""
		}
	}
	return result
}
