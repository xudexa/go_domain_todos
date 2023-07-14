package domain

import (
	"testing"
	"time"
)

func TestTotalMinutes(t *testing.T) {
	d := Duration{Day: 1, Hour: 1, Minute: 1}
	expected := 1440 + 60 + 1 // 1 jour + 1 heure + 1 minute
	if d.TotalMinutes() != expected {
		t.Errorf("TotalMinutes was incorrect, got: %d, want: %d.", d.TotalMinutes(), expected)
	}
}

func TestSortTodos(t *testing.T) {
	todos := Todos{
		&Todo{
			ID:           "1",
			Title:        "Test Todo 1",
			Importance:   Important,
			Urgency:      Urgent,
			ExpectedWork: Duration{Minute: 1},
			PlanningDate: time.Now().AddDate(0, 0, 1),
			State:        ToDo,
			Status:       Actived,
		},
		&Todo{
			ID:           "2",
			Title:        "Test Todo 2",
			Importance:   NotImportant,
			Urgency:      NotUrgent,
			ExpectedWork: Duration{Minute: 3},
			PlanningDate: time.Now(),
			State:        ToDo,
			Status:       Actived,
		},
	}

	SortTodos(todos)

	// Vérifier que la tâche la plus urgente et la plus importante est maintenant la première de la liste
	if todos[0].ID != "1" {
		t.Errorf("SortTodos did not sort todos as expected, got first ID: %s, want first ID: %s.", todos[0].ID, "1")
	}
}
