package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	ans := 0

	for _,i := range birdsPerDay{
		ans+=i
	}

	return ans
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	
	ans :=0
	start:=(week-1)  * 7

	for i:=0;i<7;i++{
		ans+=birdsPerDay[start+i]
	}

	return ans
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	
	for i:=0;i<len(birdsPerDay);i=i+2{
		birdsPerDay[i]++
	}

	return birdsPerDay
}
