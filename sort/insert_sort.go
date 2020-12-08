package sort

import (
	"fmt"
	"sort"
	"strconv"
)

type contestants []contestant
type contestant struct {
	Name  string
	Place int
}

//DoExample runs example
func DoExample() {
	drivers := contestants{
		{Name: "Zach", Place: 1},
		{Name: "Jack", Place: 2},
		{Name: "Jeff", Place: 3},
		{Name: "Dale", Place: 5},
		{Name: "Carl", Place: 6},
	}
	driver4 := contestant{Name: "Kevin", Place: 4}

	fmt.Printf("Insert sort example\n=====================\n")
	fmt.Println("Initial list:")
	drivers.print()
	fmt.Printf("\nAdding ")
	driver4.print()
	finalResult := insertSortOnPlace(drivers, driver4)

	fmt.Println("\n\nFinal result: ")
	finalResult.print()
}

func insertSortOnPlace(cl contestants, c contestant) contestants {

	curPlace := c.Place

	indexToInsert := sort.Search(len(cl), func(i int) bool {
		return cl[i].Place > curPlace
	})

	//append empty object
	cl = append(cl, contestant{})

	//shift entries one
	copy(cl[(indexToInsert+1):], cl[indexToInsert:])

	//set right index to the object to insert
	cl[indexToInsert] = c

	return cl

}

func (c contestants) print() {
	for _, driver := range c {
		driver.print()
		fmt.Printf("\n")
	}

}

func (c contestant) print() {
	fmt.Printf("Name: " + c.Name + ", Place: " + strconv.Itoa(c.Place))
}
