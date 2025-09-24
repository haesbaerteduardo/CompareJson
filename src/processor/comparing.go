package processor

type ComparisonResult struct {
	InBothFiles []string
	OnlyInFile1 []string
	OnlyInFile2 []string
	File1Count  int
	File2Count  int
}

func CompareIDs(ids1, ids2 map[string]bool) ComparisonResult {
	var inBoth, onlyIn1, onlyIn2 []string

	for id := range ids1 {
		if ids2[id] {
			inBoth = append(inBoth, id)
		} else {
			onlyIn1 = append(onlyIn1, id)
		}
	}

	for id := range ids2 {
		if !ids1[id] {
			onlyIn2 = append(onlyIn2, id)
		}
	}

	return ComparisonResult{
		InBothFiles: inBoth,
		OnlyInFile1: onlyIn1, OnlyInFile2: onlyIn2,
		File1Count: len(ids1),
		File2Count: len(ids2),
	}
}
