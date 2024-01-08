package main

import (
	"fmt"
)

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func NewCar(speed int, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

func Drive(car Car) Car {
	car.distance = car.distance + car.speed
	car.battery = car.battery - car.batteryDrain
	return car
}

func CanFinish(car Car, track Track) bool {
	iterations := car.battery / car.batteryDrain
	totalDistance := car.speed * iterations

	if totalDistance >= track.distance {
		return true
	} else {
		return false
	}
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	fmt.Println(car)

	distance := 800
	track := NewTrack(distance)
	fmt.Println(track)

	car = Drive(car)
	fmt.Println(car)

	speed = 5
	batteryDrain = 2
	car = NewCar(speed, batteryDrain)

	distance = 100
	track = NewTrack(distance)

	fmt.Println(CanFinish(car, track))
}
