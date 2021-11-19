package main

// Program conditional logic

import (
	"fmt"
)

func main() {

	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "Less than zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)

	if x := -42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)

}

// // Switch Statements

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	rand.Seed(time.Now().Unix())

// 	// fmt.Println("Day", dow)

// 	var result string
// 	switch dow := rand.Intn(7) + 1; dow {
// 	case 1:
// 		result = "It's Sunday!"
// 		// fallthrough
// 	case 2:
// 		result = "It's Monday!"
// 		// fallthrough
// 	default:
// 		result = "It's some other day!"
// 	}
// 	fmt.Println(result)
// }

// // Create loops - for statements

// import (
// 	"fmt"
// )

// func main() {
// 	colors := []string{"Red", "Green", "Blue"}
// 	fmt.Println(colors)

// 	for i := 0; i < len(colors); i++ {
// 		fmt.Println(colors[i])
// 	}

// 	for i := range colors {
// 		fmt.Println(colors[i])
// 	}

// 	for _, color := range colors {
// 		fmt.Println(color)
// 	}

// 	value := 1
// 	for value < 10 {
// 		fmt.Println("Value:", value)
// 		value++
// 	}

// 	sum := 1
// 	for sum < 1000 {
// 		sum += sum
// 		fmt.Println("Sum:", sum)
// 		if sum > 200 {
// 			goto theEnd
// 		}
// 	}

// theEnd:
// 	fmt.Println("End of program")
// }
