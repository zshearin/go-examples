package basic

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//DoExample is a function skeleton for handling batch requests
func DoExample() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	filename := path + "/file.txt"
	fmt.Println(path)
	//takes all file contents into memory
	data, err := ioutil.ReadFile(filename)
	check(err)
	fmt.Print(string(data))

	//for more control over how and what parts of a file are read:
	//f, err := os.Open("file.txt")

}
