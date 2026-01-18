package moodview

import (
	"github.com/yihao03/reminding/internal/database/sqlc"
)

type MoodReadView struct {
	LoggedToday bool            `json:"loggedToday"`
	MoodCount   []MoodCountView `json:"moodCount"`
}

func ToMoodReadView(loggedToday bool, moodCount *[]sqlc.GetMonthlyMoodCountByUserUidRow) *MoodReadView {
	return &MoodReadView{
		LoggedToday: loggedToday,
		MoodCount:   *ToMoodCountView(moodCount),
	}
}
