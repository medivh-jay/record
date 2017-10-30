package pubg

import "encoding/json"

type Season struct {
	Season [5]struct {
		Season        string `json:"season"`
		DisplayYear   string `json:"display_year"`
		DisplaySeason string `json:"display_season"`
	} `json:"season"`
}

func (season *Season) GetSeason() []byte {
	sea := [5]struct {
		Season        string `json:"season"`
		DisplayYear   string `json:"display_year"`
		DisplaySeason string `json:"display_season"`
	}{}
	season.Season = sea

	season.Season[0].Season = "2017-pre5"
	season.Season[0].DisplayYear = "2017"
	season.Season[0].DisplaySeason = "第5赛季"

	season.Season[1].Season = "2017-pre4"
	season.Season[1].DisplayYear = "2017"
	season.Season[1].DisplaySeason = "第4赛季"

	season.Season[2].Season = "2017-pre3"
	season.Season[2].DisplayYear = "2017"
	season.Season[2].DisplaySeason = "第3赛季"

	season.Season[3].Season = "2017-pre2"
	season.Season[3].DisplayYear = "2017"
	season.Season[3].DisplaySeason = "第2赛季"

	season.Season[4].Season = "2017-pre1"
	season.Season[4].DisplayYear = "2017"
	season.Season[4].DisplaySeason = "第1赛季"

	js, _ := json.Marshal(season)
	return js
}

type UserInfo struct {
	AccountID     string `json:"account_id"`
	Avatar        string `json:"avatar"`
	DefaultSeason string `json:"default_season"`
	Nickname      string `json:"nickname"`
	Platform      int    `json:"platform"`
	SelectRegion  string `json:"select_region"`
	SteamID       int    `json:"steam_id"`
}

func (userInfo *UserInfo) ToJSON(playerData *PlayerData) []byte {
	userInfo.AccountID = playerData.AccountID
	userInfo.Avatar = playerData.Avatar
	userInfo.DefaultSeason = playerData.DefaultSeason
	userInfo.Nickname = playerData.PlayerName
	userInfo.Platform = playerData.Platform
	userInfo.SelectRegion = playerData.SelectedRegion
	userInfo.SteamID = playerData.SteamID
	js, _ := json.Marshal(userInfo)
	return js
}
