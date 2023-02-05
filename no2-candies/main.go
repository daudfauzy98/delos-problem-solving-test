package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter amount of students, candies, and first_student (separated by a space): ")
	inputText, _ := reader.ReadString('\n')
	inputText = strings.Replace(inputText, "\r\n", "", -1)

	students, candies, firstStudent, err := extractInput(inputText)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Student's number with sour candy:", getStudentWithSourCandy(students, candies, firstStudent))
}

func extractInput(s string) (int, int, int, error) {
	s = strings.TrimSpace(s)
	textSplit := strings.Split(s, " ")
	if len(textSplit) != 3 {
		return 0, 0, 0, errors.New("error: wrong input format")
	}

	// Students ammount constraint
	st, err := strconv.Atoi(textSplit[0])
	if err != nil {
		return 0, 0, 0, errors.New("error: input only positive integers number")
	}
	if st < 1 || st > int(math.Pow(10, 9)) {
		return 0, 0, 0, errors.New("error: students ammounts cannot be less than 1 or greater than 1 billion")
	}

	// Candies ammount constraint
	cd, err := strconv.Atoi(textSplit[1])
	if err != nil {
		return 0, 0, 0, errors.New("error: input only positive integers number")
	}
	if cd < 1 || cd > int(math.Pow(10, 9)) {
		return 0, 0, 0, errors.New("error: candies ammounts cannot be less than 1 or greater than 1 billion")
	}

	// First student ammount constraint
	fs, err := strconv.Atoi(textSplit[2])
	if err != nil {
		return 0, 0, 0, errors.New("error: input only positive integers number")
	}
	if fs < 1 || fs > st {
		return 0, 0, 0, errors.New("error: first student cannot be less than 1 or greater than students ammount")
	}

	return st, cd, fs, nil
}

func getStudentWithSourCandy(st, cd, fs int) int {
	var cdArr []int
	var srCd int
	currStd := fs

	for i := 1; i <= cd; i++ {
		cdArr = append(cdArr, currStd)
		srCd = currStd

		if currStd == st {
			currStd = 1
		} else {
			currStd++
		}
	}
	return srCd
}
