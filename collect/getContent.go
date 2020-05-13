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
	Rank     string `json:"rank"`
	Team     string `json:"team"`
	Player   string `json:"player"`
	Value    string `json:"value"`
}

// Start parsing
func Start() {
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
