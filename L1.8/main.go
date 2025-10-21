package main

import (
	"errors"
	"fmt"
	"os"
)

func changeBit(num int64, pos int, bit int) (int64, error) {
	if bit != 1 && bit != 0 {
		return 0, errors.New("Bit must be 1 or 0")
	}
	pos = pos - 1
	if pos < 0 || pos > 64 {
		return 0, errors.New("Position must be between 1 and 64")
	}
	mask := (int64(1) << uint(pos))
	if bit == 1 {
		num = num | mask
	} else {
		num = num &^ mask
	}

	return num, nil
}

func main() {
	fmt.Println("Enter number, position (1-63), and bit (0 or 1):")
	var num int64
	var position int
	var bit int
	fmt.Fscan(os.Stdin, &num, &position, &bit)
	value, err := changeBit(num, position, bit)
	if err != nil {
		println(err.Error())
		return
	}
	println(value)
	if position == 64 {
		println("Установка 64-го бита в 1 даёт отрицательное число, т.к. это знаковый бит int64")
	}
}