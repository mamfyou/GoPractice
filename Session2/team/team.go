package team

type LeagueType int

const (
	League1 LeagueType = iota + 1
	League2
	PremiereLeague
)


type Person struct {
	Name string
	Age  int
	IsAlive bool
}

type Team struct {
	Name string
	League LeagueType
	Captain *Person
}

func NewPerson(name string, age int) *Person {
	new_person := new(Person)
	new_person.IsAlive = true
	new_person.Age = age
	new_person.Name = name

	return new_person
}

func IncrementAge(person *Person) {
	person.Age += 1
}

func NewTeam(name string, league LeagueType, captain *Person) *Team {
	new_team := new(Team)
	new_team.Captain = captain
	new_team.Name = name
	new_team.League = league

	return new_team
}

func ChangeCaptain(team *Team, captain *Person) {
	team.Captain = captain
}