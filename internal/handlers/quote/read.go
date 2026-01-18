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
	{Text: "Dementia care is not measured by perfection, but by presence.", Author: ""},
	{Text: "Even when recognition is lost, connection is still possible.", Author: ""},
	{Text: "When memory fades, love becomes the language that remains.", Author: ""},
	{Text: "Caregivers learn to see the person behind the confusion.", Author: ""},
	{Text: "In dementia care, understanding matters more than correction.", Author: ""},
	{Text: "Caregivers carry stories when those they love no longer can.", Author: ""},
	{Text: "Dementia does not erase a person - it changes how we reach them.", Author: ""},
	{Text: "They may forget your name, but they can feel your kindness.", Author: ""},
	{Text: "In a world of forgetting, caregivers become the keepers of identity.", Author: ""},
	{Text: "The heart remembers what the mind forgets.", Author: ""},
	{Text: "Every gentle response preserves dignity.", Author: ""},
	{Text: "Dementia changes the story, but it does not erase the meaning.", Author: ""},
	{Text: "Dementia caregiving is love practiced through repetition.", Author: ""},
	{Text: "To care is to carry another's burden, even when the road is long.", Author: ""},
}

func HandleReadQuote(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	day := time.Now().Day() % 14
	quote := quotes[day]

	api.WriteResponse(quoteview.ToReadQuoteView(quote.Text, quote.Author), w)
	return nil
}
