package builder

import (
	"bytes"
	"strconv"
)

func interfaceToString(input_int []interface{}) []string {
	length := len(input_int)
	columns := make([]string, length)
	for index, element := range input_int {
		columns[index] = convertValueToString(element)
	}
	return columns
}

func convertValueToString(value interface{}) string {
	var result string
	switch value := value.(type) {
	case int:
		result = strconv.Itoa(int(value))
	case float64:
		result = strconv.FormatFloat(value, 'f', 6, 64)
	case string:
		result = value
	case []int:
		var res bytes.Buffer
		for index, el := range value {
			res.WriteString(strconv.Itoa(int(el)))
			if index != len(value)-1 {
				res.WriteString(", ")
			}
		}
		result = res.String()
	case []float64:
		var res bytes.Buffer
		for index, el := range value {
			res.WriteString(strconv.FormatFloat(el, 'f', 6, 64))
			if index != len(value)-1 {
				res.WriteString(", ")
			}
		}
		result = res.String()

	}
	return result
}
