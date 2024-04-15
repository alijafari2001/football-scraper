package main

import (
	"flag"
	"fmt"
	"football-scraper/pkg"
	"log"
)

func handle(config map[string]string, c pkg.Command) error {
	if c.StandingTable {
		url := config[c.LeagueName+"-"+c.Season]
		doc, err := pkg.Scrape(url)
		if err != nil {
			return err
		}
		parser := pkg.NewStandingTableParser()
		parser.Parse(doc)
		pkg.NewStandingTableDrawer().Draw(parser.GetData())
	} else if c.Fixures != 0 {
		url := config["fixtures-"+c.LeagueName+"-"+c.Season]
		doc, err := pkg.Scrape(url)
		if err != nil {
			return err
		}
		parser := pkg.NewFixturesTableParser(c.Fixures)
		parser.Parse(doc)
		pkg.NewFixturesTableDrawer().Draw(parser.GetData())
	} else if c.GoalTable {
		url := config[c.LeagueName+"-"+c.Season]
		doc, err := pkg.Scrape(url)
		if err != nil {
			return err
		}
		parser := pkg.NewGoalTableParser()
		parser.Parse(doc)
		pkg.NewGoalTableDrawer().Draw(parser.GetData())
	}
	return nil
}

func main() {
	config, err := pkg.LoadConfigs()
	if err != nil {
		fmt.Println("hey")
		log.Fatal(err)
	}
	command := parseCommand()
	command.SetDefaultFlags()
	err = command.CheckFlags()
	if err != nil {
		log.Fatal(err)
	}
	err = handle(config, *command)
	if err != nil {
		log.Fatal(err)
	}
}

func parseCommand() *pkg.Command {
	var leagueName, season string
	var standingTable, goalTable, assistTable bool
	var fixtures int

	flag.StringVar(&leagueName, "l", "pl", `league names are:
	'pl': premier league
	'la': la liga
	'se': serie A
	'bu': bundes liga
	'l1': ligue 1`)
	flag.StringVar(&season, "s", "23/24", `specify your season [you can pick a season from 09/10 to 23/24]`)
	flag.BoolVar(&standingTable, "t", false, "display league standing table")
	flag.BoolVar(&goalTable, "gt", false, "display league top scorers")
	flag.BoolVar(&assistTable, "at", false, "display league top assist providers table")
	flag.IntVar(&fixtures, "f", 0, "display fixures of the specified game week")
	flag.Parse()
	return &pkg.Command{
		LeagueName:    leagueName,
		Season:        season,
		StandingTable: standingTable,
		GoalTable:     goalTable,
		AssistTable:   assistTable,
		Fixures:       fixtures,
	}
}
