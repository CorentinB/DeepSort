package logging

import "github.com/labstack/gommon/color"

// Error log error messages.
func Error(str string) {
	color.Println(color.Red("[✖] ") + color.Yellow(str))
}
