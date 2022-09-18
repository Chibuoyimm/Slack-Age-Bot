package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	// Setting environment variables
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4094726156805-4094745245205-qmOg929JjEZzw9QQHIvIwpMW")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A042P2KJP70-4097597882915-656e42ae5d745a79ee5dab2cd40d4bd84b56bb53fa135c004733c3e5484c1ca4")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "age calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	// This block makes sure the bot cancels at the end of the main function by calling defer
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
