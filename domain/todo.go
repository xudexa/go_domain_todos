package domain

import (
	"sort"
	"time"
)

type Importance int
type Urgency int
type State string
type Status string

const (
	Important    Importance = 1
	NotImportant Importance = 0
	Urgent       Urgency    = 1
	NotUrgent    Urgency    = 0
	ToDo         State      = "todo"
	InProgress   State      = "in progress"
	Done         State      = "done"
	Actived      Status     = "actived"
	Archived     Status     = "archived"
)

type Duration struct {
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type Todo struct {
	ID             string     `json:"id"`
	Title          string     `json:"title"`
	Description    string     `json:"description,omitempty"`
	Importance     Importance `json:"importance"`
	Urgency        Urgency    `json:"urgency"`
	PlanningDate   time.Time  `json:"planning_date,omitempty"`
	ExpectedWork   Duration   `json:"expected_work"`
	ActualWork     Duration   `json:"actual_work,omitempty"`
	State          State      `json:"state"`
	Status         Status     `json:"status"`
	AssignedPerson string     `json:"assigned_person,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	LastModifiedAt time.Time  `json:"last_modified_at"`
}

// TodoUseCases définit tous les cas d'utilisation possibles sur un Todo.
type TodoUseCases interface {
	Add(todo *Todo) error
	Get(id string) (*Todo, error)
	Update(todo *Todo) error
	Delete(id string) error
	GetAll() (Todos, error)
	GetAllArchived() (Todos, error)
}

// TodoRepository définit toutes les actions possibles à réaliser sur le storage de Todo.
type TodoRepository interface {
	AddTodo(todo *Todo) error
	GetTodo(id string) (*Todo, error)
	UpdateTodo(todo *Todo) error
	DeleteTodo(id string) error
	GetAllTodos() (Todos, error)
	GetAllArchivedTodos() (Todos, error)
}

type Todos []*Todo

func (t Todos) Len() int {
	return len(t)
}

// commentaire test
func (t Todos) Less(i, j int) bool {
	if t[i].Status == Archived {
		return false
	}
	if t[j].Status == Archived {
		return true
	}
	if t[i].ExpectedWork.TotalMinutes() < 2 {
		return true
	}
	if t[j].ExpectedWork.TotalMinutes() < 2 {
		return false
	}
	if t[i].PlanningDate != t[j].PlanningDate {
		return t[i].PlanningDate.Before(t[j].PlanningDate)
	}
	if t[i].Urgency != t[j].Urgency {
		return t[i].Urgency < t[j].Urgency
	}
	return t[i].Importance < t[j].Importance
}

func (t Todos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (d Duration) TotalMinutes() int {
	return d.Day*1440 + d.Hour*60 + d.Minute
}

func SortTodos(todos Todos) {
	sort.Sort(todos)
}
