package main

import (
	"fmt"
)

type Robot struct {
	coordinates     []int
	directionFacing string
}

func MakeRobot(coordinates []int, direction string) Robot {
	return Robot{
		coordinates:     coordinates,
		directionFacing: direction,
	}
}

func TurnRight(directionFacing string) string {
	if directionFacing == "North" {
		return "East"
	}
	if directionFacing == "East" {
		return "South"
	}
	if directionFacing == "South" {
		return "West"
	}
	if directionFacing == "West" {
		return "North"
	}
	return directionFacing
}

func TurnLeft(directionFacing string) string {
	if directionFacing == "North" {
		return "West"
	}
	if directionFacing == "East" {
		return "North"
	}
	if directionFacing == "South" {
		return "East"
	}
	if directionFacing == "West" {
		return "South"
	}
	return directionFacing
}

func Advance(currentPosition []int, directionFacing string) []int {
	if directionFacing == "North" {
		return []int{currentPosition[0], currentPosition[1] + 1}
	}
	if directionFacing == "East" {
		return []int{currentPosition[0] + 1, currentPosition[1]}
	}
	if directionFacing == "South" {
		return []int{currentPosition[0], currentPosition[1] - 1}
	}
	if directionFacing == "West" {
		return []int{currentPosition[0] - 1, currentPosition[1]}
	}
	return currentPosition
}

func InterpretInstructions(robot Robot, instructions string) Robot {
	for _, letter := range instructions {
		if string(letter) == "R" {
			robot.directionFacing = TurnRight(robot.directionFacing)
		}
		if string(letter) == "L" {
			robot.directionFacing = TurnLeft(robot.directionFacing)
		}
		if string(letter) == "A" {
			robot.coordinates = Advance(robot.coordinates, robot.directionFacing)
		}
	}
	return robot
}

func main() {
	robot := MakeRobot([]int{7, 3}, "North")
	instructions := "RAALAL"

	fmt.Println("Starting location", robot.coordinates, "facing", robot.directionFacing)
	robot = InterpretInstructions(robot, instructions)
	fmt.Println("Final location", robot.coordinates, "facing", robot.directionFacing)
}
