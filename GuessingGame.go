// Classic guess a number game, every programmer should write
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Guess a number from 0 to 100, you have 5 tries:")
	answear := rand.Intn(100)
	var playerInput int
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &playerInput)

		if playerInput == answear {
			fmt.Println("Correct, You win!")
			break
		} else if playerInput == -1 { //this is because scanf doesn't consume all line :c :c :c
			i--
		} else if answear < playerInput {
			fmt.Printf("Number is lower than %d\n", playerInput)
		} else if answear > playerInput {
			fmt.Printf("Number is greater than %d\n", playerInput)
		}
		playerInput = -1
	}

}
