package service

import (
	"errors"
	"strings"

	"github.com/yanetrodriguez25/twitter/src/domain"
)

type TweetManager struct {
	tweets       []domain.Tweet
	tweetsMap    map[string][]domain.Tweet
	tweetWritter TweetWritter
}

func NewTweetManager(unTweetWritter TweetWritter) *TweetManager {
	tweetManager := TweetManager{}
	tweetManager.initializeService(unTweetWritter)

	return &tweetManager
}

func (tm *TweetManager) initializeService(unTweetWritter TweetWritter) {
	tm.tweets = make([]domain.Tweet, 0)
	tm.tweetsMap = make(map[string][]domain.Tweet)
	tm.tweetWritter = unTweetWritter

}

func (tm *TweetManager) PublishTweet(unTweet domain.Tweet) (int, error) {
	var err error
	if unTweet.GetUser() == "" {
		err = errors.New("user is required")
		return -1, err
	}
	if unTweet.GetText() == "" {
		err = errors.New("text is required")
		return -1, err
	}
	if len(unTweet.GetText()) > 140 {
		err = errors.New("text exceeds 140 characters")
		return -1, err
	}
	unTweet.SetId(len(tm.tweets))
	tm.tweets = append(tm.tweets, unTweet)

	tm.tweetsMap[unTweet.GetUser()] = append(tm.tweetsMap[unTweet.GetUser()], unTweet)

	tm.tweetWritter.Write(unTweet)
	return unTweet.GetId(), nil

}

func (tm *TweetManager) GetTweetWritter() TweetWritter {
	return tm.tweetWritter
}

func (tm *TweetManager) GetTweets() []domain.Tweet {
	return tm.tweets
}

func (tm *TweetManager) GetTweetById(id int) domain.Tweet {
	return tm.tweets[id]
}

func (tm *TweetManager) CountTweetsByUser(user string) int {
	var count int
	for _, tweet := range tm.tweets {
		if tweet.GetUser() == user {
			count++
		}
	}

	return count
}

func (tm *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	return tm.tweetsMap[user]
}

func (tm *TweetManager) SearchTweetsContaining(query string, channel chan domain.Tweet) {
	go func() {
		for _, t := range tm.GetTweets() {
			if strings.Contains(t.GetUser(), query) || strings.Contains(t.GetText(), query) {
				channel <- t
			}
		}
	}()
}
