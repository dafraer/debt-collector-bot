package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.MouseSleep = 100

	robotgo.Move(700, 870)

	robotgo.Click()

	robotgo.TypeStr("Ты должен мне миллион рублей")

	robotgo.KeyTap("enter")
}
