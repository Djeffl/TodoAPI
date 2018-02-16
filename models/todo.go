package models

import "time"

type Todo struct {
	ID           string    `json:"id"`
	Name         string    `json:"name";`
	CreationDate time.Time `json:"creationDate"`
	FinishTime   time.Time `json:"finishDate"`
	Completed    bool      `json:"completed"`
	HasTime      bool      `json:"hasHours"`
	HoursLeft    time.Time `json:"hoursLeft"`
}
