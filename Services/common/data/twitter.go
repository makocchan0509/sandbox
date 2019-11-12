package data

type StandardSearchRes struct {
	Statuses []struct {
		CreatedAt string `json:"created_at"`
		ID        int64  `json:"id"`
		IDStr     string `json:"id_str"`
		Text      string `json:"text"`
		Truncated bool   `json:"truncated"`
		Entities  struct {
			Hashtags     []interface{} `json:"hashtags"`
			Symbols      []interface{} `json:"symbols"`
			UserMentions []interface{} `json:"user_mentions"`
			Urls         []struct {
				URL         string `json:"url"`
				ExpandedURL string `json:"expanded_url"`
				DisplayURL  string `json:"display_url"`
				Indices     []int  `json:"indices"`
			} `json:"urls"`
		} `json:"entities"`
		Metadata struct {
			IsoLanguageCode string `json:"iso_language_code"`
			ResultType      string `json:"result_type"`
		} `json:"metadata"`
		Source               string      `json:"source"`
		InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
		InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
		InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
		User                 struct {
			ID          int64  `json:"id"`
			IDStr       string `json:"id_str"`
			Name        string `json:"name"`
			ScreenName  string `json:"screen_name"`
			Location    string `json:"location"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Entities    struct {
				URL struct {
					Urls []struct {
						URL         string `json:"url"`
						ExpandedURL string `json:"expanded_url"`
						DisplayURL  string `json:"display_url"`
						Indices     []int  `json:"indices"`
					} `json:"urls"`
				} `json:"url"`
				Description struct {
					Urls []struct {
						URL         string `json:"url"`
						ExpandedURL string `json:"expanded_url"`
						DisplayURL  string `json:"display_url"`
						Indices     []int  `json:"indices"`
					} `json:"urls"`
				} `json:"description"`
			} `json:"entities"`
			Protected                      bool        `json:"protected"`
			FollowersCount                 int         `json:"followers_count"`
			FriendsCount                   int         `json:"friends_count"`
			ListedCount                    int         `json:"listed_count"`
			CreatedAt                      string      `json:"created_at"`
			FavouritesCount                int         `json:"favourites_count"`
			UtcOffset                      interface{} `json:"utc_offset"`
			TimeZone                       interface{} `json:"time_zone"`
			GeoEnabled                     bool        `json:"geo_enabled"`
			Verified                       bool        `json:"verified"`
			StatusesCount                  int         `json:"statuses_count"`
			Lang                           string      `json:"lang"`
			ContributorsEnabled            bool        `json:"contributors_enabled"`
			IsTranslator                   bool        `json:"is_translator"`
			IsTranslationEnabled           interface{} `json:"is_translation_enabled"`
			ProfileBackgroundColor         string      `json:"profile_background_color"`
			ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
			ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
			ProfileBackgroundTile          interface{} `json:"profile_background_tile"`
			ProfileImageURL                string      `json:"profile_image_url"`
			ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
			ProfileBannerURL               string      `json:"profile_banner_url"`
			ProfileLinkColor               string      `json:"profile_link_color"`
			ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
			ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
			ProfileTextColor               string      `json:"profile_text_color"`
			ProfileUseBackgroundImage      interface{} `json:"profile_use_background_image"`
			HasExtendedProfile             interface{} `json:"has_extended_profile"`
			DefaultProfile                 bool        `json:"default_profile"`
			DefaultProfileImage            bool        `json:"default_profile_image"`
			Following                      bool        `json:"following"`
			FollowRequestSent              bool        `json:"follow_request_sent"`
			Notifications                  bool        `json:"notifications"`
			TranslatorType                 string      `json:"translator_type"`
		} `json:"user"`
		Geo               interface{} `json:"geo"`
		Coordinates       interface{} `json:"coordinates"`
		Place             interface{} `json:"place"`
		Contributors      interface{} `json:"contributors"`
		IsQuoteStatus     bool        `json:"is_quote_status"`
		QuotedStatusID    int64       `json:"quoted_status_id,omitempty"`
		QuotedStatusIDStr string      `json:"quoted_status_id_str,omitempty"`
		QuotedStatus      struct {
			CreatedAt string `json:"created_at"`
			ID        int64  `json:"id"`
			IDStr     string `json:"id_str"`
			Text      string `json:"text"`
			Truncated bool   `json:"truncated"`
			Entities  struct {
				Hashtags     []interface{} `json:"hashtags"`
				Symbols      []interface{} `json:"symbols"`
				UserMentions []interface{} `json:"user_mentions"`
				Urls         []struct {
					URL         string `json:"url"`
					ExpandedURL string `json:"expanded_url"`
					DisplayURL  string `json:"display_url"`
					Indices     []int  `json:"indices"`
				} `json:"urls"`
			} `json:"entities"`
			Metadata struct {
				IsoLanguageCode string `json:"iso_language_code"`
				ResultType      string `json:"result_type"`
			} `json:"metadata"`
			Source               string      `json:"source"`
			InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
			InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
			InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
			InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
			InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
			User                 struct {
				ID          int    `json:"id"`
				IDStr       string `json:"id_str"`
				Name        string `json:"name"`
				ScreenName  string `json:"screen_name"`
				Location    string `json:"location"`
				Description string `json:"description"`
				URL         string `json:"url"`
				Entities    struct {
					URL struct {
						Urls []struct {
							URL         string `json:"url"`
							ExpandedURL string `json:"expanded_url"`
							DisplayURL  string `json:"display_url"`
							Indices     []int  `json:"indices"`
						} `json:"urls"`
					} `json:"url"`
					Description struct {
						Urls []interface{} `json:"urls"`
					} `json:"description"`
				} `json:"entities"`
				Protected                      bool        `json:"protected"`
				FollowersCount                 int         `json:"followers_count"`
				FriendsCount                   int         `json:"friends_count"`
				ListedCount                    int         `json:"listed_count"`
				CreatedAt                      string      `json:"created_at"`
				FavouritesCount                int         `json:"favourites_count"`
				UtcOffset                      interface{} `json:"utc_offset"`
				TimeZone                       interface{} `json:"time_zone"`
				GeoEnabled                     bool        `json:"geo_enabled"`
				Verified                       bool        `json:"verified"`
				StatusesCount                  int         `json:"statuses_count"`
				Lang                           string      `json:"lang"`
				ContributorsEnabled            bool        `json:"contributors_enabled"`
				IsTranslator                   bool        `json:"is_translator"`
				IsTranslationEnabled           interface{} `json:"is_translation_enabled"`
				ProfileBackgroundColor         string      `json:"profile_background_color"`
				ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
				ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
				ProfileBackgroundTile          interface{} `json:"profile_background_tile"`
				ProfileImageURL                string      `json:"profile_image_url"`
				ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
				ProfileBannerURL               string      `json:"profile_banner_url"`
				ProfileLinkColor               string      `json:"profile_link_color"`
				ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
				ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
				ProfileTextColor               string      `json:"profile_text_color"`
				ProfileUseBackgroundImage      interface{} `json:"profile_use_background_image"`
				HasExtendedProfile             interface{} `json:"has_extended_profile"`
				DefaultProfile                 bool        `json:"default_profile"`
				DefaultProfileImage            bool        `json:"default_profile_image"`
				Following                      bool        `json:"following"`
				FollowRequestSent              bool        `json:"follow_request_sent"`
				Notifications                  bool        `json:"notifications"`
				TranslatorType                 string      `json:"translator_type"`
			} `json:"user"`
			Geo               interface{} `json:"geo"`
			Coordinates       interface{} `json:"coordinates"`
			Place             interface{} `json:"place"`
			Contributors      interface{} `json:"contributors"`
			IsQuoteStatus     bool        `json:"is_quote_status"`
			RetweetCount      int         `json:"retweet_count"`
			FavoriteCount     int         `json:"favorite_count"`
			Favorited         bool        `json:"favorited"`
			Retweeted         bool        `json:"retweeted"`
			PossiblySensitive bool        `json:"possibly_sensitive"`
			Lang              string      `json:"lang"`
		} `json:"quoted_status,omitempty"`
		RetweetCount      int    `json:"retweet_count"`
		FavoriteCount     int    `json:"favorite_count"`
		Favorited         bool   `json:"favorited"`
		Retweeted         bool   `json:"retweeted"`
		PossiblySensitive bool   `json:"possibly_sensitive"`
		Lang              string `json:"lang"`
	} `json:"statuses"`
	SearchMetadata struct {
		CompletedIn float64 `json:"completed_in"`
		MaxID       int64   `json:"max_id"`
		MaxIDStr    string  `json:"max_id_str"`
		NextResults string  `json:"next_results"`
		Query       string  `json:"query"`
		RefreshURL  string  `json:"refresh_url"`
		Count       int     `json:"count"`
		SinceID     int     `json:"since_id"`
		SinceIDStr  string  `json:"since_id_str"`
	} `json:"search_metadata"`
}

type TweetServiceReq struct {
	Query      string `json:"query"`
	OrderCount int    `json:"ordercount"`
	MaxResults string `json:"maxresults"`
	SearchMode string `json:"searchmode"`
}

type TweetServiceRes struct {
	Result   string `json:"result"`
	GetCount int    `json:"getcount"`
}

type PremiumSearchReq struct {
	Query      string `json:"query"`
	MaxResults string `json:"maxResults"`
}

type PremiumSearchNextReq struct {
	Query      string `json:"query"`
	MaxResults string `json:"maxResults"`
	Next       string `json:"next"`
}

type PremiumSearchRes struct {
	Results           []TweetList `json:"results"`
	Next              string      `json:"next"`
	RequestParameters struct {
		MaxResults int    `json:"maxResults"`
		FromDate   string `json:"fromDate"`
		ToDate     string `json:"toDate"`
	} `json:"requestParameters"`
}

type TweetList struct {
	CreatedAt            string      `json:"created_at"`
	ID                   int64       `json:"id"`
	IDStr                string      `json:"id_str"`
	Text                 string      `json:"text"`
	Source               string      `json:"source"`
	Truncated            bool        `json:"truncated"`
	InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
	InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
	User                 struct {
		ID                             int64       `json:"id"`
		IDStr                          string      `json:"id_str"`
		Name                           string      `json:"name"`
		ScreenName                     string      `json:"screen_name"`
		Location                       string      `json:"location"`
		URL                            interface{} `json:"url"`
		Description                    interface{} `json:"description"`
		TranslatorType                 string      `json:"translator_type"`
		Protected                      bool        `json:"protected"`
		Verified                       bool        `json:"verified"`
		FollowersCount                 int         `json:"followers_count"`
		FriendsCount                   int         `json:"friends_count"`
		ListedCount                    int         `json:"listed_count"`
		FavouritesCount                int         `json:"favourites_count"`
		StatusesCount                  int         `json:"statuses_count"`
		CreatedAt                      string      `json:"created_at"`
		UtcOffset                      interface{} `json:"utc_offset"`
		TimeZone                       interface{} `json:"time_zone"`
		GeoEnabled                     bool        `json:"geo_enabled"`
		Lang                           interface{} `json:"lang"`
		ContributorsEnabled            bool        `json:"contributors_enabled"`
		IsTranslator                   bool        `json:"is_translator"`
		ProfileBackgroundColor         string      `json:"profile_background_color"`
		ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool        `json:"profile_background_tile"`
		ProfileLinkColor               string      `json:"profile_link_color"`
		ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string      `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
		ProfileImageURL                string      `json:"profile_image_url"`
		ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
		ProfileBannerURL               string      `json:"profile_banner_url"`
		DefaultProfile                 bool        `json:"default_profile"`
		DefaultProfileImage            bool        `json:"default_profile_image"`
		Following                      interface{} `json:"following"`
		FollowRequestSent              interface{} `json:"follow_request_sent"`
		Notifications                  interface{} `json:"notifications"`
	} `json:"user"`
	Geo             interface{} `json:"geo"`
	Coordinates     interface{} `json:"coordinates"`
	Place           interface{} `json:"place"`
	Contributors    interface{} `json:"contributors"`
	RetweetedStatus struct {
		CreatedAt            string      `json:"created_at"`
		ID                   int64       `json:"id"`
		IDStr                string      `json:"id_str"`
		Text                 string      `json:"text"`
		DisplayTextRange     []int       `json:"display_text_range"`
		Source               string      `json:"source"`
		Truncated            bool        `json:"truncated"`
		InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
		InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
		InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
		User                 struct {
			ID                             int64       `json:"id"`
			IDStr                          string      `json:"id_str"`
			Name                           string      `json:"name"`
			ScreenName                     string      `json:"screen_name"`
			Location                       interface{} `json:"location"`
			URL                            string      `json:"url"`
			Description                    string      `json:"description"`
			TranslatorType                 string      `json:"translator_type"`
			Protected                      bool        `json:"protected"`
			Verified                       bool        `json:"verified"`
			FollowersCount                 int         `json:"followers_count"`
			FriendsCount                   int         `json:"friends_count"`
			ListedCount                    int         `json:"listed_count"`
			FavouritesCount                int         `json:"favourites_count"`
			StatusesCount                  int         `json:"statuses_count"`
			CreatedAt                      string      `json:"created_at"`
			UtcOffset                      interface{} `json:"utc_offset"`
			TimeZone                       interface{} `json:"time_zone"`
			GeoEnabled                     bool        `json:"geo_enabled"`
			Lang                           interface{} `json:"lang"`
			ContributorsEnabled            bool        `json:"contributors_enabled"`
			IsTranslator                   bool        `json:"is_translator"`
			ProfileBackgroundColor         string      `json:"profile_background_color"`
			ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
			ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
			ProfileBackgroundTile          bool        `json:"profile_background_tile"`
			ProfileLinkColor               string      `json:"profile_link_color"`
			ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
			ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
			ProfileTextColor               string      `json:"profile_text_color"`
			ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
			ProfileImageURL                string      `json:"profile_image_url"`
			ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
			ProfileBannerURL               string      `json:"profile_banner_url"`
			DefaultProfile                 bool        `json:"default_profile"`
			DefaultProfileImage            bool        `json:"default_profile_image"`
			Following                      interface{} `json:"following"`
			FollowRequestSent              interface{} `json:"follow_request_sent"`
			Notifications                  interface{} `json:"notifications"`
		} `json:"user"`
		Geo           interface{} `json:"geo"`
		Coordinates   interface{} `json:"coordinates"`
		Place         interface{} `json:"place"`
		Contributors  interface{} `json:"contributors"`
		IsQuoteStatus bool        `json:"is_quote_status"`
		QuoteCount    int         `json:"quote_count"`
		ReplyCount    int         `json:"reply_count"`
		RetweetCount  int         `json:"retweet_count"`
		FavoriteCount int         `json:"favorite_count"`
		Entities      struct {
			Hashtags []struct {
				Text    string `json:"text"`
				Indices []int  `json:"indices"`
			} `json:"hashtags"`
			Urls         []interface{} `json:"urls"`
			UserMentions []interface{} `json:"user_mentions"`
			Symbols      []interface{} `json:"symbols"`
			Media        []struct {
				ID                  int64  `json:"id"`
				IDStr               string `json:"id_str"`
				Indices             []int  `json:"indices"`
				AdditionalMediaInfo struct {
					Monetizable bool `json:"monetizable"`
				} `json:"additional_media_info"`
				MediaURL      string `json:"media_url"`
				MediaURLHTTPS string `json:"media_url_https"`
				URL           string `json:"url"`
				DisplayURL    string `json:"display_url"`
				ExpandedURL   string `json:"expanded_url"`
				Type          string `json:"type"`
				Sizes         struct {
					Thumb struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"thumb"`
					Medium struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"medium"`
					Small struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"small"`
					Large struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"large"`
				} `json:"sizes"`
			} `json:"media"`
		} `json:"entities"`
		ExtendedEntities struct {
			Media []struct {
				ID                  int64  `json:"id"`
				IDStr               string `json:"id_str"`
				Indices             []int  `json:"indices"`
				AdditionalMediaInfo struct {
					Monetizable bool `json:"monetizable"`
				} `json:"additional_media_info"`
				MediaURL      string `json:"media_url"`
				MediaURLHTTPS string `json:"media_url_https"`
				URL           string `json:"url"`
				DisplayURL    string `json:"display_url"`
				ExpandedURL   string `json:"expanded_url"`
				Type          string `json:"type"`
				VideoInfo     struct {
					AspectRatio    []int `json:"aspect_ratio"`
					DurationMillis int   `json:"duration_millis"`
					Variants       []struct {
						ContentType string `json:"content_type"`
						URL         string `json:"url"`
						Bitrate     int    `json:"bitrate,omitempty"`
					} `json:"variants"`
				} `json:"video_info"`
				Sizes struct {
					Thumb struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"thumb"`
					Medium struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"medium"`
					Small struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"small"`
					Large struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"large"`
				} `json:"sizes"`
			} `json:"media"`
		} `json:"extended_entities"`
		Favorited         bool   `json:"favorited"`
		Retweeted         bool   `json:"retweeted"`
		PossiblySensitive bool   `json:"possibly_sensitive"`
		FilterLevel       string `json:"filter_level"`
		Lang              string `json:"lang"`
	} `json:"retweeted_status"`
	IsQuoteStatus bool `json:"is_quote_status"`
	QuoteCount    int  `json:"quote_count"`
	ReplyCount    int  `json:"reply_count"`
	RetweetCount  int  `json:"retweet_count"`
	FavoriteCount int  `json:"favorite_count"`
	Entities      struct {
		Hashtags []struct {
			Text    string `json:"text"`
			Indices []int  `json:"indices"`
		} `json:"hashtags"`
		Urls         []interface{} `json:"urls"`
		UserMentions []struct {
			ScreenName string `json:"screen_name"`
			Name       string `json:"name"`
			ID         int64  `json:"id"`
			IDStr      string `json:"id_str"`
			Indices    []int  `json:"indices"`
		} `json:"user_mentions"`
		Symbols []interface{} `json:"symbols"`
		Media   []struct {
			ID                  int64  `json:"id"`
			IDStr               string `json:"id_str"`
			Indices             []int  `json:"indices"`
			AdditionalMediaInfo struct {
				Monetizable bool `json:"monetizable"`
			} `json:"additional_media_info"`
			MediaURL      string `json:"media_url"`
			MediaURLHTTPS string `json:"media_url_https"`
			URL           string `json:"url"`
			DisplayURL    string `json:"display_url"`
			ExpandedURL   string `json:"expanded_url"`
			Type          string `json:"type"`
			Sizes         struct {
				Thumb struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"thumb"`
				Medium struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"medium"`
				Small struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"small"`
				Large struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"large"`
			} `json:"sizes"`
			SourceStatusID    int64  `json:"source_status_id"`
			SourceStatusIDStr string `json:"source_status_id_str"`
			SourceUserID      int64  `json:"source_user_id"`
			SourceUserIDStr   string `json:"source_user_id_str"`
		} `json:"media"`
	} `json:"entities"`
	ExtendedEntities struct {
		Media []struct {
			ID                  int64  `json:"id"`
			IDStr               string `json:"id_str"`
			Indices             []int  `json:"indices"`
			AdditionalMediaInfo struct {
				Monetizable bool `json:"monetizable"`
			} `json:"additional_media_info"`
			MediaURL      string `json:"media_url"`
			MediaURLHTTPS string `json:"media_url_https"`
			URL           string `json:"url"`
			DisplayURL    string `json:"display_url"`
			ExpandedURL   string `json:"expanded_url"`
			Type          string `json:"type"`
			VideoInfo     struct {
				AspectRatio    []int `json:"aspect_ratio"`
				DurationMillis int   `json:"duration_millis"`
				Variants       []struct {
					ContentType string `json:"content_type"`
					URL         string `json:"url"`
					Bitrate     int    `json:"bitrate,omitempty"`
				} `json:"variants"`
			} `json:"video_info"`
			Sizes struct {
				Thumb struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"thumb"`
				Medium struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"medium"`
				Small struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"small"`
				Large struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"large"`
			} `json:"sizes"`
			SourceStatusID    int64  `json:"source_status_id"`
			SourceStatusIDStr string `json:"source_status_id_str"`
			SourceUserID      int64  `json:"source_user_id"`
			SourceUserIDStr   string `json:"source_user_id_str"`
		} `json:"media"`
	} `json:"extended_entities"`
	Favorited         bool   `json:"favorited"`
	Retweeted         bool   `json:"retweeted"`
	PossiblySensitive bool   `json:"possibly_sensitive"`
	FilterLevel       string `json:"filter_level"`
	Lang              string `json:"lang"`
	MatchingRules     []struct {
		Tag interface{} `json:"tag"`
	} `json:"matching_rules"`
}
