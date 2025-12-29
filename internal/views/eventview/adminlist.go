package eventview

import (
	"time"

	"github.com/yihao03/reminding/internal/database/sqlc"
)

type EventListAdminView struct {
	ID           int32     `json:"id"`
	EventName    string    `json:"eventName"`
	Organiser    string    `json:"organiser"`
	IsOnline     bool      `json:"isOnline"`
	LocationName string    `json:"locationName"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	Details      string    `json:"details"`
	UserCount    int64     `json:"userCount"`
}

func ToEventListAdminView(events *[]sqlc.ListEventsAdminRow) *[]EventListAdminView {
	list := make([]EventListAdminView, len(*events))
	for i, event := range *events {
		view := EventListAdminView{
			ID:           event.ID,
			EventName:    event.EventName,
			Organiser:    event.Organiser.String,
			IsOnline:     event.IsOnline,
			LocationName: event.LocationName.String,
			StartTime:    event.StartTime.Time,
			EndTime:      event.EndTime.Time,
			Details:      event.Details.String,
			UserCount:    event.UserCount,
		}
		list[i] = view
	}
	return &list
}
