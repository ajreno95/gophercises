/*
Part 1
Create a program that will read in a quiz provided via a CSV file (more details below)
and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect.
Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

The CSV file should default to problems.csv (example shown below), but the user should be able to customize the filename via a flag.

*/

package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var csvfile string

func init() {
	flag.StringVar(&csvfile, "csvfile", "problems.csv", "set filename for csv quiz file")
	flag.Parse()
}

func main() {
	var correct, wrong int

	file, _ := os.Open(csvfile)
	r := csv.NewReader(bufio.NewReader(file))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("What is " + record[0])

		var answer string
		fmt.Scanln(&answer)

		//troubleshoot this comparison
		if answer == record[1] {
			correct++
		} else {
			wrong++
		}

	}

	fmt.Printf("You got %d correct answers and %d wrong answers", correct, wrong)
}
