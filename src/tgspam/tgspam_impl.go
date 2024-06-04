package tgspam

import (
	"github.com/go-vgo/robotgo"
)

func SendMessage(username, message string) error {
	//set default sleep time
	robotgo.MouseSleep = 1000
	//Open telegram
	robotgo.Move(0, 540)
	robotgo.Click()
	//Open saved messages
	robotgo.Move(100, 80)
	robotgo.Click()
	robotgo.TypeStr("Saved messages")
	robotgo.Move(100, 160)
	robotgo.Click()
	//Send username to the saved messages
	robotgo.Move(960, 1060)
	robotgo.Click()
	robotgo.TypeStr(username)
	if err := robotgo.KeyTap("enter"); err != nil {
		return err
	}
	//Open chat with the user and send message
	//TODO If user has long bio, message wont be sent; fix is needed
	//Maybe copy bio and check if its long manually, cant think of a better solution rn
	robotgo.Move(760, 1010)
	robotgo.Click()
	robotgo.Move(960, 500)
	robotgo.Click()
	robotgo.TypeStr(message)
	if err := robotgo.KeyTap("enter"); err != nil {
		return err
	}
	return nil
}

func ChangeAccount(accountNumber int) error {

	return nil
}
