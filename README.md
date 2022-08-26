# Slice-Helper

It's just a convinience to work with slices.
<br><br>

## Examples
Consider these data for the next examples:
```go
var programmingLanguages = []ProgrammingLanguage{
	{Id: 1, Name: "JavaScript", ReleasedYear: 1995, GithutRank: 1, PyplRank: 3},
	{Id: 3, Name: "Java", ReleasedYear: 1995, GithutRank: 3, PyplRank: 2},
	{Id: 2, Name: "Python", ReleasedYear: 1991, GithutRank: 2, PyplRank: 1},
	{Id: 4, Name: "TypeScript", ReleasedYear: 2012, GithutRank: 7, PyplRank: 10},
	{Id: 5, Name: "C#", ReleasedYear: 2000, GithutRank: 9, PyplRank: 4},
}
```

### **Find**
```go
p := collections.Find(programmingLanguages, func(p ProgrammingLanguage) bool {
		return p.Id == 4
})
fmt.Println(*p)
```
#### Output
```
{4 TypeScript 2012 7 10}
```
<br>

### **Filter**
```go
languages := collections.FilterArray(programmingLanguages, func(p ProgrammingLanguage) bool {
		return p.GithutRank > 6
})
fmt.Println(languages)
```
#### Output
```
[{4 TypeScript 2012 7 10} {5 C# 2000 9 4}]
```
<br>

### **Map**
```go
langNames := collections.MapArray(programmingLanguages, func(p ProgrammingLanguage) string {
		return p.Name
})
fmt.Println(langNames)
```
#### Output
```
[JavaScript Java Python TypeScript C#]
```
<br>

### **Some**
```go
shouldBeTrue := collections.Some(programmingLanguages, func(p ProgrammingLanguage) bool {
	return p.GithutRank == 7
})

shouldBeFalse := collections.Some(programmingLanguages, func(p ProgrammingLanguage) bool {
	return p.GithutRank == 4
})
fmt.Println("shouldBeTrue: ", shouldBeTrue, "shouldBeFalse: ", shouldBeFalse)
```
#### Output
```
shouldBeTrue:  true shouldBeFalse:  false
```
<br>

### **Every**
```go
shouldBeTrue = collections.Every(programmingLanguages, func(p ProgrammingLanguage) bool {
	return p.GithutRank < 10
})

shouldBeFalse = collections.Every(programmingLanguages, func(p ProgrammingLanguage) bool {
    return p.GithutRank > 4
})
fmt.Println("shouldBeTrue: ", shouldBeTrue, "shouldBeFalse: ", shouldBeFalse)

```
#### Output
```
shouldBeTrue:  true shouldBeFalse:  false
```
<br>

### **Group**
```go
sumOfRankings := collections.Group(programmingLanguages, func(p ProgrammingLanguage, amount int) int {
	return p.GithutRank + amount
})
fmt.Println("Some of Rankings: ", sumOfRankings)
```
#### Output
```
Sum of Rankings:  22
```
<br>