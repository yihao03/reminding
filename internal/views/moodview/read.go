package moodview

import "github.com/yihao03/reminding/internal/database/sqlc"

type MoodReadView struct {
	Mood  int32 `json:"mood"`
	Count int64 `json:"count"`
}

func ToMoodReadViewArray(rows []sqlc.GetMonthlyMoodCountByUserUidRow) []MoodReadView {
	res := make([]MoodReadView, 0, len(rows))
	for _, v := range rows {
		res = append(res, MoodReadView{
			Mood:  v.Mood,
			Count: v.OccurrenceCount,
		})
	}
	return res
}
