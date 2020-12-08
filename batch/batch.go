package batch

import "fmt"

//DoBatch is a function skeleton for handling batch requests
func DoBatch() {
	slice := make([]int, 159)

	// Split the slice into batches of 20 items.
	batch := 20

	for i := 0; i < len(slice); i += batch {
		j := i + batch
		if j > len(slice) {
			j = len(slice)
		}

		fmt.Println(slice[i:j]) // Process the batch.
	}

}
