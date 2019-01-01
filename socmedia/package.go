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
	TwMsg struct {
		CreatedAt            string `json:"created_at"`
		FavoriteCount        int    `json:"favorite_count"`
		ID                   int64  `json:"id"`
		IDStr                string `json:"id_str"`
		InReplyToScreenName  string `json:"in_reply_to_screen_name"`
		InReplyToStatusID    int64  `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr string `json:"in_reply_to_status_id_str"`
		InReplyToUserID      int64  `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   string `json:"in_reply_to_user_id_str"`
		Lang                 string `json:"lang"`
		PossiblySensitive    bool   `json:"possibly_sensitive"`
		RetweetCount         int    `json:"retweet_count"`
		Source               string `json:"source"`
		Text                 string `json:"text"`
		FullText             string `json:"full_text"`
		User                 TwUser `json:"user"`
		QuotedStatusID       int64  `json:"quoted_status_id"`
		QuotedStatusIDStr    string `json:"quoted_status_id_str"`
		SourceSearchTopic    string `json:"source_topic"`
	}
	TwUser struct {
		CreatedAt       string `json:"created_at"`
		Description     string `json:"description"`
		Email           string `json:"email"`
		FavouritesCount int    `json:"favourites_count"`
		FollowersCount  int    `json:"followers_count"`
		FriendsCount    int    `json:"friends_count"`
		ID              int64  `json:"id"`
		IDStr           string `json:"id_str"`
		Lang            string `json:"lang"`
		Location        string `json:"location"`
		Name            string `json:"name"`
		Protected       bool   `json:"protected"`
		ScreenName      string `json:"screen_name"`
		Timezone        string `json:"time_zone"`
		URL             string `json:"url"`
		UtcOffset       int    `json:"utc_offset"`
		Verified        bool   `json:"verified"`
	}
)
