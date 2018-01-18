package pubg

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"record/db"
	"record/log"
	"regexp"
	"strings"
	"time"
)

const (
	UserStatUrl    = "https://pubgtracker.com/profile/pc/{nickname}"
	RobotAuthUrl   = "https://pubgtracker.com/cdn-cgi/l/chk_jschl?"
	UserAgent      = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36"
	FindJschlVc    = "input[name=jschl_vc]"
	FindPass       = "input[name=pass]"
	JschlVcKey     = "jschl_vc"
	PassKey        = "pass"
	AnswerKey      = "jschl_answer"
	PlayerDataRule = `playerData.*?=.*?({.*}).*?;`
)

type Pubg struct {
	client *http.Client
	mongo  *db.Mongo
}

// pubgtracker 网站机器人验证
func (pubg *Pubg) AuthRobot(response *http.Response) {
	responseContent, _ := ioutil.ReadAll(response.Body)
	query, _ := goquery.NewDocumentFromReader(strings.NewReader(string(responseContent)))
	request, _ := http.NewRequest("GET", RobotAuthUrl, nil)
	request.Header.Add("user-agent", UserAgent)
	q := request.URL.Query()
	jschlVc, _ := query.Find(FindJschlVc).Attr("value")
	pass, _ := query.Find(FindPass).Attr("value")
	q.Add(JschlVcKey, jschlVc)
	q.Add(PassKey, pass)
	q.Add(AnswerKey, getJsAnswer(string(responseContent)))
	request.URL.RawQuery = q.Encode()
	time.Sleep(4 * time.Second)
	pubg.client.Do(request)
}

// 获取数据
func (pubg *Pubg) Get(nickname string) *PlayerData {
	request, _ := http.NewRequest("GET", strings.Replace(UserStatUrl, "{nickname}", nickname, -1), nil)
	request.Header.Add("user-agent", UserAgent)
	response, err := pubg.client.Do(request)
	if err != nil {
		log.Info("连接请求错误" + err.Error())
	} else {
		if response.StatusCode == 503 {
			pubg.AuthRobot(response)
			return pubg.Get(nickname)
		} else {
			body, _ := ioutil.ReadAll(response.Body)
			playerData := &PlayerData{}
			json.Unmarshal([]byte(pubg.PlayerData(body)), playerData)
			if playerData.AccountID != "" {
				pubg.Save(playerData)
			}
			return playerData
		}
	}
	return nil
}

func (pubg *Pubg) FindBySteamId(steamId int64) *PlayerData {
	query := struct {
		SteamID int64 `bson:"steam_id"`
	}{}
	query.SteamID = steamId
	find := pubg.mongo.Select(query)
	count, _ := find.Count()
	if count == 0 {
		return nil
	} else {
		playerData := &PlayerData{}
		err := find.One(playerData)
		if err != nil {
			log.Info("mongo查询错误" + err.Error())
		}
		TrackerGo.New().Add(playerData.PlayerName)
		return playerData
	}
}

func (pubg *Pubg) FindByAccountId(accountId string) *PlayerData {
	query := struct {
		AccountID string `bson:"_id"`
	}{}
	query.AccountID = accountId
	find := pubg.mongo.Select(query)
	count, _ := find.Count()
	if count == 0 {
		return nil
	} else {
		playerData := &PlayerData{}
		err := find.One(playerData)
		if err != nil {
			log.Info("mongo查询错误" + err.Error())
		}
		return playerData
	}
}

func (pubg *Pubg) Find(nickname string) *PlayerData {
	query := struct {
		PlayerName string `bson:"player_name"`
	}{}
	query.PlayerName = nickname
	find := pubg.mongo.Select(query)
	count, _ := find.Count()
	if count == 0 {
		return pubg.Get(nickname)
	} else {
		playerData := &PlayerData{}
		err := find.One(playerData)
		if err != nil {
			log.Info("mongo查询错误:" + err.Error())
		}
		TrackerGo.New().Add(nickname)
		return playerData
	}
}

func (pubg *Pubg) Save(playerData *PlayerData) {
	query := struct {
		AccountId string `bson:"_id"`
	}{}
	query.AccountId = playerData.AccountID
	data := pubg.mongo.Select(query)
	count, err := data.Count()
	if err != nil {
		log.Info("mongo查询错误:" + err.Error())
	} else {
		if count > 0 {
			recordInfo := &PlayerData{}
			err := data.One(recordInfo)
			if err != nil {
				log.Info(err)
			} else {
				if playerData.SteamID == 0 {
					playerData.SteamID = recordInfo.SteamID
				}
				playerData.CreatedAt = recordInfo.CreatedAt
			}

			playerData.UpdatedAt = time.Now().Unix()
			pubg.mongo.Update(query, playerData)
		} else {
			playerData.CreatedAt = time.Now().Unix()
			playerData.UpdatedAt = time.Now().Unix()
			pubg.mongo.Insert(playerData)
		}
	}
}

// 获取用户战绩信息
func (pubg *Pubg) PlayerData(find []byte) string {
	reg := regexp.MustCompile(PlayerDataRule)
	playerData := string(reg.Find(find))
	playerData = strings.Replace(playerData, "playerData", "", -1)
	playerData = strings.Replace(playerData, "=", "", -1)
	playerData = strings.Replace(playerData, ";", "", -1)
	return playerData
}

func New(dialInfo *mgo.DialInfo) *Pubg {
	pubg := &Pubg{client: &http.Client{}, mongo: db.New(dialInfo)}
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Info("cookieJar实例化错误:" + err.Error())
	} else {
		pubg.client.Jar = jar
	}
	return pubg
}

var TrackerGo *Tracker

type Tracker struct {
	nicknames chan string
}

func (tracker *Tracker) New() *Tracker {
	if TrackerGo == nil {
		TrackerGo = &Tracker{
			nicknames: make(chan string, 1000),
		}
	}
	return TrackerGo
}

func (tracker *Tracker) Add(nickname string) {
	if len(tracker.nicknames) < 1000 {
		tracker.nicknames <- nickname
	}
}

func (tracker *Tracker) Do(pubgTracker *Pubg) {
	go func(nicknames chan string) {
		for {
			select {
			case nickname := <-nicknames:
				log.Info("用户查询更新:" + nickname)
				pubgTracker.Get(nickname)
				//log.Info(playerData)
			}
		}
	}(tracker.nicknames)
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
