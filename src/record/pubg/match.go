package pubg

type PlayerData struct {
	AccountID     string `json:"AccountId" bson:"_id"`
	Avatar        string `json:"Avatar" bson:"avatar"`
	LastUpdated   string `json:"LastUpdated" bson:"last_updated"`
	LifeTimeStats []struct {
		Key   string `json:"Key" bson:"key"`
		Value string `json:"Value" bson:"value"`
	} `json:"LifeTimeStats" bson:"life_time_stats"`
	LiveTracking []struct {
		Date         string      `json:"Date" bson:"date"`
		Delta        float64     `json:"Delta" bson:"delta"`
		Match        int         `json:"Match" bson:"match"`
		MatchDisplay string      `json:"MatchDisplay" bson:"match_display"`
		Region       string      `json:"Region" bson:"region"`
		RegionID     int         `json:"RegionId" bson:"region_id"`
		Season       int         `json:"Season" bson:"season"`
		Value        float64     `json:"Value" bson:"value"`
		Message      interface{} `json:"message" bson:"message"`
	} `json:"LiveTracking" bson:"live_tracking"`
	MatchHistory []struct {
		Assists              int     `json:"Assists" bson:"assists"`
		Damage               float64 `json:"Damage" bson:"damage"`
		Headshots            float64 `json:"Headshots" bson:"headshots"`
		ID                   int     `json:"Id" bson:"id"`
		Kd                   float64 `json:"Kd" bson:"kd"`
		KillRank             float64 `json:"KillRank" bson:"kill_rank"`
		KillRating           float64 `json:"KillRating" bson:"kill_rating"`
		KillRatingChange     float64 `json:"KillRatingChange" bson:"kill_rating_change"`
		KillRatingRankChange float64 `json:"KillRatingRankChange" bson:"kill_rating_rank_rhange"`
		Kills                float64 `json:"Kills" bson:"kills"`
		Match                float64 `json:"Match" bson:"match"`
		MatchDisplay         string  `json:"MatchDisplay" bson:"match_display"`
		MoveDistance         float64 `json:"MoveDistance" bson:"move_distance"`
		Rating               float64 `json:"Rating" bson:"rating"`
		RatingChange         float64 `json:"RatingChange" bson:"rating_change"`
		RatingRank           float64 `json:"RatingRank" bson:"rating_rank"`
		RatingRankChange     float64 `json:"RatingRankChange" bson:"rating_rank_change"`
		Region               float64 `json:"Region" bson:"region"`
		RegionDisplay        string  `json:"RegionDisplay" bson:"region_display"`
		Rounds               float64 `json:"Rounds" bson:"rounds"`
		Season               float64 `json:"Season" bson:"season"`
		SeasonDisplay        string  `json:"SeasonDisplay" bson:"season_display"`
		TimeSurvived         float64 `json:"TimeSurvived" bson:"time_survived"`
		Top10                int     `json:"Top10" bson:"top10"`
		Updated              string  `json:"Updated" bson:"updated"`
		UpdatedJS            string  `json:"UpdatedJS" bson:"updated_js"`
		WinRank              float64 `json:"WinRank" bson:"win_rank"`
		WinRating            float64 `json:"WinRating" bson:"win_rating"`
		WinRatingChange      float64 `json:"WinRatingChange" bson:"win_rating_change"`
		WinRatingRankChange  float64 `json:"WinRatingRankChange" bson:"win_rating_rank_change"`
		Wins                 int     `json:"Wins" bson:"wins"`
	} `json:"MatchHistory" bson:"match_history"`
	Platform      int    `json:"Platform" bson:"platform"`
	PlayerName    string `json:"PlayerName" bson:"player_name"`
	PubgTrackerID int    `json:"PubgTrackerId" bson:"pubg_tracker_id"`
	Stats         []struct {
		Match  string `json:"Match" bson:"match"`
		Region string `json:"Region" bson:"region"`
		Season string `json:"Season" bson:"season"`
		Stats  []struct {
			ValueDec     float64     `json:"ValueDec" bson:"value_dec"`
			ValueInt     interface{} `json:"ValueInt" bson:"value_int"`
			Category     string      `json:"category" bson:"category"`
			DisplayValue string      `json:"displayValue" bson:"display_value"`
			Field        string      `json:"field" bson:"field"`
			Label        string      `json:"label" bson:"label"`
			Percentile   interface{} `json:"percentile" bson:"percentile"`
			Rank         interface{} `json:"rank" bson:"rank"`
			Value        string      `json:"value" bson:"value"`
		} `json:"Stats" bson:"stats"`
	} `json:"Stats" bson:"stats"`
	SteamID        int         `json:"SteamId" bson:"steam_id"`
	SteamName      interface{} `json:"SteamName" bson:"steam_name"`
	TimePlayed     int         `json:"TimePlayed" bson:"time_played"`
	Twitch         interface{} `json:"Twitch" bson:"twitch"`
	UserID         interface{} `json:"UserId" bson:"userId"`
	DefaultSeason  string      `json:"defaultSeason" bson:"default_season"`
	SelectedMatch  interface{} `json:"selectedMatch" bson:"selected_match"`
	SelectedRegion string      `json:"selectedRegion" bson:"selected_region"`
	SelectedSeason string      `json:"selectedSeason" bson:"selected_season"`
	CreatedAt      int64       `bson:"created_at"`
	UpdatedAt      int64       `bson:"updated_at"`
}
