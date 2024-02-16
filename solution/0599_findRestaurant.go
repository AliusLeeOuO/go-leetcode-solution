package solution

import "math"

func FindRestaurant(list1 []string, list2 []string) []string {
	map1 := make(map[string]int)
	for i, v := range list1 {
		map1[v] = i
	}
	// 最小值标记
	lessIndex := math.MaxInt
	var res []string
	for i, v := range list2 {
		if index, ok := map1[v]; ok {
			sum := i + index
			if sum < lessIndex {
				res = make([]string, 0)
				res = []string{v}
				lessIndex = sum
			} else if sum == lessIndex {
				res = append(res, v)
			}
		}
	}

	return res
}
