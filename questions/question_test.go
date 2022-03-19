package questions

import (
	"reflect"
	"testing"
)

func TestQuestion1(t *testing.T) {
	question1Testcase := []int{1, 2, 3, 4, 4}
	result := []int{1, 2, 3, 4, 5}

	if reflect.DeepEqual(result, Question1(question1Testcase)) == false {
		t.Error("Question 1: Incremental Array : ", Question1(question1Testcase))
	}
}

func TestQuestion2(t *testing.T) {
	question2Testcase := "We test coders. Give us a try?"
	result := 4

	if result != Question2(question2Testcase) {
		t.Error("Question 2: Maximum words in sentence : ", Question2(question2Testcase))
	}
}

func TestQuestion3(t *testing.T) {
	question3TestcaseA := []int{3, 8, 9, 7, 6}
	question3TestcaseK := 3
	result := []int{9, 7, 6, 3, 8}

	if reflect.DeepEqual(result, Question3(question3TestcaseA, question3TestcaseK)) == false {
		t.Error("Question 3: Shifted Array : ", Question3(question3TestcaseA, question3TestcaseK))
	}
}

func TestQuestion4(t *testing.T) {
	question4Testcase := []int{9, 3, 9, 3, 9, 7, 9}
	result := 7

	if result != Question4(question4Testcase) {
		t.Error("Question 4: Find pair : ", Question4(question4Testcase))
	}
}

func TestQuestion5(t *testing.T) {
	question5Testcase := []int{1, 5, 3, 4, 7}
	result := false

	if result != Question5(question5Testcase) {
		t.Error("Question 5: Can be sorted? : ", Question5(question5Testcase))
	}

	question5Testcase = []int{1, 3, 5, 3, 7}
	result = true
	if result != Question5(question5Testcase) {
		t.Error("Question 5: Can be sorted? : ", Question5(question5Testcase))
	}

	question5Testcase = []int{1}
	result = false
	if result != Question5(question5Testcase) {
		t.Error("Question 5: Can be sorted? : ", Question5(question5Testcase))
	}
}
