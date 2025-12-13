package eventview

import (
	"time"

	"github.com/yihao03/reminding/internal/database/sqlc"
)

type EventListView struct {
	ID           int32     `json:"id"`
	Organiser    string    `json:"organiser"`
	IsOnline     bool      `json:"isOnline"`
	LocationName string    `json:"locationName"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	EventName    string    `json:"eventName"`
}

func ToEventListView(events *[]sqlc.ListEventsWithUserRegistrationRow) *[]EventListView {
	list := make([]EventListView, len(*events))
	for i, event := range *events {
		view := EventListView{
			ID:           event.ID,
			Organiser:    event.Organiser.String,
			IsOnline:     event.IsOnline,
			LocationName: event.LocationName.String,
			StartTime:    event.StartTime.Time,
			EndTime:      event.EndTime.Time,
			EventName:    event.EventName,
		}
		list[i] = view
	}
	return &list
}
