package todo

import "time"

type Todo struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	CreationDate time.Time `json:"creationDate,omitempty"`
	FinishTime   time.Time `json:"finishDate"`
	Completed    bool      `json:"completed"`
}
