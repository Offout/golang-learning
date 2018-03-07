package main

func showMeTheType(i interface{}) string {
	switch t := i.(type) {
	case uint:
		return "uint"
	case int8:
		return "int8"
	case float64:
		return "float64"
	case string:
		return "string"
	case byte:
		return "byte"
	case int:
		return "int"
	case int32:
		return "int32"
	case []int:
		return "[]int"
	case map[string]bool:
		return "map[string]bool"
	default:
		_ = t
		return "unknown"
	}
}

func main() {

}
