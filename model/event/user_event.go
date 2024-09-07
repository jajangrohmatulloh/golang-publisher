package event

import "time"

type UserEvent struct {
	Id uint64 `json:"id"`
	Nama string `json:"nama"`
	LoginDatetime time.Time `json:"login_datetime"`
	Agent string `json:"agent"`
}