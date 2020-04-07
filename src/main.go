package main

import (
	"errors"
	"fmt"
)

type testStruct struct {
	field1 string
	field2 string
}

var globalNum int

func test_Functoin() string {

	return "test"
}

func testError() (string, error) {
	return "test error", errors.New("test error")
}

func unused() {
	println("unused")
}

func main(){

	//this is test commentary, langauge
	fmt.Sprintf("hello langauge")

	ts := &testStruct{
		field1:"test",
	}



	testMsg, _ := testError()
	println(testMsg)

	println(test_Functoin())
	println(test_Functoin())
	println(ts.field1)

}
