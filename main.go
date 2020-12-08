package main

import (
	"github.com/zshearin/go-examples/ins"
)

func main() {
	//batch.DoBatch()

	drivers := ins.Contestants{
		{Name: "Zach", Place: 1},
		{Name: "Bob", Place: 2},
		{Name: "Jeff", Place: 3},
		{Name: "Dale", Place: 5},
		{Name: "Carl", Place: 6},
	}

	driver4 := ins.Contestant{Name: "Kevin", Place: 4}

	finalResult := ins.InsertSortOnPlace(drivers, driver4)

	finalResult.Print()

}
