package service

import (
	"github.com/yanetrodriguez25/twitter/src/domain"
)

type TweetWritter interface {
	Write(domain.Tweet)
}
