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
		seconds = 3
	}
	correct_answer_ch := make(chan bool)
	end_quiz_ch := make(chan bool)
	timer_ch := make(chan bool)
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
	go startTimer(seconds, timer_ch)
	go askQuestions(lines, correct_answer_ch, end_quiz_ch)
	total_correct := 0
out:
	for {
		select {
		case <-correct_answer_ch:
			total_correct++
		case <-end_quiz_ch:
			break out
		case <-timer_ch:
			break out
		}
	}
	fmt.Printf("total correct: %d/%d", total_correct, len(lines))
}

func askQuestions(lines [][]string, correct_answer_ch chan<- bool, end_quiz_ch chan<- bool) {
	total_correct := 0
	for _, line := range lines {
		question := line[0]
		answer := line[1]
		fmt.Println("question: ", question)
		userAnswer := getInputAndCleanIt()
		if userAnswer == answer {
			total_correct++
			correct_answer_ch <- true
		}
	}
	end_quiz_ch <- true
}

func getInputAndCleanIt() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func startTimer(seconds int, timer_ch chan<- bool) {
	time.Sleep(time.Duration(seconds) * time.Second)
	timer_ch <- true
}
