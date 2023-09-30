package ebird

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	key := "wololo"
	client, err := NewClient(key)
	if err != nil {
		t.Fatal(err)
	}
	if client.APIKey != key {
		t.Errorf("APIKey %s; want %s", client.APIKey, key)
	}
}
