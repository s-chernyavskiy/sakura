package util

func ToString(v any) string {
	if s, ok := v.(string); ok {
		return s
	}

	return ""
}
