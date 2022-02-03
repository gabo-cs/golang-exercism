package cars

// IndividualCarProductionCost stores how much each car normally costs to produce
const IndividualCarProductionCost = 10_000

// GroupedCarsForBetterProduction stores the number of cars that can
// be produced together at a lower cost.
const GroupedCarsForBetterProduction = 10

// GroupedCarsProductionCost stores how much it costs to produce 10 cars
// together with a bit of planning.
const GroupedCarsProductionCost = 95_000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	groupedCars := carsCount / GroupedCarsForBetterProduction
	remainingCars := carsCount % GroupedCarsForBetterProduction
	return uint((groupedCars * GroupedCarsProductionCost) + (remainingCars * IndividualCarProductionCost))
}
