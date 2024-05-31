package main

import (
	"collector/src/tgspam"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // Set to false to run Chrome in a non-headless mode
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
	)

	// Setup context and chromedp
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	username := "fiodop"
	message := "Hello, this is a test message!"

	// Send message
	if err := tgspam.SendMessage(ctx, username, message); err != nil {
		fmt.Println("Failed to send message:", err)
	} else {
		fmt.Println("Message sent successfully!")
	}
}
