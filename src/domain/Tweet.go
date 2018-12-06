package domain
import "time"

type Tweet struct{
	User string
	Text string
	Date *time.Time
}

func NewTweet(user, text string) *Tweet {
	fecha := time.Now()
 	nuevoTweet := Tweet{User: user, Text: text, Date: &fecha}
 return &nuevoTweet
}