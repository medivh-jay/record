package main

import (
	"fmt"
	"net/http"
	"net/url"
	"record/pubg"
)

func main() {
	pubgTracker := pubg.New()

	http.HandleFunc("/season.json", func(writer http.ResponseWriter, request *http.Request) {
		season := &pubg.Season{}
		writer.Write(season.GetSeason())
	})

	http.HandleFunc("/userinfo.json", func(writer http.ResponseWriter, request *http.Request) {

		params, err := url.ParseQuery(request.URL.RawQuery)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(params["nickname"][0])
			data := pubgTracker.Find(params["nickname"][0])
			userInfo := &pubg.UserInfo{}
			writer.Write(userInfo.ToJSON(data))
		}
	})

	http.ListenAndServe(":9999", nil)
}
