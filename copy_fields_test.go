package gotoo

import (
	"testing"
	"time"
)

type Source struct {
	ID int
	Name string
	CreatedAt time.Time
}

type Target struct {
	ID int
	Name string
	CreatedAt time.Time
}

func TestCopyFields(t *testing.T) {
	s := Source{1, "name", time.Now()}
	target_i, err := CopyFields(s, Target{})
	if err != nil {
		t.Errorf("%s", err)
	}
	target, ok := target_i.(*Target)
	if ok == false {
		t.Errorf("Don't match type")
	}

	if target.ID != s.ID {
		t.Errorf("Wrong value: target id is expected to be %d, got %d", s.ID, target.ID)
	}
	if target.Name != s.Name {
		t.Errorf("Wrong value: target id is expected to be %s, got %s", s.Name, target.Name)
	}
	if target.CreatedAt != s.CreatedAt {
		t.Errorf("Wrong value: target id is expected to be %s, got %s", s.CreatedAt, target.CreatedAt)
	}
}
