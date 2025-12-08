package eventview

import (
	"time"

	"github.com/yihao03/reminding/internal/database/sqlc"
)

type EventDetailedView struct {
	ID           int32     `json:"id"`
	Organiser    string    `json:"organiser"`
	IsOnline     bool      `json:"isOnline"`
	LocationName string    `json:"locationName"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	EventName    string    `json:"eventName"`
	CreatedAt    time.Time `json:"createdAt"`
	Details      string    `json:"details"`
}

func ToDetailedEventView(event *sqlc.Event) *EventDetailedView {
	return &EventDetailedView{
		ID:           event.ID,
		Organiser:    event.Organiser.String,
		IsOnline:     event.IsOnline,
		LocationName: event.LocationName.String,
		StartTime:    event.StartTime.Time,
		EndTime:      event.EndTime.Time,
		EventName:    event.EventName,
		CreatedAt:    event.CreatedAt.Time,
		Details:      event.Details.String,
	}
}
