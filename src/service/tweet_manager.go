package service
import (
	"github.com/yanetrodriguez25/twitter/src/domain"
	"errors"
)

var tweet *domain.Tweet

func PublishTweet(unTweet *domain.Tweet) error{
	var err error
	if unTweet.User == ""{
		err = errors.New("user is required")
		return err
	}
	 if unTweet.Text == ""{
		err = errors.New("text is required")
		return err
	} 
	 if len(unTweet.Text) > 140 {
		err = errors.New("text exceeds 140 characters")
		return err
	} 
	tweet = unTweet
	return nil
}

func GetTweet() *domain.Tweet{
	return tweet;
}