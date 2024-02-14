package solution

func SupplyWagon(supplies []int) []int {
	if len(supplies) == 0 {
		return []int{}
	} else if len(supplies) == 1 {
		return supplies
	}
	var flag = len(supplies) / 2
	for len(supplies) != flag {
		minSum := supplies[0] + supplies[1]
		minIndex := 0
		// 找到物资之和最小的两辆相邻马车
		for i := 0; i < len(supplies)-1; i++ {
			if sum := supplies[i] + supplies[i+1]; sum < minSum {
				minSum = supplies[i] + supplies[i+1]
				minIndex = i
			}
		}
		// 合并马车
		supplies[minIndex] = minSum
		// 删除后面一辆马车
		supplies = append(supplies[:minIndex+1], supplies[minIndex+2:]...)
	}
	return supplies
}
