package domain

import (
	"fmt"
	"time"
)

//TextTweet del tipo texto
type TextTweet struct {
	User string
	Text string
	Date *time.Time
	Id   int
}

func NewTextTweet(user, text string) *TextTweet {
	fecha := time.Now()
	nuevoTweet := TextTweet{User: user, Text: text, Date: &fecha}
	return &nuevoTweet
}

func (t TextTweet) String() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}

func (t TextTweet) GetUser() string {
	return t.User
}

func (t TextTweet) GetText() string {
	return t.Text
}

func (t TextTweet) GetId() int {
	return t.Id
}

func (t TextTweet) GetDate() *time.Time {
	return t.Date
}

func (t *TextTweet) SetId(id int) {
	t.Id = id
}

func (t *TextTweet) SetDate(date time.Time) {
	t.Date = &date
}
