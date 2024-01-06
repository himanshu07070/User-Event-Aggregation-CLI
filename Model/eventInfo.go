package events

type Event struct {
	UserID    int    `json:"userId"`
	EventType string `json:"eventType"`
	Timestamp int64  `json:"timestamp"`
}
