package pubg

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"record/log"
	"reflect"
	"strings"
)

type Season struct {
	jsonFile string

	Season []struct {
		Season        string `json:"season"`
		DisplayYear   string `json:"display_year"`
		DisplaySeason string `json:"display_season"`
	} `json:"season"`
}

func (season *Season) Load() {
	reader, err := os.Open(getCurrentDirectory() + "/season.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	seasons, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(seasons, season)
	reader.Close()
}

func (season *Season) GetSeason() []byte {
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

// 用户总览信息的结构
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

func (match *Match) ToJSON() []byte {
	js, _ := json.Marshal(match)
	return js
}

// 生成用户总览的信息
func (match *Match) GetStats(matchType string, region string, season string, playerData *PlayerData) *Match {
	match.AccountID = playerData.AccountID
	match.Avatar = playerData.Avatar
	match.Match = matchType
	match.Nickname = playerData.PlayerName
	match.Region = region
	match.Season = season
	match.Stats.Performance = &Performance{}
	match.Stats.SkillRating = &SkillRating{}
	match.Stats.PerGame = &PerGame{}
	match.Stats.Combat = &Combat{}
	match.Stats.Survival = &Survival{}
	match.Stats.Distance = &Distance{}
	match.Stats.Support = &Support{}

	playerStats := map[string]reflect.Value{
		`Performance`:  reflect.ValueOf(match.Stats.Performance).Elem(),
		`Skill Rating`: reflect.ValueOf(match.Stats.SkillRating).Elem(),
		`Per Game`:     reflect.ValueOf(match.Stats.PerGame).Elem(),
		`Combat`:       reflect.ValueOf(match.Stats.Combat).Elem(),
		`Survival`:     reflect.ValueOf(match.Stats.Survival).Elem(),
		`Distance`:     reflect.ValueOf(match.Stats.Distance).Elem(),
		`Support`:      reflect.ValueOf(match.Stats.Support).Elem(),
	}

	for _, stat := range playerData.Stats {
		if stat.Match == matchType && stat.Region == region && stat.Season == season {
			for _, sonStat := range stat.Stats {
				playerStats[sonStat.Category].FieldByName(sonStat.Field).Set(reflect.ValueOf(sonStat.Value))
			}
		}
	}
	if match.Stats.Performance.KillDeathRatio == "" {
		return nil
	}

	return match
}

// 赛季信息
type RegionInfo struct {
	Season string `json:"season"`
	Data   *struct {
		Duo   []string `json:"duo"`
		Solo  []string `json:"solo"`
		Squad []string `json:"squad"`
	} `json:"data"`
}

// 赛季信息结构
type Regions struct {
	RegionInfo map[string]*RegionInfo `json:"region_info"`
}

// 获取用户赛季信息
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

// 历史比赛信息http返回结构
type History struct {
	Histories []interface{} `json:"histories"`
}

// 获取用户历史比赛信息
func (history *History) GetHistory(match string, playerData *PlayerData) []byte {
	history.Histories = []interface{}{}
	for _, stat := range playerData.MatchHistory {
		if match != "all" {
			if strings.EqualFold(match, stat.MatchDisplay) {
				history.Histories = append(history.Histories, stat)
			}
		} else {
			history.Histories = append(history.Histories, stat)
		}
	}

	js, _ := json.Marshal(history)
	return js
}

// 战绩积分时间kd结构
type Overview struct {
	Score  string
	KD     string
	Played int
}

// 获取用户战绩积分，kd，时间
func (overview *Overview) GetOverview(data *PlayerData, matchType string, region string, season string) {
	overview.Played = data.TimePlayed
	match := &Match{}
	match.GetStats(matchType, region, season, data)
	overview.KD = match.Stats.Performance.KillDeathRatio
	overview.Score = match.Stats.SkillRating.BestRating
}
