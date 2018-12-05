package service_test // este  no se compila.. tiene que tener el mismo nombre del paquete + _test
import (
	"testing"

	"github.com/yanetrodriguez25/twitter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T){
	var tweet string = "This is my first tweet"

	service.PublishTweet(tweet)

	if(service.Tweet != tweet){
		t.Error("Expected tweet is", tweet)
	}
}