package quoteview

type ReadQuoteView struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func ToReadQuoteView(Text string, Author string) *ReadQuoteView {
	return &ReadQuoteView{Text, Author}
}
