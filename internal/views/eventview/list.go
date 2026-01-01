package eventview

import (
	"time"

	"github.com/yihao03/reminding/internal/database/sqlc"
)

type UserEventList struct {
	ID           int32     `json:"id"`
	Organiser    string    `json:"organiser"`
	IsOnline     bool      `json:"isOnline"`
	LocationName string    `json:"locationName"`
	State        string    `json:"state" validate:"oneof='Johor' 'Kedah' 'Kelantan' 'Melaka' 'Negeri Sembilan' 'Pahang' 'Perak' 'Perlis' 'Penang' 'Sabah' 'Sarawak' 'Selangor' 'Terengganu'"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	EventName    string    `json:"eventName"`
	IsRegistered bool      `json:"isRegistered"`
}

func ToUserEventList(events *[]sqlc.ListEventsUserRow) *[]UserEventList {
	list := make([]UserEventList, len(*events))
	for i, event := range *events {
		view := UserEventList{
			ID:           event.ID,
			Organiser:    event.Organiser.String,
			IsOnline:     event.IsOnline,
			LocationName: event.LocationName.String,
			State:        string(event.State.States),
			StartTime:    event.StartTime.Time,
			EndTime:      event.EndTime.Time,
			EventName:    event.EventName,
			IsRegistered: event.IsRegistered,
		}
		list[i] = view
	}
	return &list
}
