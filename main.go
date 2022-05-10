package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type QuesAns struct {
	question string
	answer   string
}

func main() {
	//set what kind of file we are looking for with usage
	csvFile := flag.String("csv", "questions.csv", "a csv file contain'question and answer'")
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

	//perform quiz, show results
	fmt.Println(ShowQuesCalcAns(qa))
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

func ShowQuesCalcAns(qa []QuesAns) string {

	correct := 0

	for index, value := range qa {
		fmt.Printf("Question: %d --> %s =\n", index+1, value.question)
		var userAnswer string
		fmt.Scanf("%s\n", &userAnswer)
		if userAnswer == value.answer {
			correct++
		}
	}
	return fmt.Sprintf("Correct Answers: %d out of total %d", correct, len(qa))
}
