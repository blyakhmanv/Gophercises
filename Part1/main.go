package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type problem struct {
	Question string
	Answer   string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	//parce csv file and store quiz in the problems slice
	csvFile, error := os.Open(*csvFilename)
	if error != nil {
		log.Fatal(error)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var problems []problem
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		problems = append(problems, problem{
			Question: line[0],
			Answer:   line[1],
		})
	}

	//start quiz
	consolereader := bufio.NewReader(os.Stdin)
	var score int
	fmt.Println("Game started")
	for _, problem := range problems {
		fmt.Printf(problem.Question + "=")
		answer, _ := consolereader.ReadString('\n')
		answer = strings.Replace(answer, "\r\n", "", -1)
		if strings.Compare(problem.Answer, answer) == 0 {
			score++
		}
	}
	fmt.Printf("Game Over. Your score is %d out of %d.\n", score, len(problems))
}
