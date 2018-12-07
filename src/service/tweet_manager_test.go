package service_test // este  no se compila.. tiene que tener el mismo nombre del paquete + _test
import (
	"testing"

	"github.com/yanetrodriguez25/twitter/src/domain"
	"github.com/yanetrodriguez25/twitter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	manager := service.NewTweetManager()
	var tweet *domain.TextTweet
	user := "yanetrodriguez25"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	manager.PublishTweet(tweet)

	publishedTweet := manager.GetTweets()[0]
	if publishedTweet.GetUser() != user && publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \n but is %s: %s", user, text, publishedTweet.GetUser(), publishedTweet.GetText())
	}

	if publishedTweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	manager := service.NewTweetManager()
	var tweet *domain.TextTweet
	var user string
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	var err error
	_, err = manager.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	manager := service.NewTweetManager()

	var tweet *domain.TextTweet
	var text string
	user := "yanetrodriguez25"
	tweet = domain.NewTextTweet(user, text)

	var err error
	_, err = manager.PublishTweet(tweet)

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	manager := service.NewTweetManager()

	var tweet *domain.TextTweet
	user := "yanetrodriguez25"
	text := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	tweet = domain.NewTextTweet(user, text)

	var err error
	_, err = manager.PublishTweet(tweet)

	if err != nil && err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	manager := service.NewTweetManager()

	var tweet, secondTweet *domain.TextTweet

	text := "hola soy un tweet 1"
	user := "yanetrodriguez25"
	tweet = domain.NewTextTweet(user, text)

	text2 := "hola soy un tweet 2"
	user2 := "andres"
	secondTweet = domain.NewTextTweet(user2, text2)

	id, _ := manager.PublishTweet(tweet)
	id2, _ := manager.PublishTweet(secondTweet)

	publishedTweets := manager.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, id, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, id2, user2, text2) {
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T) {
	manager := service.NewTweetManager()

	var tweet *domain.TextTweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	id, _ = manager.PublishTweet(tweet)

	publishedTweet := manager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	manager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	manager.PublishTweet(tweet)
	manager.PublishTweet(secondTweet)
	manager.PublishTweet(thirdTweet)

	count := manager.CountTweetsByUser(user)

	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	manager := service.NewTweetManager()
	var tweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	// publish the 3 tweets
	id, _ := manager.PublishTweet(tweet)
	id2, _ := manager.PublishTweet(secondTweet)
	manager.PublishTweet(thirdTweet)

	// Operation
	tweets := manager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 { /* handle error */
	}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet

	if !isValidTweet(t, firstPublishedTweet, id, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, id2, user, secondText) {
		return
	}
}

func isValidTweet(t *testing.T, publishedTweet domain.Tweet, id int, user string, text string) bool {
	if publishedTweet.GetUser() == user && publishedTweet.GetText() == text && publishedTweet.GetId() == id {
		return true
	}
	t.Errorf("Expected %s: %s, %d get %s, %s %d", user, text, id, publishedTweet.GetUser(), publishedTweet.GetText(), publishedTweet.GetId())
	return false
}
