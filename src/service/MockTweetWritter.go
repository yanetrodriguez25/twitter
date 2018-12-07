package service

import (
	"github.com/yanetrodriguez25/twitter/src/domain"
)

type MockTweetWritter struct {
	Tweet domain.Tweet
}

func NewMockTweetWritter() *MockTweetWritter {
	return &MockTweetWritter{}
}

func (m *MockTweetWritter) Write(unTweet domain.Tweet) {
	m.Tweet = unTweet
}

func (m *MockTweetWritter) GetLastSavedTweet() domain.Tweet {
	return m.Tweet
}
