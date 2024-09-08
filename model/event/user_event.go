package event

type UserEvent struct {
	Id uint64 `json:"id"`
	Nama string `json:"nama"`
	LoginDatetime string `json:"login_datetime"`
	Agents string `json:"agents"`
}