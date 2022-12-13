package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	car := map[int]string{-1: option1, 0: option1, 1: option2}[strings.Compare(option1, option2)]
	return fmt.Sprintf("%s is clearly the better choice.", car)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * .80
	}
	if age < 10 {
		return originalPrice * .70
	}
	return originalPrice * .50
}
