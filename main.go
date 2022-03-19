package main

import (
	"fmt"

	"github.com/popoiuioopp/ace/questions"
)

func main() {
	question1Testcase := []int{1, 2, 3, 4, 4}

	question2Testcase := "We test coders. Give us a try?"

	question3TestcaseA := []int{3, 8, 9, 7, 6}
	question3TestcaseK := 3

	question4Testcase := []int{9, 3, 9, 3, 9, 7, 9}

	question5Testcase := []int{1, 5, 3, 3, 7}

	fmt.Println("Question 1: Incremental Array : ", questions.Question1(question1Testcase))
	fmt.Println("Question 2: Maximum words in sentence : ", questions.Question2(question2Testcase))
	fmt.Println("Question 3: Shifted Array : ", questions.Question3(question3TestcaseA, question3TestcaseK))
	fmt.Println("Question 4: Find pair : ", questions.Question4(question4Testcase))
	fmt.Println("Question 5: Can be sorted? : ", questions.Question5(question5Testcase))
}
