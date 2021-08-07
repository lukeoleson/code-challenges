package main

import (
	"fmt"
	"math"
)

func main() {
	invertBits()
}

func invertBits() {
	var input uint32 = 43261596
	fmt.Println("Have:   43261596 -> 00000010100101000001111010011100")
	fmt.Println("Want: 4251705699 -> 11111101011010111110000101100011")

	var num uint32
	for i := 0; i < 32; i++ {

		rightMostBit := input & 1 // get the right-most bit
		input = input >> 1 // shave 1 bit off the right

		if i == 0 {
			num = invertBit(rightMostBit)
		} else {
			if invertBit(rightMostBit) == 1 {
				num += uint32(math.Pow(2, float64(i)))
			}
		}
	}

	fmt.Printf("Got : %v -> %b\n", num, num)
}

func invertBit(b uint32) uint32 {
	if b == 1 {
		return 0
	}
	return 1
}
