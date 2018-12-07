package domain

import (
	"fmt"
)

type QuoteTweet struct {
	TextTweet
	Retweet Tweet
}

func NewQuoteTweet(unTweet *TextTweet, retweet Tweet) *QuoteTweet {
	nuevoTweet := QuoteTweet{*unTweet, retweet}
	return &nuevoTweet
}

func (t QuoteTweet) String() string {
	return fmt.Sprintf("@%s: %s\n\tRT:\"%s\"", t.User, t.Text, t.Retweet)
}
