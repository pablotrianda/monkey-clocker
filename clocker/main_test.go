package clocker

import "testing"

func TestNewClocker(t *testing.T) {
	name := "Pablo"
	busyHours := []int{12, 15, 16}

	newClocker, err := NewClocker(name, busyHours)

	if err != nil {
		t.Fatal("fail the new clocker cant be created")
	}

	if newClocker.Name != name {
		t.Fatalf("expected %s, got %s", name, newClocker.Name)
	}

	if len(newClocker.BusyHours) != len(busyHours) {
		t.Fatalf("expected %d, got %d", len(busyHours), len(newClocker.Name))
	}
}
