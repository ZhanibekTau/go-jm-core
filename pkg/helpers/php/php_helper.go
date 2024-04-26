package php

import "fmt"
import "github.com/leeqvip/gophp"

func UnSerializePhpString(serializedString string) map[string]interface{} {
	resultMap := make(map[string]interface{})

	var dat = []byte(serializedString)
	out, err := gophp.Unserialize(dat)

	if err != nil {
		fmt.Println(err)

		return resultMap
	}

	if m, ok := out.(map[string]interface{}); ok {
		resultMap = m

		return resultMap
	}

	return resultMap
}
