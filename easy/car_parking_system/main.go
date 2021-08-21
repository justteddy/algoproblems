package main

import "fmt"

func main() {
	ps := Constructor(1, 1, 0)
	fmt.Println(ps.AddCar(1))
	fmt.Println(ps.AddCar(1))
}

type ParkingSystem struct {
	big, medium, small int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if ps.big-1 < 0 {
			return false
		}
		ps.big--
	case 2:
		if ps.medium-1 < 0 {
			return false
		}
		ps.medium--
	case 3:
		if ps.small-1 < 0 {
			return false
		}
		ps.small--
	default:
		return false
	}

	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
