// Program takes up to 8 binary digits and convert them to dec
// by Strawfervor
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Please enter binary number:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	//fmt.Printf("len of input: %d\n", len(input)) //debug
	var result float64 = 0
	powerBase := 0.0

	for i := (len(input) - 3); i >= 0; i-- {
		number, _ := strconv.Atoi(string(input[i]))
		if number != 1 && number != 0 {
			fmt.Println("It's not binary number")
			result = 0
			break
		}
		result += float64(number) * math.Pow(2, powerBase)
		powerBase++
	}

	fmt.Println(result)
}
