package pkg

import (
	"fmt"
)

type Command struct {
	LeagueName    string
	Season        string
	StandingTable bool
	GoalTable     bool
	AssistTable   bool
	Fixures       int
}

func (c *Command) SetDefaultFlags() {
	if !c.StandingTable && !c.GoalTable && !c.AssistTable && c.Fixures == 0 {
		c.StandingTable = true
	}
}

func (c Command) CheckFlags() error {
	switch c.LeagueName {
	case "pl", "la", "se", "l1", "bu":
		break
	default:
		return fmt.Errorf("invalid league name!")
	}
	switch c.Season {
	case "09/10", "10/11", "11/12", "12/13", "13/14", "14/15", "15/16", "16/17", "17/18", "18/19", "19/20", "20/21", "21/22", "22/23", "23/24":
		break
	default:
		return fmt.Errorf("invalid season!")
	}
	switch {
	case c.Fixures < 0 || c.Fixures > 38:
		return fmt.Errorf("invalid fixures flag")
	case (c.LeagueName == "bu" || c.LeagueName == "l1") && c.Fixures > 34:
		return fmt.Errorf("invalid fixtures flag")
	}
	return nil
}
