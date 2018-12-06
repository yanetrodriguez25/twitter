package service_test // este  no se compila.. tiene que tener el mismo nombre del paquete + _test
import (
	"testing"
	"github.com/yanetrodriguez25/twitter/src/domain"
	"github.com/yanetrodriguez25/twitter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T){
	var tweet *domain.Tweet
	user :="yanetrodriguez25"
	text :="This is my first tweet"
	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	publishedTweet := service.GetTweet()
	if(publishedTweet.User != user && publishedTweet.Text != text){
		t.Errorf("Expected tweet is %s: %s \n but is %s: %s", user, text, publishedTweet.User, publishedTweet.Text)
	}

	if(publishedTweet.Date == nil){
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T){
	var tweet *domain.Tweet
	var user string
	text :="This is my first tweet"
	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required"{
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T){
	var tweet *domain.Tweet
	var text string
	user :="yanetrodriguez25"
	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "text is required"{
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T){
	var tweet *domain.Tweet
	user :="yanetrodriguez25"
	text :="aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "text exceeds 140 characters"{
		t.Error("Expected error is text exceeds 140 characters")
	}
}