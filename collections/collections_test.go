package collections_test

import (
	"fmt"
	"generics_example/collections"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestFind(t *testing.T) {

	for _, testCase := range []struct {
		Name   string
		Test   func(p ProgrammingLanguage) bool
		Assert func(t *testing.T, p *ProgrammingLanguage)
	}{
		{
			Name: "FindSuccessful",
			Test: func(p ProgrammingLanguage) bool { return p.Id == 1 },
			Assert: func(t *testing.T, p *ProgrammingLanguage) {
				require.NotNil(t, p)
				assert.Equal(t, "JavaScript", p.Name)
			},
		},

		{
			Name: "FindNotFound",
			Test: func(p ProgrammingLanguage) bool { return p.Id == 9 },
			Assert: func(t *testing.T, p *ProgrammingLanguage) {
				require.Nil(t, p)
			},
		},

		{
			Name: "FindFirstItem",
			Test: func(p ProgrammingLanguage) bool { return p.GithutRank > 2 },
			Assert: func(t *testing.T, p *ProgrammingLanguage) {
				require.NotNil(t, p)
				assert.Equal(t, "Java", p.Name)
			},
		},
	} {

		t.Run(testCase.Name, func(t *testing.T) {
			p := collections.Find(programmingLanguages, testCase.Test)
			testCase.Assert(t, p)
		})

	}
}

func TestFilter(t *testing.T) {

	for _, testCase := range []struct {
		Name   string
		Test   func(p ProgrammingLanguage) bool
		Assert func(t *testing.T, p []ProgrammingLanguage)
	}{
		{
			Name: "FilterFindTwoItems",
			Test: func(p ProgrammingLanguage) bool { return p.GithutRank > 6 },
			Assert: func(t *testing.T, p []ProgrammingLanguage) {
				assert.NotEmpty(t, p)
				assert.Equal(t, 2, len(p))
				assert.Equal(t, "TypeScript", p[0].Name)
				assert.Equal(t, "C#", p[1].Name)
			},
		},

		{
			Name: "FilterFindAnyItem",
			Test: func(p ProgrammingLanguage) bool { return p.GithutRank == 0 },
			Assert: func(t *testing.T, p []ProgrammingLanguage) {
				assert.Empty(t, p)
			},
		},

		{
			Name: "FilterFindAllItems",
			Test: func(p ProgrammingLanguage) bool { return p.GithutRank < 10 },
			Assert: func(t *testing.T, p []ProgrammingLanguage) {
				assert.NotEmpty(t, p)
				assert.Equal(t, programmingLanguages, p)
			},
		},
	} {

		t.Run(testCase.Name, func(t *testing.T) {
			languages := collections.FilterArray(programmingLanguages, testCase.Test)
			testCase.Assert(t, languages)
		})
	}
}

func TestMap(t *testing.T) {

	for _, testCase := range []struct {
		Name   string
		Test   func(p ProgrammingLanguage) interface{}
		Assert func(t *testing.T, out []interface{})
	}{
		{
			Name: "MapLanguagesName",
			Test: func(p ProgrammingLanguage) interface{} { return p.Name },
			Assert: func(t *testing.T, out []interface{}) {
				assert.NotEmpty(t, out)
				assert.Equal(t, []interface{}{"JavaScript", "Java", "Python", "TypeScript", "C#"}, out)
			},
		},

		{
			Name: "MapLanguageNamesAndRanking",
			Test: func(p ProgrammingLanguage) interface{} {
				return fmt.Sprintf("Name : %s, Git Rank: %d, Pypl Rank: %d", p.Name, p.GithutRank, p.PyplRank)
			},
			Assert: func(t *testing.T, out []interface{}) {
				assert.NotEmpty(t, out)
				assert.Equal(t, "Name : JavaScript, Git Rank: 1, Pypl Rank: 3", out[0])
			},
		},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			out := collections.MapArray(programmingLanguages, testCase.Test)
			testCase.Assert(t, out)
		})
	}
}

func TestSome(t *testing.T) {

	for _, testCase := range []struct {
		Name   string
		Test   func(p ProgrammingLanguage) bool
		Assert bool
	}{
		{
			Name:   "SomeSuccessful",
			Test:   func(p ProgrammingLanguage) bool { return p.GithutRank == 9 },
			Assert: true,
		},
		{
			Name:   "SomeReturnsFalse",
			Test:   func(p ProgrammingLanguage) bool { return p.Name == "PHP" },
			Assert: false,
		},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			b := collections.Some(programmingLanguages, testCase.Test)
			assert.Equal(t, testCase.Assert, b)
		})
	}
}

func TestEvery(t *testing.T) {

	for _, testCase := range []struct {
		Name   string
		Test   func(p ProgrammingLanguage) bool
		Assert bool
	}{
		{
			Name:   "EverySuccessful",
			Test:   func(p ProgrammingLanguage) bool { return p.GithutRank < 10 },
			Assert: true,
		},
		{
			Name:   "SomeReturnsFalse",
			Test:   func(p ProgrammingLanguage) bool { return p.GithutRank == 9 },
			Assert: false,
		},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			b := collections.Every(programmingLanguages, testCase.Test)
			assert.Equal(t, testCase.Assert, b)
		})
	}
}

func TestGroup(t *testing.T) {
	t.Run("SumOfRankings", func(t *testing.T) {
		out := collections.Group(programmingLanguages, func(p ProgrammingLanguage, amount int) int {
			return p.GithutRank + amount
		})

		assert.Equal(t, 22, out)

	})
}
