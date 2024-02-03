package solution

func CountGoodRectangles(rectangles [][]int) int {
	maxRectHeight := 0
	maxRectCount := 0
	for i := 0; i < len(rectangles); i++ {
		maxRect := min(rectangles[i][0], rectangles[i][1])
		if maxRect == maxRectHeight {
			maxRectCount++
		} else if maxRect > maxRectHeight {
			maxRectHeight = maxRect
			maxRectCount = 1
		}
	}
	return maxRectCount
}
