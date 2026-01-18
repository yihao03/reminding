package moodview

import (
	"github.com/yihao03/reminding/internal/database/sqlc"
)

type MoodCountView struct {
	Mood  int32 `json:"mood"`
	Count int64 `json:"count"`
}

func ToMoodCountView(rows *[]sqlc.GetMonthlyMoodCountByUserUidRow) *[]MoodCountView {
	if rows == nil {
		res := make([]MoodCountView, 0)
		return &res
	}

	res := make([]MoodCountView, len(*rows))
	for i, v := range *rows {
		res[i] = MoodCountView{
			Mood:  v.Mood,
			Count: v.OccurrenceCount,
		}
	}
	return &res
}
