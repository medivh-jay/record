package main

import (
	"encoding/json"
	"flag"
	"gopkg.in/mgo.v2"
	"net/http"
	"net/url"
	"record/log"
	"record/pubg"
	"strconv"
	"time"
)

func main() {
	port := flag.String("port", "9999", "start port")
	mgoIp := flag.String("host", "127.0.0.1:27017", "mongo ip")
	username := flag.String("username", "tracker", "mongo username")
	password := flag.String("password", "e5d$e(Gs%epN3nDb", "mongo password")
	flag.Parse()
	log.Info("服务启动!时间:" + time.Now().String() + ",mongoDB地址:" + *mgoIp + ",mongo验证（用户名:" + *username + ", 密码:" + *password + "）")
	mgoDialInfo := &mgo.DialInfo{Username: *username, Password: *password, Addrs: []string{*mgoIp}}

	pubgTracker := pubg.New(mgoDialInfo)
	pubg.TrackerGo.New().Do(pubgTracker)

	nilStruct := struct {
		ErrorInfo string `json:"error_info"`
	}{}

	successStruct := struct {
		Result bool `json:"result"`
	}{}

	nilStruct.ErrorInfo = "false"

	js, _ := json.Marshal(nilStruct)

	http.HandleFunc("/pubg/season.json", func(writer http.ResponseWriter, request *http.Request) {
		season := &pubg.Season{}
		season.Load()
		writer.Write(season.GetSeason())
	})

	http.HandleFunc("/pubg/userinfo.json", func(writer http.ResponseWriter, request *http.Request) {

		params, err := url.ParseQuery(request.URL.RawQuery)
		if err != nil {
			log.Info(err.Error())
		} else {
			if _, ok := params["nickname"]; ok {
				log.Info("查询用户nickname：" + params["nickname"][0])
				data := pubgTracker.Find(params["nickname"][0])
				if data.AccountID == "" {
					writer.Write(js)
				} else {
					userInfo := &pubg.UserInfo{}
					writer.Write(userInfo.ToJSON(data))
				}
			} else if _, ok := params["steamid"]; ok {
				log.Info("查询用户steamId：" + params["steamid"][0])
				steamID, err := strconv.ParseInt(params["steamid"][0], 10, 64)
				if err != nil {
					writer.Write(js)
				} else {
					data := pubgTracker.FindBySteamId(steamID)
					if data.AccountID == "" {
						writer.Write(js)
					} else {
						userInfo := &pubg.UserInfo{}
						writer.Write(userInfo.ToJSON(data))
					}
				}
			} else {
				writer.Write(js)
			}
		}
	})

	http.HandleFunc("/pubg/match.json", func(writer http.ResponseWriter, request *http.Request) {
		params, _ := url.ParseQuery(request.URL.RawQuery)

		match := &pubg.Match{}
		accountId := params["account_id"][0]
		data := pubgTracker.FindByAccountId(accountId)
		if data != nil {

			if data.AccountID == "" {
				writer.Write(js)
			}

			season := data.DefaultSeason

			if _, ok := params["season"]; ok {
				season = params["season"][0]
			}

			regionInfo := &pubg.Regions{}
			regionInfo.GetUserRegion(data)

			soloLen := len(regionInfo.RegionInfo[season].Data.Solo)
			duoLen := len(regionInfo.RegionInfo[season].Data.Duo)
			squadLen := len(regionInfo.RegionInfo[season].Data.Squad)
			matchType := "solo"

			region := "agg"
			if soloLen > 0 {
				matchType = "solo"
				region = regionInfo.RegionInfo[season].Data.Solo[0]
			} else if duoLen > 0 {
				matchType = "duo"
				region = regionInfo.RegionInfo[season].Data.Duo[0]
			} else if squadLen > 0 {
				matchType = "squad"
				region = regionInfo.RegionInfo[season].Data.Squad[0]
			}

			if _, ok := params["match"]; ok {
				matchType = params["match"][0]
			}

			if _, ok := params["region"]; ok {
				region = params["region"][0]
			}
			result := match.GetStats(matchType, region, season, data)
			if result != nil {
				writer.Write(result.ToJSON())
			} else {
				writer.Write(js)
			}
		} else {
			writer.Write(js)
		}
	})

	http.HandleFunc("/pubg/regioninfo.json", func(writer http.ResponseWriter, request *http.Request) {
		params, _ := url.ParseQuery(request.URL.RawQuery)
		accountId := params["account_id"][0]
		data := pubgTracker.FindByAccountId(accountId)
		if data == nil || data.AccountID == "" {
			writer.Write(js)
		} else {
			regionInfo := pubg.Regions{}
			writer.Write(regionInfo.GetUserRegion(data))
		}
	})

	http.HandleFunc("/pubg/history.json", func(writer http.ResponseWriter, request *http.Request) {
		params, _ := url.ParseQuery(request.URL.RawQuery)
		accountId := params["account_id"][0]

		match := "all"
		if _, ok := params["match"]; ok {
			match = params["match"][0]
		}

		data := pubgTracker.FindByAccountId(accountId)
		if data == nil || data.AccountID == "" {
			writer.Write(js)
		} else {
			history := pubg.History{}
			writer.Write(history.GetHistory(match, data))
		}
	})

	http.HandleFunc("/pubg/overview", func(writer http.ResponseWriter, request *http.Request) {
		params, _ := url.ParseQuery(request.URL.RawQuery)
		if _, ok := params["nickname"]; ok {
			nickname := params["nickname"][0]
			overview := &pubg.Overview{}
			playerData := pubgTracker.Find(nickname)
			season := &pubg.Season{}
			season.Load()
			overview.GetOverview(playerData, "solo", "agg", season.Season[0].Season)
			js, _ := json.Marshal(overview)
			writer.Write(js)
		} else if _, ok := params["steamid"]; ok {
			steamId := params["steamid"][0]
			overview := &pubg.Overview{}
			steamID, _ := strconv.ParseInt(steamId, 10, 64)
			playerData := pubgTracker.FindBySteamId(steamID)
			season := &pubg.Season{}
			season.Load()
			overview.GetOverview(playerData, "solo", "agg", season.Season[0].Season)
			js, _ := json.Marshal(overview)
			writer.Write(js)
		}
	})

	http.HandleFunc("/pubg/upload.json", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		if request.Method != "POST" {
			nilStruct.ErrorInfo = "请求方式错误，当前请求方式:" + request.Method + ",期待的请求方式: POST"
			result, _ := json.Marshal(nilStruct)
			writer.Write(result)
		} else {
			if uploadString, ok := request.Form["pubg"]; ok {
				loadString := uploadString[0]
				log.Info(loadString)
				info := pubgTracker.PlayerData([]byte("var playerData = " + loadString))
				playerData := &pubg.PlayerData{}
				json.Unmarshal([]byte(info), playerData)
				if playerData.AccountID != "" {
					if steamIDs, ok := request.Form["steamId"]; ok {
						steamID := steamIDs[0]
						playerData.SteamID, _ = strconv.Atoi(steamID)
					}
					pubgTracker.Save(playerData)
					successStruct.Result = true
					result, _ := json.Marshal(successStruct)
					writer.Write(result)
				} else {
					nilStruct.ErrorInfo = "数据解析失败，无法存入,接收的数据为： " + request.Form["pubg"][0]
					result, _ := json.Marshal(nilStruct)
					writer.Write(result)
				}
			} else {
				nilStruct.ErrorInfo = "接收到的数据没有找到key : pubg"
				result, _ := json.Marshal(nilStruct)
				writer.Write(result)
			}
		}

	})
	http.ListenAndServe(":"+*port, nil)
}
