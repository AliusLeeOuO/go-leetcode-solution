package solution

func MaximumPopulation(logs [][]int) int {
	if len(logs) == 0 {
		return 0
	}
	minLog := 0
	maxLog := 0
	for i := 0; i < len(logs); i++ {
		for j := 0; j < len(logs[i]); j++ {
			if logs[i][j] < minLog || minLog == 0 {
				minLog = logs[i][j]
			}
			if logs[i][j] > maxLog {
				maxLog = logs[i][j]
			}
		}
	}
	maxCount := 0
	maxYear := 0
	for i := minLog; i <= maxLog; i++ {
		currentCount := 0
		for j := 0; j < len(logs); j++ {
			if i >= logs[j][0] && i < logs[j][1] {
				currentCount++
			}
		}
		if currentCount > maxCount {
			maxCount = currentCount
			maxYear = i
		}
	}
	return maxYear
}
