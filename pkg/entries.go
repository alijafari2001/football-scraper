package pkg

type standingTableEntry struct {
	rank, squad, mp, w, l, d, gf, ga, gd, pts,
	attendance, topScorer, goalKeeper string
}

type fixturesTableEntry struct {
	wk, day, date, time, home, score, away, venue, refree string
}

type goalTableEntry struct {
	rank, player, goal string
}
