package eventview

import (
	"time"

	"github.com/yihao03/reminding/internal/database/sqlc"
)

type EventListWithRegistrationStatus struct {
	ID           int32     `json:"id"`
	Organiser    string    `json:"organiser"`
	IsOnline     bool      `json:"isOnline"`
	LocationName string    `json:"locationName"`
	State       string `json:"state" validate:"oneof='Johor' 'Kedah' 'Kelantan' 'Melaka' 'Negeri Sembilan' 'Pahang' 'Perak' 'Perlis' 'Penang' 'Sabah' 'Sarawak' 'Selangor' 'Terengganu'"`
	EndTime      time.Time `json:"endTime"`
	EventName    string    `json:"eventName"`
	IsRegistered bool      `json:"isRegistered"`
}

func ToEventListWithRegistrationStatus(events *[]sqlc.ListEventsWithRegistrationStatusRow) *[]EventListWithRegistrationStatus {
	list := make([]EventListWithRegistrationStatus, len(*events))
	for i, event := range *events {
		view := EventListWithRegistrationStatus{
			ID:           event.ID,
			Organiser:    event.Organiser.String,
			IsOnline:     event.IsOnline,
			LocationName: event.LocationName.String,
			StartTime:    event.StartTime.Time,
			EndTime:      event.EndTime.Time,
			EventName:    event.EventName,
			IsRegistered: event.IsRegistered,
		}
		list[i] = view
	}
	return &list
}
