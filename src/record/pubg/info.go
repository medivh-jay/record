package pubg

import (
	"encoding/json"
	"reflect"
)

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

type Performance struct {
	KillDeathRatio string `json:"KillDeathRatio"`
	Losses         string `json:"Losses"`
	RoundsPlayed   string `json:"RoundsPlayed"`
	TimeSurvived   string `json:"TimeSurvived"`
	Top10Ratio     string `json:"Top10Ratio"`
	Top10s         string `json:"Top10s"`
	WinPoints      string `json:"WinPoints"`
	WinRatio       string `json:"WinRatio"`
	WinTop10Ratio  string `json:"WinTop10Ratio"`
	Wins           string `json:"Wins"`
}

type Combat struct {
	Assists           string `json:"Assists"`
	DailyKills        string `json:"DailyKills"`
	HeadshotKillRatio string `json:"HeadshotKillRatio"`
	HeadshotKills     string `json:"HeadshotKills"`
	Kills             string `json:"Kills"`
	MaxKillStreaks    string `json:"MaxKillStreaks"`
	RoadKills         string `json:"RoadKills"`
	RoundMostKills    string `json:"RoundMostKills"`
	Suicides          string `json:"Suicides"`
	TeamKills         string `json:"TeamKills"`
	VehicleDestroys   string `json:"VehicleDestroys"`
	WeaponAcquired    string `json:"WeaponAcquired"`
	WeeklyKills       string `json:"WeeklyKills"`
}

type Distance struct {
	AvgRideDistance string `json:"AvgRideDistance"`
	AvgWalkDistance string `json:"AvgWalkDistance"`
	LongestKill     string `json:"LongestKill"`
	MoveDistance    string `json:"MoveDistance"`
	RideDistance    string `json:"RideDistance"`
	WalkDistance    string `json:"WalkDistance"`
}

type PerGame struct {
	DamagePg        string `json:"DamagePg"`
	HeadshotKillsPg string `json:"HeadshotKillsPg"`
	HealsPg         string `json:"HealsPg"`
	KillsPg         string `json:"KillsPg"`
	MoveDistancePg  string `json:"MoveDistancePg"`
	RevivesPg       string `json:"RevivesPg"`
	RoadKillsPg     string `json:"RoadKillsPg"`
	TeamKillsPg     string `json:"TeamKillsPg"`
	TimeSurvivedPg  string `json:"TimeSurvivedPg"`
	Top10sPg        string `json:"Top10sPg"`
}

type SkillRating struct {
	BestRank   string `json:"BestRank"`
	BestRating string `json:"BestRating"`
	Rating     string `json:"Rating"`
}

type Support struct {
	Boosts      string `json:"Boosts"`
	DBNOs       string `json:"DBNOs"`
	DamageDealt string `json:"DamageDealt"`
	Heals       string `json:"Heals"`
	Revives     string `json:"Revives"`
}

type Survival struct {
	AvgSurvivalTime     string `json:"AvgSurvivalTime"`
	Days                string `json:"Days"`
	LongestTimeSurvived string `json:"LongestTimeSurvived"`
	MostSurvivalTime    string `json:"MostSurvivalTime"`
}

type Match struct {
	AccountID string `json:"account_id"`
	Avatar    string `json:"avatar"`
	Match     string `json:"match"`
	Nickname  string `json:"nickname"`
	Region    string `json:"region"`
	Season    string `json:"season"`
	Stats     struct {
		Combat      *Combat      `json:"Combat"`
		Distance    *Distance    `json:"Distance"`
		PerGame     *PerGame     `json:"PerGame"`
		Performance *Performance `json:"Performance"`
		SkillRating *SkillRating `json:"SkillRating"`
		Support     *Support     `json:"Support"`
		Survival    *Survival    `json:"Survival"`
	} `json:"stats"`
}

