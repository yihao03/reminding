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
	State       string `json:"state" validate:"oneof='Johor' 'Kedah' 'Kelantan' 'Melaka' 'Negeri Sembilan' 'Pahang' 'Perak' 'Perlis' 'Penang' 'Sabah' 'Sarawak' 'Selangor' 'Terengganu'"`
	EndTime      time.Time `json:"endTime"`
	EventName    string    `json:"eventName"`
	CreatedAt    time.Time `json:"createdAt"`
	Details      string    `json:"details"`
	IsRegistered bool      `json:"isRegistered"`
}

func ToDetailedEventView(event *sqlc.GetEventByIdAndUidRow) *EventDetailedView {
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
		IsRegistered: event.IsRegistered,
	}
}
