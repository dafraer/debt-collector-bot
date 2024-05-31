package tgspam

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"log"
	"time"
)

func SendMessage(ctx context.Context, username, message string) error {
	log.Println("entered messaging function")
	// Open Telegram Web
	err := chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf("https://web.telegram.org/k/#@%s", username)),
		chromedp.Sleep(30*time.Second), // Wait for loading
	)
	if err != nil {
		return err
	}
	log.Println("Navigated to the chat page")

	// Find the message input box and send the message
	err = chromedp.Run(ctx,
		chromedp.KeyEvent(message),
		chromedp.KeyEvent("\t"),
		chromedp.KeyEvent(kb.Enter),
		chromedp.Sleep(2*time.Second), // Wait for the message to be sent
	)
	if err != nil {
		return err
	}
	log.Println("message sent")
	return nil
}
