package slicekit

func StringIncludes(records []string, obj string) bool {
	for _, v := range records {
		if v == obj {
			return true
		}
	}

	return false
}
