package collect

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

type Stats struct {
	ItemName string
	Rank     string
	Team     string
	Player   string
	Value    string
}

func Start() {
	c := colly.NewCollector()
	stats := []Stats{}
	// // Find and visit all links
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })

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
			neStat := Stats{
				Rank:     rank,
				Team:     team,
				Player:   name,
				Value:    value,
				ItemName: item,
			}
			stats = append(stats, neStat)
		}

	})

	c.Visit("http://www.cpbl.com.tw/stats/toplist.html")

	fmt.Println("final result", stats)
}
