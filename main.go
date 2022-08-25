package main

import (
	"fmt"
	"generics_example/collections"
)

type ProgrammingLanguage struct {
	Id           int
	Name         string
	ReleasedYear int
	GithutRank   int
	PyplRank     int
}

var programmingLanguages = []ProgrammingLanguage{
	{Id: 1, Name: "JavaScript", ReleasedYear: 1995, GithutRank: 1, PyplRank: 3},
	{Id: 3, Name: "Java", ReleasedYear: 1995, GithutRank: 3, PyplRank: 2},
	{Id: 2, Name: "Python", ReleasedYear: 1991, GithutRank: 2, PyplRank: 1},
	{Id: 4, Name: "TypeScript", ReleasedYear: 2012, GithutRank: 7, PyplRank: 10},
	{Id: 5, Name: "C#", ReleasedYear: 2000, GithutRank: 9, PyplRank: 4},
}

func main() {

	fmt.Println("01 - Find")
	p := collections.Find(programmingLanguages, func(p ProgrammingLanguage) bool {
		return p.Id == 4
	})
	fmt.Println(*p)

	fmt.Println()

	fmt.Println("02 - Filter")
	languages := collections.FilterArray(programmingLanguages, func(p ProgrammingLanguage) bool {
		return p.GithutRank > 6
	})
	fmt.Println(languages)

	fmt.Println()

	fmt.Println("03 - Map")
	langNames := collections.MapArray(programmingLanguages, func(p ProgrammingLanguage) string {
		return p.Name
	})
	fmt.Println(langNames)

	fmt.Println()

	fmt.Println("04 - Some")
	shoudBeTrue := collections.Some(programmingLanguages, func(p ProgrammingLanguage) bool {
		return p.GithutRank == 7
	})

	shoudBeFalse := collections.Some(programmingLanguages, func(p ProgrammingLanguage) bool {
		return p.GithutRank == 4
	})

	fmt.Println("shoudBeTrue: ", shoudBeTrue, "shoudBeFalse: ", shoudBeFalse)

	fmt.Println()

	fmt.Println("04 - Every")
	shoudBeTrue = collections.Every(programmingLanguages, func(p ProgrammingLanguage) bool {
		return p.GithutRank < 10
	})

	shoudBeFalse = collections.Every(programmingLanguages, func(p ProgrammingLanguage) bool {
		return p.GithutRank > 4
	})
	fmt.Println("shoudBeTrue: ", shoudBeTrue, "shoudBeFalse: ", shoudBeFalse)

	fmt.Println()

	fmt.Println("05 - Group")
	sumOfRankings := collections.Group(programmingLanguages, func(p ProgrammingLanguage, amount int) int {
		return p.GithutRank + amount
	})
	fmt.Println("Some of Rankings: ", sumOfRankings)
}
