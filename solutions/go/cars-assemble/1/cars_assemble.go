package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	i := successRate / float64(100)
    return float64(productionRate) * i
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	perHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    return int(perHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tenth := int(carsCount / 10)
    remain := carsCount - (tenth * 10)

    return uint((tenth * 95000) + (remain * 10000))
}
