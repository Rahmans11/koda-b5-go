package triangle

import "fmt"

func Segitiga(input int) {
	// for i := 1; i <= input; i++ {
	// 	for j := 1; j < i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("*")
	// } // using for loop logic
	// i := 1
	// for i <= input {
	// 	for j := 1; j < i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("*")
	// 	i++
	// } // using while loop logic
	for i := range input + 1 {
		if i <= 0 {
			// prevent to print 0
		} else {
			for j := 1; j < i; j++ {
				fmt.Print("*")
			}
			fmt.Println("*")
			i++
		}
	} // using range
}
