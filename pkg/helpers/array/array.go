package array

func IntersectUint64(slice1 []uint64, slice2 []uint64) []uint64 {
	var intersect []uint64
	for _, element1 := range slice1 {
		for _, element2 := range slice2 {
			if element1 == element2 {
				intersect = append(intersect, element1)
			}
		}
	}
	return intersect //return slice after intersection
}
