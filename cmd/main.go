package main

import (
	"log"

	"github.com/celestix/gotgproto"
	"github.com/celestix/gotgproto/sessionMaker"
	"github.com/glebarez/sqlite"
)

func main() {
	client, err := gotgproto.NewClient(
		// Get AppID from https://my.telegram.org/apps
		21326414,
		// Get ApiHash from https://my.telegram.org/apps
		"311562b87b3326cb20a864529520a121",
		// ClientType, as we defined above
		gotgproto.ClientTypePhone("+995557627581"),
		// Optional parameters of client
		&gotgproto.ClientOpts{
			Session: sessionMaker.SqlSession(sqlite.Open("reminderbot")),
		},
	)
	if err != nil {
		log.Fatalln("failed to start client:", err)
	}
	if err := client.Idle(); err != nil {
		log.Fatalln(err)
	}

}
