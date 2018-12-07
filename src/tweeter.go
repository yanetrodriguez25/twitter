package main

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/yanetrodriguez25/twitter/src/domain"
	"github.com/yanetrodriguez25/twitter/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")
	manager := service.NewTweetManager()
	var tweet domain.Tweet
	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Write your user: ")
			user := c.ReadLine()
			selectedOption := c.MultiChoice([]string{"Text tweet", "Image tweet", "Quote tweet"}, "Select an option")
			switch selectedOption {
			case 0:
				defer c.ShowPrompt(true)
				c.Print("Write your text: ")
				text := c.ReadLine()
				tweet = domain.NewTextTweet(user, text)
			case 1:
				defer c.ShowPrompt(true)
				c.Print("Write your text: ")
				text := c.ReadLine()
				c.Print("Write your url: ")
				url := c.ReadLine()
				textTweet := domain.NewTextTweet(user, text)
				tweet = domain.NewImageTweet(textTweet, url)
			case 2:
				defer c.ShowPrompt(true)
				c.Print("Write your text: ")
				text := c.ReadLine()
				textTweet := domain.NewTextTweet(user, text)
				tweets := manager.GetTweets()
				tweetsToQuoteArray := make([]string, 0)
				for _, t := range tweets {
					tweetsToQuoteArray = append(tweetsToQuoteArray, t.String())
				}
				tweetSelected := c.MultiChoice(tweetsToQuoteArray, "Select a tweet")
				retweet := manager.GetTweetById(tweetSelected)
				tweet = domain.NewQuoteTweet(textTweet, retweet)
			}

			manager.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := manager.GetTweets()
			tweet := tweets[len(tweets)-1]
			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "listTweets",
		Help: "Shows a list of tweets",
		Func: func(c *ishell.Context) {
			tweets := manager.GetTweets()
			for i := 0; i < len(tweets); i++ {
				c.Println(tweets[i])
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "findByID",
		Help: "find a tweet by ID",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Write an id: ")
			input := c.ReadLine()
			id, ok := strconv.Atoi(input)
			tweetsCount := len(manager.GetTweets())
			if id <= tweetsCount && ok == nil {
				tweet := manager.GetTweetById(id)
				c.Println(tweet)
			} else {
				c.Println("Incorrect id")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "CountTweetsByUser",
		Help: "Counts User's tweets ",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Write an user: ")
			user := c.ReadLine()

			c.Println(manager.CountTweetsByUser(user))

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "GetTweetsByUser",
		Help: "Gets User's tweets ",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Write an user: ")
			user := c.ReadLine()
			for _, tweet := range manager.GetTweetsByUser(user) {
				c.Println(tweet)
			}

			return
		},
	})

	shell.Run()

}
