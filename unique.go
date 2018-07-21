package slicekit

func IntUniq(records []int) []int {
	results := []int{}
	flags := make(map[int]bool)

	for _, record := range records {
		if flags[record] {
			continue
		}

		flags[record] = true
		results = append(results, record)
	}

	return results
}

func StringUniq(records []string) []string {
	results := []string{}
	flags := make(map[string]bool)

	for _, record := range records {
		if flags[record] {
			continue
		}

		flags[record] = true
		results = append(results, record)
	}

	return results
}
