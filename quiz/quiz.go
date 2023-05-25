package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := 0
	if len(os.Args) > 1 {
		seconds, _ = strconv.Atoi(os.Args[1])
	} else {
		seconds = 30
	}
	correctAnswerCh := make(chan bool)
	endQuizCh := make(chan bool)
	timerCh := make(chan bool)
	// read problems.csv and parse it with csv package
	csvFile, err := os.Open("problems.csv")
	defer csvFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Press enter to start the quiz")
	_ = getInputAndCleanIt()
	go startTimer(seconds, timerCh)
	go askQuestions(lines, correctAnswerCh, endQuizCh)
	totalCorrect := 0
out:
	for {
		select {
		case <-correctAnswerCh:
			totalCorrect++
		case <-endQuizCh:
			break out
		case <-timerCh:
			break out
		}
	}
	fmt.Printf("total correct: %d/%d", totalCorrect, len(lines))
}

func askQuestions(lines [][]string, correctAnswerCh chan<- bool, endQuizCh chan<- bool) {
	totalCorrect := 0
	for _, line := range lines {
		question := line[0]
		answer := line[1]
		fmt.Println("question: ", question)
		userAnswer := getInputAndCleanIt()
		if userAnswer == answer {
			totalCorrect++
			correctAnswerCh <- true
		}
	}
	endQuizCh <- true
}

func getInputAndCleanIt() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func startTimer(seconds int, timerCh chan<- bool) {
	time.Sleep(time.Duration(seconds) * time.Second)
	timerCh <- true
}
