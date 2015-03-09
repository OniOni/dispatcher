package store

import (
	"fmt"
	"testing"
)

func Test_AddSubscriber(t *testing.T) {
	store, error := NewStore()
	if error != nil {
		t.Error("Error creating store.")
	}
	defer store.Close()
	store.db.Clear()

	fmt.Println("Store db", store.db)

	store.AddSubsriber("lol", "one")
	if error != nil {
		t.Error("Error setting value.")
	}

	value, error := store.GetSubscribers("lol")
	if error != nil {
		t.Error("Error fecthing value.")
	}

	if value != "one" {
		t.Error("Value was not as expected.", value)
	}
}
