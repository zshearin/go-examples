package string

import "fmt"

//DoExample runs example
func DoExample() {

	//Strings are immutable and behave like read-only byte slices (with a few extra properties)

	//To update the data, use a rune slice instead:

	myString := "I want this"

	buf := []rune(myString)

	buf[9] = 'a'
	buf[10] = 't'

	s := string(buf)

	fmt.Println(s)

}
