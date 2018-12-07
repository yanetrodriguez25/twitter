package domain

import "time"

//Tweet ..Interfaz para Tweets
type Tweet interface {
	String() string
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetId() int
	SetId(int)
	SetDate(time.Time)
}
