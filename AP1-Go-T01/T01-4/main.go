package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

// type Data struct {
// 	Doctor string
// 	VisitDate string
// }

type UserNotFoundError struct {
    Text string
}

type DoctorNotFoundError struct {
    Text string
}

func main() {
	journal := make(map[string]map[string][]string)
	getOperation(journal)
}

func getOperation(journal map[string]map[string][]string) {
	fmt.Println("Input: Save / GetHistory / GetLastVisit")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operation := scanner.Text()

	operationProcessing(journal, operation)
}

func journalEntry(name string, doctor string, date string, journal map[string]map[string][]string) {
	_, ok := journal[name]
	if !ok {
        journal[name] = make(map[string][]string)
	}

    _, ok = journal[name][doctor]
	if !ok {
        journal[name][doctor] = []string{}
    }

    journal[name][doctor] = append(journal[name][doctor], date)
	
	getOperation(journal)
}

func dateProcessing() string {
	scanner := bufio.NewScanner(os.Stdin)
	var date string

	for {
		fmt.Println("Enter date in YYYY-MM-DD format")
		scanner.Scan()

		date = scanner.Text()

		_, err := time.Parse("2006-01-02", date)
		if err == nil {
			break
		} else {
			fmt.Println("Invalid date format. Please try again.")
		}
	}
	return date
}

func doctorProcessing() string {
	scanner := bufio.NewScanner(os.Stdin)
	var doctor string
	for {
		fmt.Println("Enter a doctor")
		scanner.Scan()

		doctor = scanner.Text()
		checkDoctor := strings.Split(doctor, " ")

		if len(checkDoctor) == 1 {
			break
		}
	}
	return doctor
}

func nameProcessing() string {
	scanner := bufio.NewScanner(os.Stdin)
	var name string

	for {
		fmt.Println("Enter name in \"Иванов Иван Иванович\" format")
		scanner.Scan()

		name = scanner.Text()
		checkName := strings.Split(name, " ")

		if len(checkName) >= 3 {
			break
		}
	}
	return name
}

func operationProcessing(journal map[string]map[string][]string, operation string) {
	switch strings.ReplaceAll(strings.ToLower(operation), " ", "") {
	case "save":
		save(journal)
	case "gethistory":
		getHistory(journal)
	case "getlastvisit":
		getLastVisit(journal)
	default:
		getOperation(journal)
	}
}

func save(journal map[string]map[string][]string) {
	name := nameProcessing()
	doctor := doctorProcessing()
	date := dateProcessing()
	journalEntry(name, doctor, date, journal)
}

func getHistory(journal map[string]map[string][]string) {
	if len(journal) > 0 {
		name := nameProcessing()
		err := checkUser(name, journal)
		if err == nil {
			for doctor, visitDate := range journal[name] {
				for _, date := range visitDate {
					fmt.Println(doctor, date)
				}
			}
			getOperation(journal)
		} else {
			err.Error()
			getOperation(journal)
		}
	} else {
		fmt.Println("Journal is empty")
		getOperation(journal)
	}
}

func getLastVisit(journal map[string]map[string][]string) {
	if len(journal) > 0 {
		name := nameProcessing()
		err1 := checkUser(name, journal)
		if err1 == nil {
			doctor := doctorProcessing()
			err2 := checkDoctor(name, doctor, journal)
			if err2 == nil {
				lastDate := checkLastVisit(name, doctor, journal)
				fmt.Println(lastDate)
				getOperation(journal)
			} else {
				err2.Error()
				getOperation(journal)
			}
		} else {
			err1.Error()
			getOperation(journal)
		}
	} else {
		fmt.Println("Journal is empty")
		getOperation(journal)
	}
}

func checkLastVisit(name string, doctor string, journal map[string]map[string][]string) string {
	rangeDate := make([]string, len(journal))
	for _, el := range journal[name][doctor] {
		rangeDate = append(rangeDate, el)
	}
	
	sort.Slice(rangeDate, func(i, j int) bool {
        return rangeDate[i] > rangeDate[j] // Сортировка по убыванию
    })

	return rangeDate[0]
}

func (userError *UserNotFoundError) Error() string {
	return "user not found"
}

func (userError *DoctorNotFoundError) Error() string {
	return "doctor not found"
}

func checkUser(name string, journal map[string]map[string][]string) *UserNotFoundError {
	_, ok := journal[name]
	if !ok {
		return &UserNotFoundError{}
	}
	return nil
}

func checkDoctor(name string, doctor string, journal map[string]map[string][]string) *DoctorNotFoundError {
	_, ok := journal[name][doctor]
	if !ok {
		return &DoctorNotFoundError{}
	}
	return nil
}