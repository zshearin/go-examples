package ifc

import "fmt"

func DoExample() {
	var code interface{}
	code = "started as an interface but am going to cast to a string"

	myString := code.(string)
	fmt.Println(myString)
}
