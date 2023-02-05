package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter array lenght: ")
	arrLenText, _ := reader.ReadString('\n')
	arrLenText = strings.Replace(arrLenText, "\r\n", "", -1)

	arrLen, err := strconv.Atoi(arrLenText)
	if err != nil {
		fmt.Printf("error: input not a number or wrong format, detail: %v\n", err)
		return
	}
	if arrLen < 1 || arrLen > 100000 {
		fmt.Println("error: array length cannot be less than 1 or greater than 100 thousand")
		return
	}

	fmt.Print("Enter array elements (separated by a space): ")
	arrValuesText, _ := reader.ReadString('\n')
	arrValuesText = strings.Replace(arrValuesText, "\r\n", "", -1)

	arrVals, err := stringToIntArr(arrValuesText, arrLen)
	if err != nil {
		fmt.Println(err)
		return
	}

	var result string
	isMatch, meanElement := isLeftEqualToRight(arrVals)
	if isMatch {
		result = "YES"
	} else {
		result = "NO"
	}

	// fmt.Println("Input (string)", arrLenText)
	// fmt.Println("Input (integer)", arrLen)
	// fmt.Println("Array:", arrVals)
	fmt.Println("Mean element:", meanElement)
	fmt.Println("Left and right sub-arrays is Matched?", result)
}

// With number of array elements constraint
// So, number of array elements must be matched with array length
func stringToIntArr(s string, l int) ([]int, error) {
	s = strings.TrimSpace(s)
	stringArr := strings.Split(s, " ")
	if len(stringArr) != l {
		err := errors.New("error: number of entered elements doesn't match with the array lenght or wrong format!")
		return []int{}, err
	}

	intArr := make([]int, l)
	for i, val := range stringArr {
		d, err := strconv.Atoi(val)
		if err != nil {
			return []int{}, err
		}
		if d < 1 || d > 10000 {
			err := errors.New("error: array element cannot be less than 1 or greater than 10K")
			return []int{}, err
		}
		intArr[i] = d
	}

	return intArr, nil
}

func isLeftEqualToRight(arr []int) (bool, int) {
	l := len(arr)
	var divider int
	for i, val := range arr {
		if i == 0 {
			if arrSum(arr[1:l]) == 0 {
				return true, val
			}
		}

		leftSubArr := arrSum(arr[0:i])
		rightSubArr := arrSum(arr[i+1 : l])
		divider = val

		if leftSubArr == rightSubArr {
			return true, divider
		}
	}
	return false, 0
}

func arrSum(arr []int) int {
	sumResult := 0
	for _, val := range arr {
		sumResult += val
	}
	return sumResult
}
