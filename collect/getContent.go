package collect

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

const (
	cpblTopListURL = "http://www.cpbl.com.tw/stats/toplist.html"
)

// SingleStats :
type SingleStats struct {
	ItemName string `json:"itemName"`
	ItemType string `json:"itmeType"` // batter, pitcher
	Rank     string `json:"rank"`
	Team     string `json:"team"` // m, l, g, b
	Player   string `json:"player"`
	Value    string `json:"value"`
}

var (
	batterTypeInTopList = []string{
		"AVG",
		"H",
		"HR",
		"RBI",
		"SB",
		"TB",
	}
	pitcherTypeInTopList = []string{
		"ERA",
		"W",
		"SV",
		"WHIP",
		"HLD",
		"SO",
	}

	teamTranslator = map[string]string{
		"樂天":   "m",
		"統一獅":  "l",
		"中信兄弟": "b",
		"富邦":   "g",
	}
)

const (

	// batter
	getBBURL       = "http://www.cpbl.com.tw/stats/all.html?&game_type=01&online=0&year=2020&stat=pbat&sort=BB&order=desc&over=1"
	getOBPURL      = "http://www.cpbl.com.tw/stats/all.html?&game_type=01&online=0&year=2020&stat=pbat&sort=OBP&order=desc"
	getSLGURL      = "http://www.cpbl.com.tw/stats/all.html?&game_type=01&online=0&year=2020&stat=pbat&sort=SLG&order=desc"
	getBatterSOURL = "http://www.cpbl.com.tw/stats/all.html?&game_type=01&online=0&year=2020&stat=pbat&sort=SO&order=desc"
	// 雙殺打
	getGIDPURL = "http://www.cpbl.com.tw/stats/all.html?&game_type=01&online=0&year=2020&stat=pbat&sort=GIDP&order=desc&over=1"

	// pitcher
	getPitchGameURL = "http://www.cpbl.com.tw/stats/all.html?&game_type=01&online=0&year=2020&stat=ppit"
	getIPURL        = "http://www.cpbl.com.tw/stats/all.html?&game_type=01&online=0&year=2020&stat=ppit&sort=IP&order=desc"
	getNPURL        = "http://www.cpbl.com.tw/stats/all.html?&game_type=01&online=0&year=2020&stat=ppit&sort=NP&order=desc&over=1"
	getLoseURL      = "http://www.cpbl.com.tw/stats/all.html?&game_type=01&online=0&year=2020&stat=ppit&sort=LOSE&order=desc"
	getHRURL        = "http://www.cpbl.com.tw/stats/all.html?&game_type=01&online=0&year=2020&stat=ppit&sort=HR&order=desc&over=1"
	getBSURL        = "http://www.cpbl.com.tw/stats/all.html?&game_type=01&online=0&year=2020&stat=ppit&sort=BS&order=desc"

	// @@todo
	// 1 find error
	// 2 filter the SLG, with G or PA
)

// StartGetTop5ByURL :
func StartGetTop5ByURL(url string) {
	c := colly.NewCollector()
	stats := []SingleStats{}

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visting", string(r.Body))
	})

	c.OnHTML("table.std_tb", func(e *colly.HTMLElement) {

		for i := 1; i <= 5; i++ {
			rank := e.ChildText("tr:nth-of-type(" + strconv.Itoa(i+1) + ") td:nth-of-type(1)")
			name := e.ChildText("tr:nth-of-type(" + strconv.Itoa(i+1) + ") td:nth-of-type(2) a")
			team := e.ChildAttr("tr:nth-of-type("+strconv.Itoa(i+1)+") td:nth-of-type(2) img", "src")
			// 		value := e.ChildText("tr:nth-of-type(" + strconv.Itoa(i+1) + ") td:nth-of-type(4)")

			fmt.Printf("Rank%s: %s %s\n", rank, name, team)
			// 		neStat := SingleStats{
			// 			Rank:     rank,
			// 			Team:     team,
			// 			Player:   name,
			// 			Value:    value,
			// 			ItemName: item,
			// 		}
			// 		stats = append(stats, neStat)
		}

	})

	c.Visit(url)

	fmt.Println("final result", stats)
}

// StartGetTopList parsing
func StartGetTopList() {
	c := colly.NewCollector()
	stats := []SingleStats{}

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visting", string(r.Body))
	})

	c.OnHTML("div.statstoplist_box", func(e *colly.HTMLElement) {
		item := e.ChildText("th")

		for i := 1; i <= 5; i++ {
			rank := e.ChildText("tr:nth-of-type(" + strconv.Itoa(i+1) + ") td:nth-of-type(1)")
			team := e.ChildText("tr:nth-of-type(" + strconv.Itoa(i+1) + ") td:nth-of-type(2)")
			name := e.ChildText("tr:nth-of-type(" + strconv.Itoa(i+1) + ") td:nth-of-type(3)")
			value := e.ChildText("tr:nth-of-type(" + strconv.Itoa(i+1) + ") td:nth-of-type(4)")

			fmt.Printf("%s, Rank%d: %s %s %s %s \n", item, i, rank, team, name, value)
			neStat := SingleStats{
				Rank:     rank,
				Team:     team,
				Player:   name,
				Value:    value,
				ItemName: item,
			}
			stats = append(stats, neStat)
		}

	})

	c.Visit(cpblTopListURL)

	fmt.Println("final result", stats)
}

// Start :
func Start() {
	StartGetTop5ByURL(getBatterSOURL)
}
