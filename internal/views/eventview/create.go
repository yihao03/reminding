package eventview

import (
	"time"

	"github.com/yihao03/reminding/internal/database"
	"github.com/yihao03/reminding/internal/database/sqlc"
)

type EventCreateView struct {
	Organiser    string    `json:"organiser" validate:"required"`
	IsOnline     bool      `json:"isOnline" validate:"required"`
	LocationName string    `json:"locationName" validate:"required"`
	StartTime    time.Time `json:"startTime" validate:"required"`
	EndTime      time.Time `json:"endTime" validate:"required"`
	EventName    string    `json:"eventName" validate:"required"`
	Details      string    `json:"details" validate:"required"`
}

func ToCreateParams(event *EventCreateView) *sqlc.CreateEventParams {
	return &sqlc.CreateEventParams{
		Organiser:    database.ToPGText(event.Organiser),
		IsOnline:     event.IsOnline,
		LocationName: database.ToPGText(event.LocationName),
		StartTime:    database.ToPGTime(&event.StartTime),
		EndTime:      database.ToPGTime(&event.EndTime),
		EventName:    event.EventName,
		Details:      database.ToPGText(event.Details),
	}
}
