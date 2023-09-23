package t

func TypeT(v any) {
	switch v := v.(type) {
	case int:
		println("int value: %v", v)
	case float64:
		println("float64 value: %v", v)
	case string:
		println("string value: %v", v)
	default:
		println("none of them")
	}
}
