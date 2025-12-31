package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
    for _, value := range birdsPerDay {
        sum += value
    }

    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	FirstIndex := (week - 1) * 7
    lastIndex := FirstIndex + 7
    sum := 0
    for _, value := range birdsPerDay[FirstIndex:lastIndex] {
        sum += value
    }

    return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
        if i % 2 == 0 {
            birdsPerDay[i] += 1
        }
    }

    return birdsPerDay
}
