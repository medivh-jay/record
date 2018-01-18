package kotk

type GameUserIdGetPage struct {
	ErrorsList     interface{} `json:"errorsList"`
	Success        bool        `json:"success"`
	SuccessPayload struct {
		PageNumber int `json:"pageNumber"`
		Rows       []struct {
			Detail struct {
				TopMatches []struct {
					Assists     string `json:"assists"`
					Kills       string `json:"kills"`
					Rank        string `json:"rank"`
					Score       string `json:"score"`
					TeamAssists string `json:"team_assists"`
					TeamKills   string `json:"team_kills"`
					TeamRank    string `json:"team_rank"`
				} `json:"top_matches"`
			} `json:"detail"`
			Position string `json:"position"`
			Values   struct {
				GameUserID      string `json:"game_user_id"`
				KillsPerMatch   string `json:"kills_per_match"`
				SubTier         string `json:"subtier"`
				Tier            string `json:"tier"`
				Top10TotalScore string `json:"top_10_total_score"`
				TopKills        string `json:"top_kills"`
				TopScore        string `json:"top_score"`
				TopTensPerMatch string `json:"top_tens_per_match"`
				TotalKills      string `json:"total_kills"`
				TotalMatches    string `json:"total_matches"`
				TotalTopTens    string `json:"total_top_tens"`
				TotalWins       string `json:"total_wins"`
				UserName        string `json:"user_name"`
				WinsPerMatch    string `json:"wins_per_match"`
			} `json:"values"`
		} `json:"rows"`
		TotalPages int `json:"totalPages"`
	} `json:"successPayload"`
}

type NameSearch struct {
	ErrorsList     interface{} `json:"errorsList"`
	Success        bool        `json:"success"`
	SuccessPayload []struct {
		GameUserID string `json:"game_user_id"`
		UserName   string `json:"user_name"`
	} `json:"successPayload"`
}

// 数据库结构
type PlayerData struct {
	Id              string `json:"_id" bson:"_id"`
	GameUserID      string `json:"game_user_id" bson:"game_user_id"`
	KillsPerMatch   string `json:"kills_per_match" bson:"kills_per_match"`
	Position        string `json:"position" bson:"position"`
	Region          string `json:"region" bson:"region"`
	Season          string `json:"season" bson:"season"`
	SubTier         string `json:"subtier" bson:"subtier"`
	TeamModel       string `json:"team_model" bson:"team_model"`
	Tier            string `json:"tier" bson:"tier"`
	Top10TotalScore string `json:"top_10_total_score" bson:"top_10_total_score"`
	TopKills        string `json:"top_kills" bson:"top_kills"`
	TopMatches      []struct {
		Assists     string `json:"assists" bson:"assists"`
		Kills       string `json:"kills" bson:"kills"`
		Rank        string `json:"rank" bson:"rank"`
		Score       string `json:"score" bson:"score"`
		TeamAssists string `json:"team_assists" bson:"team_assists"`
		TeamKills   string `json:"team_kills" bson:"team_kills"`
		TeamRank    string `json:"team_rank" bson:"team_rank"`
	} `json:"top_matches" bson:"top_matches"`
	TopScore        string `json:"top_score" bson:"top_score"`
	TopTensPerMatch string `json:"top_tens_per_match" bson:"top_tens_per_match"`
	TotalKills      string `json:"total_kills" bson:"total_kills"`
	TotalMatches    string `json:"total_matches" bson:"total_matches"`
	TotalTopTens    string `json:"total_top_tens" bson:"total_top_tens"`
	TotalWins       string `json:"total_wins" bson:"total_wins"`
	UserName        string `json:"user_name" bson:"user_name"`
	ITime           int    `json:"itime" bson:"itime"`
	UTime           int    `json:"utime" bson:"utime"`
	WinsPerMatch    string `json:"wins_per_match" bson:"wins_per_match"`
}

// 搜索条件
type Search struct {
	UserName  string `bson:"user_name"`
	Region    string `bson:"region"`
	Season    string `bson:"season"`
	TeamModel string `bson:"team_model"`
}
