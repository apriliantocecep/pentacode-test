package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func drawChevronRight(total int, base int) {
	count := total
	countFloor := math.Floor(float64(count / 2))
	countFloorInt := int(countFloor)

	for i := 0; i < countFloorInt; i++ {
		for y := 0; y < i; y++ {
			fmt.Printf(" ")
		}
		for x := 0; x < base; x++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

	for i := 0; i < countFloorInt+base; i++ {
		if i < countFloorInt {
			fmt.Printf(" ")
		} else {
			fmt.Printf("*")
		}
	}
	fmt.Printf("\n")

	for i := countFloorInt; i > 0; i-- {
		for y := 0; y < i-1; y++ {
			fmt.Printf(" ")
		}
		for x := 0; x < base; x++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func drawSilang1() {
	norows := 5

	for i := 0; i < norows; i++ {
		for x := 0; x < norows; x++ {
			if i == x || i+x == norows-1 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
func inputNumber() {
	fmt.Println("Input Number:")
	var number string
	fmt.Scanln(&number)

	numberInt, _ := strconv.Atoi(number)

	totalNumber := 1
	for x := 1; x <= numberInt; x++ {
		totalNumber *= x
	}

	fmt.Printf("Output Number = %d", totalNumber)
}

func reverseString(kata string) string {
	a := strings.Split(kata, " ")
	var b []string
	for _, v := range a {
		var hasil string
		for _, x := range v {
			hasil = string(x) + hasil
		}
		b = append(b, hasil)
	}

	return strings.Join(b, " ")
}

func main() {
	// jawaban 1a
	// drawChevronRight(9, 3)

	// jawaban 1b
	// drawSilang1()

	// jawaban 2a
	// inputNumber()

	// jawban 2b
	// s := reverseString("EDOCATNEP LATIGID")
	// fmt.Println(s)

	// jawaban 3
	// ER Diagram yang akan di buat adalah data logis

	// jawaban 4
	// OOP adalah pemrograman berbasis objek/class, bisa di turunkan/diwariskan, dan mempunyai properti/atribut dan method
}