func (match *Match) GetStats(matchType string, region string, season string, playerData *PlayerData) []byte {
	match.AccountID = playerData.AccountID
	match.Avatar = playerData.Avatar
	match.Match = matchType
	match.Nickname = playerData.PlayerName
	match.Region = region
	match.Season = season
	stats := playerData.Stats
	match.Stats.Performance = &Performance{}
	match.Stats.SkillRating = &SkillRating{}
	match.Stats.PerGame = &PerGame{}
	match.Stats.Combat = &Combat{}
	match.Stats.Survival = &Survival{}
	match.Stats.Distance = &Distance{}
	match.Stats.Support = &Support{}
	performance := reflect.ValueOf(match.Stats.Performance).Elem()
	skillRating := reflect.ValueOf(match.Stats.SkillRating).Elem()
	perGame := reflect.ValueOf(match.Stats.PerGame).Elem()
	combat := reflect.ValueOf(match.Stats.Combat).Elem()
	survival := reflect.ValueOf(match.Stats.Survival).Elem()
	distance := reflect.ValueOf(match.Stats.Distance).Elem()
	support := reflect.ValueOf(match.Stats.Support).Elem()
	for _, stat := range stats {
		if stat.Match == matchType && stat.Region == region && stat.Season == season {
			for _, sonStat := range stat.Stats {
				if sonStat.Category == "Performance" {
					performance.FieldByName(sonStat.Field).Set(reflect.ValueOf(sonStat.Value))
				}
				if sonStat.Category == "Skill Rating" {
					skillRating.FieldByName(sonStat.Field).Set(reflect.ValueOf(sonStat.Value))
				}
				if sonStat.Category == "Per Game" {
					perGame.FieldByName(sonStat.Field).Set(reflect.ValueOf(sonStat.Value))
				}
				if sonStat.Category == "Combat" {
					combat.FieldByName(sonStat.Field).Set(reflect.ValueOf(sonStat.Value))
				}
				if sonStat.Category == "Survival" {
					survival.FieldByName(sonStat.Field).Set(reflect.ValueOf(sonStat.Value))
				}
				if sonStat.Category == "Distance" {
					distance.FieldByName(sonStat.Field).Set(reflect.ValueOf(sonStat.Value))
				}
				if sonStat.Category == "Support" {
					support.FieldByName(sonStat.Field).Set(reflect.ValueOf(sonStat.Value))
				}
			}
		}
	}
	js, _ := json.Marshal(match)
	return js
}

type RegionInfo struct {
	Season string `json:"season"`
	Data   *struct {
		Duo   []string `json:"duo"`
		Solo  []string `json:"solo"`
		Squad []string `json:"squad"`
	} `json:"data"`
}

type Regions struct {
	RegionInfo map[string]*RegionInfo `json:"region_info"`
}

func (regions *Regions) GetUserRegion(playerData *PlayerData) []byte {
	regions.RegionInfo = make(map[string]*RegionInfo)
	stats := playerData.Stats
	for _, statInfo := range stats {
		_, ok := regions.RegionInfo[statInfo.Season]
		if ok == false {
			regions.RegionInfo[statInfo.Season] = &RegionInfo{Season: statInfo.Season, Data: &struct {
				Duo   []string `json:"duo"`
				Solo  []string `json:"solo"`
				Squad []string `json:"squad"`
			}{Duo: []string{}, Solo: []string{}, Squad: []string{}}}
			regions.RegionInfo[statInfo.Season].Season = statInfo.Season
		}
		if statInfo.Region != "agg" {
			if statInfo.Match == "solo" {
				regions.RegionInfo[statInfo.Season].Data.Solo = append(regions.RegionInfo[statInfo.Season].Data.Solo, statInfo.Region)
			}
			if statInfo.Match == "duo" {
				regions.RegionInfo[statInfo.Season].Data.Duo = append(regions.RegionInfo[statInfo.Season].Data.Duo, statInfo.Region)
			}
			if statInfo.Match == "squad" {
				regions.RegionInfo[statInfo.Season].Data.Squad = append(regions.RegionInfo[statInfo.Season].Data.Squad, statInfo.Region)
			}
		}
	}
	js, _ := json.Marshal(regions)
	return js
}

type History struct {
	Histories []interface{} `json:"histories"`
}

func (history *History) GetHistory(playerData *PlayerData) []byte {
	history.Histories = []interface{}{}
	for _, stat := range playerData.MatchHistory {
		history.Histories = append(history.Histories, stat)
	}

	js, _ := json.Marshal(history)
	return js
}
