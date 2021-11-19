package main

// Reference values with pointers

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("Value of p:", *p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Pointer 1:", *pointer1)
	fmt.Println("Value 1:", value1)

}

// // Store Ordered Value in Arrays

// import (
// 	"fmt"
// )

// func main() {
// 	var colors [3]string
// 	colors[0] = "Red"
// 	colors[1] = "Green"
// 	colors[2] = "Blue"
// 	fmt.Println(colors)
// 	fmt.Println(colors[0])

// 	var numbers = [5]int{5, 3, 1, 2, 4}
// 	fmt.Println(numbers)

// 	fmt.Println("Number of colors:", len(colors))
// 	fmt.Println("Number of numbers:", len(numbers))
// }

// // Manage ordered values in slices

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	var colors = []string{"Red", "Green", "Blue"}
// 	fmt.Println(colors)
// 	colors = append(colors, "Purple")
// 	fmt.Println(colors)

// 	colors = append(colors[1:len(colors)])
// 	fmt.Println(colors)

// 	colors = append(colors[:len(colors)-1])
// 	fmt.Println(colors)

// 	numbers := make([]int, 5)
// 	numbers[0] = 134
// 	numbers[1] = 72
// 	numbers[2] = 32
// 	numbers[3] = 12
// 	numbers[4] = 156
// 	fmt.Println(numbers)

// 	numbers = append(numbers, 235)
// 	fmt.Println(numbers)

// 	sort.Ints(numbers)
// 	fmt.Println(numbers)
// }

// // Store unordered values in maps

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	states := make(map[string]string)
// 	fmt.Println(states)
// 	states["WA"] = "Washington"
// 	states["OR"] = "Oregon"
// 	states["CA"] = "California"
// 	fmt.Println(states)

// 	california := states["CA"]
// 	fmt.Println(california)

// 	delete(states, "OR")
// 	states["NY"] = "New York"
// 	fmt.Println(states)

// 	for k, v := range states {
// 		fmt.Printf("%v: %v\n", k, v)
// 	}

// 	keys := make([]string, len(states))
// 	i := 0
// 	for k := range states {
// 		keys[i] = k
// 		i++
// 	}
// 	fmt.Println(keys)
// 	sort.Strings(keys)
// 	fmt.Println(keys)

// 	for i := range keys {
// 		fmt.Println(states[keys[i]])
// 	}

// }

// // Group related values in structs

// import (
// 	"fmt"
// )

// func main() {
// 	poodle := Dog{"Poodle", 10}
// 	fmt.Println(poodle)
// 	fmt.Printf("%+v\n", poodle)
// 	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
// 	poodle.Weight = 9
// 	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
// }

// // Dog is a struct
// type Dog struct {
// 	Breed  string
// 	Weight int
// }
