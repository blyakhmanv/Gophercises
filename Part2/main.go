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
	"time"
)

type problem struct {
	Question string
	Answer   string
}

var consolereader = bufio.NewReader(os.Stdin)

func userInput(didInput chan<- string) {
	var answer string
	answer, _ = consolereader.ReadString('\n')
	didInput <- answer
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	secTimer := flag.Int("sec", 30, "number of seconds for the quiz")
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
	var score int
	fmt.Printf("Press Enter to start...")
	consolereader.ReadString('\n')
	fmt.Println("Game started")

	//set and run timer
	timer := time.NewTimer(time.Duration(*secTimer) * time.Second)

BreakTimer:
	for _, problem := range problems {
		didInput := make(chan string, 1)
		fmt.Printf(problem.Question + "=")
		go userInput(didInput)
		select {
		case <-timer.C:
			fmt.Println()
			fmt.Println("Time is out")
			break BreakTimer
		case answer := <-didInput:
			answer = strings.Replace(answer, "\r\n", "", -1)
			if strings.Compare(problem.Answer, answer) == 0 {
				score++
			}

		}
	}
	fmt.Println()
	fmt.Printf("Game Over. Your score is %d out of %d.\n", score, len(problems))
}
