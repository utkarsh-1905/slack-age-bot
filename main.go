package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	_ = os.Setenv("SLACK_BOT_TOKEN", "xoxb-3665573794772-3688446187175-pSC63gT1pJQK2ZjgNsehDXY6")
	_ = os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03LNU22A11-3688419745703-d64816665cd0cc0f21b31c4c17dcce62399b72b0558017e1991b27d719c9998e")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example:     "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			year := r.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				log.Fatalln(err)
			}
			age := 2022 - yob
			res := fmt.Sprintf("Your age is %d", age)
			_ = w.Reply(res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatalln(err)
	}

}
