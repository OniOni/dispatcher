package store

import (
	"testing"
)

func Test_AddSubscriber(t *testing.T) {
	store, error := NewStore()
	if error != nil {
		t.Error("Error creating store.")
	}
	defer store.Close()
	store.db.Clear()

	store.AddSubsriber("lol", "one")
	if error != nil {
		t.Error("Error setting value.")
	}

	subs, error := store.GetSubscribers("lol")
	if error != nil {
		t.Error("Error fecthing value.")
	}

	if subs[0] != "one" {
		t.Error("Value was not as expected.", subs[0])
	}
}

func Test_IsSubscribed(t *testing.T) {
	store, error := NewStore()
	if error != nil {
		t.Error("Error creating store.")
	}
	defer store.Close()
	store.db.Clear()

	store.AddSubsriber("lol", "one")
	if error != nil {
		t.Error("Error setting value.")
	}

	is, error := store.IsSubscribed("lol", "one")
	if !is {
		t.Error("Should be subscirbed.")
	}
}

func Test_hasKey(t *testing.T) {
	store, error := NewStore()
	if error != nil {
		t.Error("Error creating store.")
	}
	defer store.Close()
	store.db.Clear()

	store.AddSubsriber("key", "one")
	if error != nil {
		t.Error("Error setting value.")
	}

	if !store.HasKey("key") {
		t.Error("Should have contained key.")
	}
}
