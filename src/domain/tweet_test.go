package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanetrodriguez25/twitter/src/domain"
)

func TestCanGetAStringFromATweet(t *testing.T) {
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	assert.Equal(t, tweet.String(), expectedText)
}
func TestCanGetAStringFromAImageTweet(t *testing.T) {
	tweetText := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweetImage := domain.NewImageTweet(tweetText, "google.com/image/gopher.png")

	// Validation
	expectedText := "@grupoesfera: This is my tweet\ngoogle.com/image/gopher.png"
	assert.Equal(t, tweetImage.String(), expectedText)
}

func TestCanGetAStringFromAQuoteTweet(t *testing.T) {

	//Tweet image
	tweetText := domain.NewTextTweet("grupoesfera", "This is my tweetImage")
	tweetImage := domain.NewImageTweet(tweetText, "google.com/image/gopher.png")
	//Tweet Quote
	tweetText2 := domain.NewTextTweet("grupoesfera", "This is my tweetQuote")
	tweetQuote := domain.NewQuoteTweet(tweetText2, tweetImage)

	// Validation
	expectedText := "@grupoesfera: This is my tweetQuote\n\tRT:\"@grupoesfera: This is my tweetImage\ngoogle.com/image/gopher.png\""
	assert.Equal(t, tweetQuote.String(), expectedText)
}
