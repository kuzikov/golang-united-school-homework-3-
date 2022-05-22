package homework

import "sort"

// MapEntry ...
type MapEntry struct {
	Key int
	Val string
}

// SortByKey ...
type SortByKey []MapEntry

func (a SortByKey) Len() int           { return len(a) }
func (a SortByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

func sortMapValues(input map[int]string) (result []string) {
	tmp := make([]MapEntry, 0, len(input)) // []struct{ {},{},{} }
	for key, val := range input {
		tmp = append(tmp, MapEntry{key, val})
	}

	sort.Sort(SortByKey(tmp))
	result = make([]string, 0, len(input))

	for _, v := range tmp {
		result = append(result, v.Val)
	}
	return
}

// Input -> {2: "a", 0: "b", 1: "c"}
// Output -> ["b", "c", "a"]
// Input -> {10: "aa", 0: "bb", 500: "cc"}
// Output -> ["bb", "aa", "cc"]
