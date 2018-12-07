package domain

import (
	"fmt"
)

type ImageTweet struct {
	TextTweet
	Url string
}

func NewImageTweet(unTweet *TextTweet, url string) *ImageTweet {
	nuevoTweet := ImageTweet{*unTweet, url}
	return &nuevoTweet
}

func (t ImageTweet) String() string {
	return fmt.Sprintf("@%s: %s\n%s", t.User, t.Text, t.Url)
}
