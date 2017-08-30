package main

import (
	"fmt"
	"time"
)

func isValidMonth(mm int) bool {
	if mm >= 1 && mm <= 12 {
		return true
	}
	return false
}

func main() {
	var i, mm, k int

	//date input validation
	for (i > 31) || (i < 1) {
		fmt.Print("Enter valid birth date: ")
		_, err := fmt.Scan(&i)
		if err != nil {
			fmt.Println(err)
		}
	}

	//month input validation
	for isValidMonth(mm) == false {
		fmt.Print("Enter the valid month: ")
		_, err := fmt.Scan(&mm)
		if err != nil {
			fmt.Println(err)
		}
	}

	t := time.Now()
	currentmonth := t.Month()
	n := int(currentmonth)
	currentdate := (t.Day())
	currentyear := (t.Year())

	//year input validation
	for (k > currentyear) || (k < 1) {
		fmt.Print("Enter the valid year: ")
		_, err := fmt.Scan(&k)
		if err != nil {
			fmt.Println(err)
		}
	}

	//validate input complete date
	for k == currentyear {
		if mm > n {
			break
		} else {
			if (mm == n) && (i > currentdate) {
				break
			}
		}
	}
	var date, month int
	if currentdate < i {
		date = (30 - i) + currentdate
		month = (n - mm) - 1
		fmt.Println("Age as per lunar calender: ", currentyear-k, "years", month, "months", date, "days")
	} else {
		fmt.Println("Age as per lunar calender: ", currentyear-k, "years", n-mm, "months", currentdate-i, "days")
	}
}
