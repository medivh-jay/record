package kotk

import (
	"net/http"
	"record/log"
)

// 区服
var Region = []string{
	1: "北美",
	2: "欧洲",
	3: "南美",
	4: "亚太",
	5: "澳大利亚",
}

// 游戏模式
var TeamModel = []string{
	1: "Solos",
	2: "Duos",
	3: "Fives",
}

// 赛季
var SeasonInfo = []string{
	1: "Pre-Season 1",
	2: "Pre-Season 2",
	3: "Pre-Season 3",
	4: "Pre-Season 4",
	5: "Pre-Season 5",
	6: "Pre-Season 6",
}

// 段位
var Tier = []string{
	1: "青铜",
	2: "白银",
	3: "黄金",
	4: "白金",
	5: "钻石",
	6: "Master",
	7: "皇家",
}

// 用名字搜索id的地址
const NameSearChUrl = "https://census.daybreakgames.com/rest/leaderboard/kotk/name-search"

// 根据用户id搜索用户数据
const UserIdGetPage = "https://census.daybreakgames.com/rest/leaderboard/kotk/game-user-id-get-page"

type Request struct {
	Search Search
	client http.Client
}

func (request *Request) NameSearch(userName string) {
	request.Search.UserName = userName

}

func (request *Request) Request(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Info(err.Error())
	} else {
		req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.78 Safari/537.36")
		req.Header.Add("Referer", "https://www.h1z1.com/king-of-the-kill/leaderboards")
		//req.Form.Add("")
	}
}
