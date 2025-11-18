package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type player struct {
	name string
	birthday string
	age int
	nationality string
	current_club string
}

// calculates the age in years based on a birth date and a reference date.
func calcAge(birthday, referenceDay time.Time) int {
	age := referenceDay.Year() - birthday.Year()
	
	// Adjust age if the birthday hasn't occurred yet in the current year
	if referenceDay.YearDay() < birthday.YearDay() {
		age--
	}
	return age
}

func strToDate(date string) (time.Time, error) {
	format_date := "2006-01-02"
	t, err := time.Parse(format_date, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}, err
	}
	
	return t, nil
}

func newPlayer(name string, birthday string, nationality, current_club string) *player {
	date, _ := strToDate(birthday)
	age := calcAge(date, time.Now())
	p := player{name: name, birthday: birthday, age: age, nationality: nationality, current_club: current_club}
	return &p
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var birthday string
	
	fmt.Println("Please input player bio:")

	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	fmt.Print("Birthday (YYYY-MM-DD): ")
	fmt.Scanf("%s\n", &birthday)
	
	fmt.Print("Nationality: ")
	nationality, _ := reader.ReadString('\n')
	nationality = strings.TrimSpace(nationality)
	
	fmt.Print("Club: ")
	current_club, _ := reader.ReadString('\n')
	current_club = strings.TrimSpace(current_club)
	
	t, _ := strToDate(birthday)
	player := newPlayer(name, birthday, nationality, current_club)
	
	fmt.Println("---- Personal information ----")
	fmt.Println("Name: ", player.name)
	fmt.Printf("Birthday: %d %s %d (age %d)\n", t.Day(), t.Month(), t.Year(), player.age)
	fmt.Println("Nationality: ", player.nationality)
	fmt.Println("Club: ", player.current_club)
}