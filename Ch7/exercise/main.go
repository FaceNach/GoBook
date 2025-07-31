//1. You have been asked to manage a basketball league and are going to write a
//program to help you. Define two types. The first one, called Team, has a field for
//the name of the team and a field for the player names. The second type is called
//League and has a field called Teams for the teams in the league and a field called
//Wins that maps a team’s name to its number of wins.
//
//2. Add two methods to League. The first method is called MatchResult. It takes
//four parameters: the name of the first team, its score in the game, the name of
//the second team, and its score in the game. This method should update the Wins
//field in League. Add a second method to League called Ranking that returns a
//slice of the team names in order of wins. Build your data structures and call these
//methods from the main function in your program using some sample data.
//
//3. Define an interface called Ranker that has a single method called Ranking that
//returns a slice of strings. Write a function called RankPrinter with two parame‐
//ters, the first of type Ranker and the second of type io.Writer. Use the io.Write
//String function to write the values returned by Ranker to the io.Writer, with a
//newline separating each result. Call this function from main.

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Team struct {
	Name        string
	PlayersName []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1, team2 string, score1, score2 int) {
	if score1 == score2 {
		fmt.Println("You cannot tie in this sport my man")
		return
	}

	if score1 > score2 {
		l.Wins[team1]++
	} else {
		l.Wins[team2]++
	}
}

func (l League) Ranking() []string {

	type Rank struct {
		Key   string
		Value int
	}

	sortedRank := make([]Rank, 0, len(l.Teams))

	for k, v := range l.Wins {
		sortedRank = append(sortedRank, Rank{k, v})
	}

	sort.Slice(sortedRank, func(i, j int) bool {
		return sortedRank[i].Value > sortedRank[j].Value
	})

	fr := make([]string, 0, len(sortedRank))

	for i, w := range sortedRank {
		fr = append(fr, strconv.Itoa(i+1)+") "+w.Key+": "+strconv.Itoa(w.Value))
	}

	return fr
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	or := r.Ranking()
	output := strings.Join(or, "\n")

	fmt.Fprintln(w, output+"\n")

}

func main() {

	union := Team{
		Name:        "Union",
		PlayersName: []string{"Rosales", "Quiroga", "Triverio"},
	}

	colon := Team{
		Name:        "Colon",
		PlayersName: []string{"Fuertes", "Ibarra", "Vignatti"},
	}

	boca := Team{
		Name:        "Boca",
		PlayersName: []string{"Riquelme", "Palermo", "Bianchi"},
	}

	lArgentina := League{
		Teams: []Team{union, colon, boca},
		Wins: map[string]int{
			union.Name: 5,
			colon.Name: 3,
			boca.Name:  1,
		},
	}

	lArgentina.MatchResult(union.Name, colon.Name, 89, 80)

	fmt.Println("La liga Argentina")
	RankPrinter(lArgentina, os.Stdout)

}
