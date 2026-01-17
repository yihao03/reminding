package quote

import (
	"net/http"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/views/quoteview"
)

var quotes = []struct {
	Text   string
	Author string
}{
	{Text: "You are doing something extraordinary. Every moment of patience, every act of love, every sacrifice you make matters more than you know.", Author: "Maya Angelou"},
	{Text: "Caring for someone with dementia is not about fixing them, it's about being present with them in their reality.", Author: "Teepa Snow"},
	{Text: "In the midst of caring for others, remember to care for yourself. You cannot pour from an empty cup.", Author: "Eleanor Brownn"},
	{Text: "The work you do may not feel heroic, but your daily acts of kindness are changing someone's world.", Author: "Rachel Naomi Remen"},
	{Text: "Love doesn't require memory. Your presence and care speak louder than any words they may have forgotten.", Author: "Lisa Genova"},
	{Text: "Strength isn't about never feeling overwhelmed. It's about continuing to show up even when you do.", Author: "Brené Brown"},
	{Text: "You are not losing the person you love—you are learning new ways to connect with them.", Author: "Judy Cornish"},
}

func HandleReadQuote(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	day := time.Now().Weekday()
	quote := quotes[day]

	api.WriteResponse(quoteview.ToReadQuoteView(quote.Text, quote.Author), w)
	return nil
}
