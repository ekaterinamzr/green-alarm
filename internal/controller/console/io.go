package console

import (
	"fmt"
	"time"
)

func inputString(prompt string) (string, error) {
	var s string
	fmt.Printf("%s: ", prompt)
	_, err := fmt.Scanf("%s", &s)
	return s, err
}

func inputInt(prompt string) (int, error) {
	var i int
	fmt.Printf("%s: ", prompt)
	_, err := fmt.Scanf("%d", &i)
	return i, err
}

func inputFloat(prompt string) (float64, error) {
	var f float64
	fmt.Printf("%s: ", prompt)
	_, err := fmt.Scanf("%f", &f)
	return f, err
}

func inputDate(prompt string) (time.Time, error) {
	var s string
	fmt.Printf("%s (dd-mm-yyyy): ", prompt)
	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse("02-01-2006", s)
}
