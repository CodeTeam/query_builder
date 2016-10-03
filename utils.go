package builder

import (
	"bytes"
	"strconv"
)

func interfaceToString(input_int []interface{}) []string {
	length := len(input_int)
	columns := make([]string, length)
	for index, element := range input_int {
		columns[index] = element.(string)
	}
	return columns
}

func convertValueToString(value interface{}) string {
	var result string
	switch value := value.(type) {
	case int:
		result = strconv.Itoa(int(value)) //strings.Replace(expr, "?", strconv.Itoa(int(value)), -1)
	case string:
		result = value //strings.Replace(expr, "?", value, -1)
	case []int:
		var res bytes.Buffer
		for index, el := range value {
			res.WriteString(strconv.Itoa(int(el)))
			if index != len(value)-1 {
				res.WriteString(", ")
			}
		}
		result = res.String() //strings.Replace(expr, "?", res.String(), -1)

	}
	return result
}
