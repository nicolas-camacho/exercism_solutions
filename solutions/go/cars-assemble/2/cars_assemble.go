package cars

// HourInMinutes is the number of minutes in an hour.
const HourInMinutes float64 = 60

// ProductionCost is the cost of producing a car.
const ProductionCost uint = 10000

// ProductionCostPerTenCars is the cost of producing 10 cars.
const ProductionCostPerTenCars uint = 95000

const CarBatchSize uint = 10

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) / 100) * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / HourInMinutes)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var remainingCars uint = uint(carsCount) % CarBatchSize
	var carBatches uint = uint(carsCount) / CarBatchSize
	return carBatches*ProductionCostPerTenCars + remainingCars*ProductionCost
}
