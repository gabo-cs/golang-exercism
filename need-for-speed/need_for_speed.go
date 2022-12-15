// Package speed simulates races between various types of remote controlled cars
package speed

// FullBattery stores the maximum battery percentage cars can have.
const FullBattery int = 100

// A Car represents a remote controlled car
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// A Track represents a race track
type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      FullBattery,
		batteryDrain: batteryDrain,
		speed:        speed,
	}
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	driveTimes := car.battery / car.batteryDrain
	return (car.speed * driveTimes) >= track.distance
}
