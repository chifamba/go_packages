package socmedia

type (
	CleanMsg struct {
		Lang              string     `json:"lang"`
		MediaSource       string     `json:"media_source"`
		Text              string     `json:"text"`
		FullText          string     `json:"full_text"`
		User              string     `json:"user"`
		SourceSearchTopic string     `json:"source_topic"`
		CleanCorpus       string     `json:"clean_corpus"`
		BagOfWords        []string   `json:"bag_of_words"`
		BagOfNeatWords     []string   `json:"bag_of_neat_words"`
		Sentences         []Sentence `json:"sentences"`
	}
	Sentence struct {
		Sentence              string  `json:"sentence"`
		SentimentScore        float64 `json:"sentiment_score"`
		SentimentSubjectivity float64 `json:"sentiment_subjectivity"`
	}
)
