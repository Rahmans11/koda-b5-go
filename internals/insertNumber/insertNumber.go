package insertnumber

import "fmt"

func InsertNumber(input int) {
	var sliceInt = []int{50, 75, 66, 20, 32, 90}
	firstSlice := sliceInt[0:3]
	secondSlice := sliceInt[3:6]

	firstPartioning := make([]int, len(firstSlice))
	copy(firstPartioning, firstSlice)

	secondPartioning := make([]int, len(secondSlice))
	copy(secondPartioning, secondSlice)

	firstPartioning = append(firstPartioning, input)
	firstPartioning = append(firstPartioning, secondPartioning...)
	for i, v := range firstPartioning {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}
