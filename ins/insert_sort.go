package ins

import (
	"fmt"
	"sort"
	"strconv"
)

//Contestants is a list of Contestant objects
type Contestants []Contestant

//Contestant is basic struct for demonstrating insert sort
type Contestant struct {
	Name  string
	Place int
}

//InsertSortOnPlace assumes input is already in order and inserts passed in value to existing list of values
func InsertSortOnPlace(contestants Contestants, contestant Contestant) Contestants {

	curPlace := contestant.Place

	indexToInsert := sort.Search(len(contestants), func(i int) bool {
		return contestants[i].Place > curPlace
	})

	//append empty object
	contestants = append(contestants, Contestant{})

	//shift entries one
	copy(contestants[(indexToInsert+1):], contestants[indexToInsert:])

	//set right index to the object to insert
	contestants[indexToInsert] = contestant

	return contestants

}

//Print prints Contestants
func (c Contestants) Print() {
	for _, driver := range c {
		driver.Print()
	}

}

//Print prints a Contestant
func (c Contestant) Print() {
	fmt.Printf("Name: " + c.Name + ", Place: " + strconv.Itoa(c.Place) + "\n")
}
