//intializing go file

package main

//docker                run image <cmd> <params>
//go run main.go run              <cmd> <params>

//import os
import (
	"os"
//import fmt
 "fmt"
)
//creating a new function main
func main() {

	switch os.Args[1] {

	case "run":

		run()

		default:

			panic("bad command")

		}

	}

//creating a new function run
func run() {

	fmt.Printf("Running %v \n", os.Args[2:])
}

//creating  a function must
func must(err error) {

	if err != nil {

		panic(err)

	}

}
