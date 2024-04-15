package pkg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Parser interface {
	Parse(*goquery.Document)
	GetData() [][]string
}

type StandingTableParser struct {
	entries []standingTableEntry
}

func NewStandingTableParser() *StandingTableParser {
	return &StandingTableParser{entries: make([]standingTableEntry, 0)}
}

func (p *StandingTableParser) Parse(doc *goquery.Document) {
	rank := 1
	doc.Find("table.stats_table").First().Find("tbody tr").Each(func(row int, tr *goquery.Selection) {
		entry := standingTableEntry{
			rank: strconv.Itoa(rank),
		}
		rank++
		tr.Find("td").Each(func(col int, td *goquery.Selection) {
			dataStat, _ := td.Attr("data-stat")
			switch val := strings.TrimSpace(td.Text()); dataStat {
			case "team":
				entry.squad = val
			case "games":
				entry.mp = val
			case "wins":
				entry.w = val
			case "ties":
				entry.d = val
			case "losses":
				entry.l = val
			case "goals_for":
				entry.gf = val
			case "goals_against":
				entry.ga = val
			case "goal_diff":
				entry.gd = val
			case "points":
				entry.pts = val
			case "attendance_per_g":
				entry.attendance = val
			case "top_team_scorers":
				entry.topScorer = val
			case "top_keeper":
				entry.goalKeeper = val
			}
		})
		p.entries = append(p.entries, entry)
	})
}

func (p StandingTableParser) GetData() [][]string {
	data := make([][]string, 0)
	for _, entry := range p.entries {
		row := make([]string, 0)
		row = append(row, entry.rank, entry.squad, entry.mp, entry.w, entry.d, entry.l, entry.gf, entry.ga, entry.gd, entry.pts, entry.attendance, entry.topScorer, entry.goalKeeper)
		data = append(data, row)
	}
	return data
}

type fixturesTableParser struct {
	fixtureNumber int
	entries       []fixturesTableEntry
}

func NewFixturesTableParser(num int) *fixturesTableParser {
	return &fixturesTableParser{
		fixtureNumber: num,
		entries:       make([]fixturesTableEntry, 0)}
}

func (p *fixturesTableParser) Parse(doc *goquery.Document) {
	wk := ""
	day := ""
	date := ""
	doc.Find("table.stats_table").First().Find("tbody tr").Each(func(row int, tr *goquery.Selection) {
		entry := fixturesTableEntry{}
		tr.Find("th").Each(func(i int, th *goquery.Selection) {
			dataStat, exist := th.Attr("data-stat")
			if exist {
				if dataStat == "gameweek" {
					if text := strings.TrimSpace(th.Text()); text != "" {
						entry.wk = text
						wk = text
					} else {
						entry.wk = wk
					}
				}
			}
		})
		tr.Find("td").Each(func(col int, td *goquery.Selection) {
			dataStat, _ := td.Attr("data-stat")
			switch val := strings.TrimSpace(td.Text()); dataStat {
			case "gameweek":
				if val != "" {
					entry.wk = val
					wk = val
				} else {
					entry.wk = wk
				}
			case "dayofweek":
				if val != "" {
					entry.day = val
					day = val
				} else {
					entry.day = day
				}
			case "date":
				if val != "" {
					entry.date = val
					date = val
				} else {
					entry.date = date
				}
			case "start_time":
				if val != "" {
					entry.time = val
				} else {
					entry.time = "null"
				}
			case "home_team":
				entry.home = val
			case "score":
				if val != "" {
					entry.score = val
				} else {
					entry.score = "TBD"
				}
			case "away_team":
				entry.away = val
			case "venue":
				if val != "" {
					entry.venue = val
				} else {
					entry.venue = "null"
				}
			case "referee":
				if val != "" {
					entry.refree = val
				} else {
					entry.refree = "null"
				}
			}
		})
		if entry.home != "" {
			p.entries = append(p.entries, entry)

		}
	})
}

func (p fixturesTableParser) GetData() [][]string {
	data := make([][]string, 0)
	for _, entry := range p.entries {
		if wk, _ := strconv.Atoi(entry.wk); wk == p.fixtureNumber {
			row := make([]string, 0)
			row = append(row, entry.wk, entry.day, entry.date, entry.time, entry.home, entry.score, entry.away, entry.venue, entry.refree)
			data = append(data, row)
		}
	}
	return data
}

type goalTableParser struct {
	entries []goalTableEntry
}

func NewGoalTableParser() *goalTableParser {
	return &goalTableParser{entries: make([]goalTableEntry, 0)}
}

func (p *goalTableParser) Parse(doc *goquery.Document) {
	fmt.Println(doc.Text())
	doc.Find("div.data_grid").Find("div#leaders_goals").Find("table.columns").Find("tbody tr").Each(func(row int, tr *goquery.Selection) {
		entry := goalTableEntry{}
		tr.Find("td").Each(func(col int, td *goquery.Selection) {
			class, _ := td.Attr("class")
			switch text := strings.TrimSpace(td.Text()); class {
			case "rank":
				entry.rank = text
			case "who":
				entry.player = text
			case "value":
				entry.goal = text
			}
		})
		p.entries = append(p.entries, entry)
	})
}

func (p goalTableParser) GetData() [][]string {
	data := make([][]string, 0)
	for _, entry := range p.entries {
		row := make([]string, 0)
		row = append(row, entry.rank, entry.player, entry.goal)
		data = append(data, row)
	}
	return data
}
