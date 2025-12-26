package eventview

import (
	"time"

	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/views/userview"
)

type EventAdminView struct {
	ID           int32               `json:"id"`
	Organiser    string              `json:"organiser"`
	IsOnline     bool                `json:"isOnline"`
	LocationName string              `json:"locationName"`
	StartTime    time.Time           `json:"startTime"`
	EndTime      time.Time           `json:"endTime"`
	EventName    string              `json:"eventName"`
	CreatedAt    time.Time           `json:"createdAt"`
	Details      string              `json:"details"`
	Users        []userview.UserView `json:"users"`
}

func ToAdminEventView(event *sqlc.Event, users *[]sqlc.User) *EventAdminView {
	usersView := make([]userview.UserView, len(*users))
	for i, user := range *users {
		usersView[i] = *userview.ToUserView(&user)
	}
	return &EventAdminView{
		ID:           event.ID,
		Organiser:    event.Organiser.String,
		IsOnline:     event.IsOnline,
		LocationName: event.LocationName.String,
		StartTime:    event.StartTime.Time,
		EndTime:      event.EndTime.Time,
		EventName:    event.EventName,
		CreatedAt:    event.CreatedAt.Time,
		Details:      event.Details.String,
		Users:        usersView,
	}
}
