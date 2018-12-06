package main

import (
	"github.com/abiosoft/ishell"
	"github.com/yanetrodriguez25/twitter/src/service"
	"github.com/yanetrodriguez25/twitter/src/domain"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Write your user: ")
			user := c.ReadLine()

			c.Print("Write your tweet: ")
			text := c.ReadLine()
 
			tweet := domain.NewTweet(user, text)
			service.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()
			fecha := *tweet.Date;
			fechaFormateada := fecha.Format("02-01-2006 15:04:05")
			c.Println(tweet.User, tweet.Text, fechaFormateada)

			return
		},
	})

	shell.Run()

}
