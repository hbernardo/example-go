package example

import "github.com/hbernardo/example-go/awesome"

var result string

func Setup() {

	// Comment

	result = awesome.Smile()

}

func GetResult() string {

	/*
	   Comment
	*/

	return result

}
