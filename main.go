package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type QuesAns struct {
	question string
	answer   string
}

func main() {
	//set what kind of file we are looking for with usage
	csvFile := flag.String("csv", "questions.csv", "a csv file contain'question and answer'")

	//receive time limit from user
	timeLimit := flag.Int("time", 30, "Application will wait for a limited time duration")

	flag.Parse()

	//open file
	osFile, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Unable to open the file: %v", osFile))
	}

	//read the file
	fileData := csv.NewReader(osFile)
	lines, err := fileData.ReadAll()
	if err != nil {
		exit("Unable to read file")
	}

	//divide questions and answers
	qa := GetQuestionsAnswers(lines)

	//start timer
	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))

	//perform quiz, show results
	fmt.Println(ShowQuesCalcAns(qa, *timer))
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func GetQuestionsAnswers(lines [][]string) []QuesAns {

	qa := make([]QuesAns, len(lines))

	for index, value := range lines {
		qa[index] = QuesAns{
			question: value[0],
			answer:   strings.TrimSpace(value[1]),
		}
	}
	return qa
}

func ShowQuesCalcAns(qa []QuesAns, timeLimit time.Timer) string {
	correct := 0

	for index, value := range qa {
		fmt.Printf("Question: %d --> %s =\n", index+1, value.question)
		answerChan := make(chan string)
		go func() {
			var userAnswer string
			fmt.Scanf("%s\n", &userAnswer)
			answerChan <- userAnswer
		}()
		select {
		case <-timeLimit.C:
			return fmt.Sprintf("\nRan out of time,Correct Answers: %d out of total Attempt %d of total questions %d", correct, index+1, len(qa))
		case userAnswer := <-answerChan:
			if userAnswer == value.answer {
				correct++
			}
		}

	}
	return fmt.Sprintf("Correct Answers: %d out of total %d", correct, len(qa))
}
