//intializing go file

package main

//docker                run image <cmd> <params>
//go run main.go run              <cmd> <params>
import(

	"os"
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

		cmd := exec.Command(os.Args[2:],os.Args[3:]...)

		cmd.Stdin = os.Stdin

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()

}

//creating  a function must
func must(err error) {

	if err != nil {

		panic(err)

	}

}
