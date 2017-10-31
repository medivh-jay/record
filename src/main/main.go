package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"record/pubg"
)

func main() {
	pubgTracker := pubg.New()

	nilStruct := struct {
		ErrorInfo string `json:"error_info"`
	}{}

	nilStruct.ErrorInfo = "false"

	js, _ := json.Marshal(nilStruct)

	http.HandleFunc("/pubg/season.json", func(writer http.ResponseWriter, request *http.Request) {
		season := &pubg.Season{}
		writer.Write(season.GetSeason())
	})

	http.HandleFunc("/pubg/userinfo.json", func(writer http.ResponseWriter, request *http.Request) {

		params, err := url.ParseQuery(request.URL.RawQuery)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(params["nickname"][0])
			data := pubgTracker.Find(params["nickname"][0])
			if data.AccountID == "" {
				writer.Write(js)
			} else {
				userInfo := &pubg.UserInfo{}
				writer.Write(userInfo.ToJSON(data))
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

			if soloLen > 0 {
				matchType = "solo"
			} else if duoLen > 0 {
				matchType = "duo"
			} else if squadLen > 0 {
				matchType = "squad"
			}

			if _, ok := params["match"]; ok {
				matchType = params["match"][0]
			}
			region := "agg"
			if _, ok := params["region"]; ok {
				region = params["region"][0]
			}
			result := match.GetStats(matchType, region, season, data)
			if result != nil {
				writer.Write(result)
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

	http.ListenAndServe(":9999", nil)
}
