package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter due date (dd mm yyyy): ")
	dueDateText, _ := reader.ReadString('\n')
	dueDateText = strings.Replace(dueDateText, "\r\n", "", -1)
	dueDate, err := stringToDate(dueDateText)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter return date (dd mm yyyy): ")
	returnDateText, _ := reader.ReadString('\n')
	returnDateText = strings.Replace(returnDateText, "\r\n", "", -1)
	returnDate, err := stringToDate(returnDateText)
	if err != nil {
		fmt.Println(err)
		return
	}

	if returnDate.Before(dueDate) {
		fmt.Println("error: be logic, due date cannot be greater than return date!")
		return
	}

	/* fmt.Println("Loan Date:", loanDate)
	fmt.Println("Due Date:", dueDate) */
	fmt.Println("Loan Charge:", calculateLoanFine(dueDate, returnDate))
}

func stringToDate(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	dateStringSplit := strings.Split(s, " ")
	if len(dateStringSplit) != 3 {
		return time.Time{}, errors.New("error: invalid date or wrong format")
	}

	// Year constraint will checked here
	yearString := dateStringSplit[2]
	yearInt, err := strconv.Atoi(yearString)
	if err != nil {
		return time.Time{}, err
	}
	if yearInt < 1 || yearInt > 4000 {
		return time.Time{}, errors.New("error: year is out of range")
	}

	// Processing when entered date has less than 4 digit, so it will add zero in front of it
	var addedZero []string
	if len(yearString) < 4 {
		for i := 1; i <= (4 - len(yearString)); i++ {
			addedZero = append(addedZero, "0")
		}
	}
	newYearString := " " + strings.Join(addedZero, "") + yearString
	newDateString := strings.Join(dateStringSplit[0:2], " ") + newYearString

	// Date and month constraint automatically will be checked by time.Parse() function
	dt, err := time.Parse("2 1 2006", newDateString)
	if err != nil {
		return time.Time{}, err
	}
	return dt, nil
}

func calculateLoanFine(startDate, endDate time.Time) int {
	diffYear := endDate.Year() - startDate.Year()
	diffMonth := int(endDate.Month() - startDate.Month())
	diffDay := endDate.Day() - startDate.Day()

	var charge int = 0
	if diffYear > 0 {
		charge = 12000
		return charge
	}
	if diffMonth > 0 {
		charge = diffMonth * 500
		return charge
	}
	if diffDay > 0 {
		charge = diffDay * 15
		return charge
	}
	return charge
}
