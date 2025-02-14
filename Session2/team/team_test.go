package team

import "testing"

/*
Person:
1. create struct Person with the following fields:
  - Name	(string)
  - Age		(int)
  - IsAlive	(bool)

2. implement function NewPerson:
  - returns a pointer to a newly created Person.
  - a new person must be alive.

3. implement function IncrementAge:
  - increments the Age field of the given person by 1.

LeagueType:
1. create int type LeagueType.

2. define three enumeration constants with that type:
	- 1: League1
	- 2: League2
	- 3: PremierLeague

Team:
1. create struct Team with the following fields:
  - Name	(string)
  - League	(LeagueType)
  - Captain	(Person)

2. implement function NewTeam:
  - returns a pointer to a newly created Team.
  - the team's captain should point to the provided person.

3. implement function ChangeCaptain:
  - changes the team’s captain to another person.
*/

func TestPerson(t *testing.T) {
	p := NewPerson("Alice", 25)

	if p == nil {
		t.Fatalf("expected NewPerson to return a non-nil pointer")
	}
	if p.Name != "Alice" {
		t.Errorf("expected Name to be Alice, got %q", p.Name)
	}
	if p.Age != 25 {
		t.Errorf("expected Age to be 25, got %d", p.Age)
	}
	if !p.IsAlive {
		t.Errorf("expected new person to be alive")
	}

	IncrementAge(p)
	if p.Age != 26 {
		t.Errorf("expected Age to be 26 after IncrementAge, got %d", p.Age)
	}

	IncrementAge(p)
	if p.Age != 27 {
		t.Errorf("expected Age to be 27 after second IncrementAge, got %d", p.Age)
	}
}

func TestTeam(t *testing.T) {
	captain := NewPerson("Charlie", 30)
	team := NewTeam("Gophers", League1, captain)

	if team == nil {
		t.Fatalf("expected NewTeam to return a non-nil pointer")
	}
	if team.Name != "Gophers" {
		t.Errorf("expected Name to be Gophers, got %q", team.Name)
	}
	if team.League != League1 {
		t.Errorf("expected League to be 1 (League1), got %d", team.League)
	}
	if *team.Captain != *captain {
		t.Errorf("expected Captain to be %+v, got %+v", captain, team.Captain)
	}

	newCaptain := NewPerson("Edward", 28)

	ChangeCaptain(team, newCaptain)
	if *team.Captain != *newCaptain {
		t.Errorf("expected Captain to be changed to %+v, got %+v", newCaptain, team.Captain)
	}

	IncrementAge(team.Captain)
	if newCaptain.Age != 29 {
		t.Errorf("expected Age to be 29 after IncrementAge, got %d", newCaptain.Age)
	}
}
