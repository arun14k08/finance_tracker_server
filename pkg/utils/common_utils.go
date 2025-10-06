package utils

import "fmt"

func FormatAmount(amount float64) string {
	// Convert the float to a string with two decimal places
	return fmt.Sprintf("%.2f", amount)
}