package storage

import (
	"testing"
)

func TestResource(t *testing.T) {
	storage := map[string]int{
		"wood": 5,
	}

	inventory := make(map[string]int)
	resource := "wood"
	amount := 3
	wantAmount := 2

	TakeResource(storage, inventory, resource, amount)

	leftAmount, ok := storage[resource]
	if !ok {
		t.Errorf("expected resource %v exist in storage", resource)
	}

	if leftAmount != wantAmount {
		t.Errorf("expected %d amount left in storage, got %d", wantAmount, leftAmount)
	}

	wantAmount = 3
	leftAmount, ok = inventory[resource]
	if !ok {
		t.Errorf("expected resource %v exist in inventory", resource)
	}

	if leftAmount != wantAmount {
		t.Errorf("expected %d amount left in inventory, got %d", wantAmount, leftAmount)
	}
}

func TestStorageCheck(t *testing.T) {
	storage := map[string]int{
		"wood": 0,
	}
	resource := "wood"
	StorageCheck(storage, resource)

	_, ok := storage[resource]

	if ok {
		t.Errorf("expected resource %v unexist", resource)
	}
}

func TestFixZone(t *testing.T) {
	tests := []struct {
		name        string
		required    map[string]int
		inventory   map[string]int
		requiredOk  bool
		inventoryOk bool
		wantAmount  int
	}{
		{
			name: "without deleting items",
			required: map[string]int{
				"wood": 5,
			},

			inventory: map[string]int{
				"wood": 7,
			},

			requiredOk:  false,
			inventoryOk: true,
			wantAmount:  2,
		},

		{
			name: "with deleting items",
			required: map[string]int{
				"wood": 5,
			},

			inventory: map[string]int{
				"wood": 5,
			},

			requiredOk:  false,
			inventoryOk: false,
			wantAmount:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FixZone(tt.required, tt.inventory)

			_, ok := tt.required["wood"]

			if ok != tt.requiredOk {
				t.Error("expected unexisting resource in required resources")
			}
			amount, ok := tt.inventory["wood"]
			if ok != tt.inventoryOk {
				t.Fatal("expected existing resource in inventory, got unexisting")
			}

			if amount != tt.wantAmount {
				t.Errorf("expected amount %d, got %d", tt.wantAmount, amount)
			}

		})
	}
}

func TestAccessCard(t *testing.T) {
	tests := []struct {
		name       string
		inventory  map[string]int
		step       int
		wantAmount int
		wantExist  bool
	}{
		{
			name:       "success case",
			inventory:  make(map[string]int),
			step:       6,
			wantAmount: 1,
			wantExist:  true,
		},

		{
			name:       "smalest edge case",
			inventory:  make(map[string]int),
			step:       4,
			wantAmount: 1,
			wantExist:  true,
		},

		{
			name:       "biggest edge case",
			inventory:  make(map[string]int),
			step:       7,
			wantAmount: 1,
			wantExist:  true,
		},

		{
			name:       "less than able to be",
			inventory:  make(map[string]int),
			step:       3,
			wantAmount: 0,
			wantExist:  false,
		},

		{
			name:       "more than able to be",
			inventory:  make(map[string]int),
			step:       8,
			wantAmount: 0,
			wantExist:  false,
		},
	}
	//make table driven and fix zone test

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AccessCard(tt.step, tt.inventory)

			amount, ok := tt.inventory["accessCard"]

			if ok != tt.wantExist {
				t.Fatalf("expected existing %v, got %v", tt.wantExist, ok)
			}

			if amount != tt.wantAmount {
				t.Errorf("expected amount %d, got %d", tt.wantAmount, amount)
			}
		})
	}
}

func TestSpecTool(t *testing.T) {
	inventory := make(map[string]int)
	SpecTool(inventory)

	_, ok := inventory["tool"]
	if !ok {
		t.Error("expected tool exist")
	}
}
