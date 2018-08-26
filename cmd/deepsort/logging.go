package main

import "github.com/labstack/gommon/color"

// Error log error messages.
func logError(str string, prefix string) {
	color.Println(color.Red("[✖] ") + color.Red(prefix+" ") + color.Yellow(str))
}

// Success log success messages.
func logSuccess(str string, prefix string) {
	color.Println(color.Green("[✔] ") + color.Green(prefix+" ") + color.Yellow(str))
}
