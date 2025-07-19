package tests

import (
	"github.com/prithvitewatia/gocache/src"
	"testing"
	"time"
)

func TestSetGet(t *testing.T) {
	c := src.NewCache()
	defer c.Close()

	c.Set("key1", "value1", 0)
	val, exists := c.Get("key1")
	if !exists || val != "value1" {
		t.Fatalf("Expected 'value1' got '%v'", val)
	}
}

func TestExpiration(t *testing.T) {
	c := src.NewCache()
	defer c.Close()

	c.Set("key1", "value1", 5*time.Millisecond)
	time.Sleep(100 * time.Millisecond)
	_, exists := c.Get("key1")
	if exists {
		t.Fatalf("Expectd value not to be present for 'key1'")
	}
}

func TestDelete(t *testing.T) {
	c := src.NewCache()
	defer c.Close()
	c.Set("key1", "value1", 0)
	c.Delete("key1")

	_, exists := c.Get("key1")
	if exists {
		t.Fatalf("Expectd value for not to be present for 'key1'")
	}

}

func TestOverwrite(t *testing.T) {
	c := src.NewCache()
	defer c.Close()

	c.Set("dup", "first", 0)
	c.Set("dup", "second", 0)

	val, ok := c.Get("dup")
	if !ok || val != "second" {
		t.Fatalf("expected second, got %v", val)
	}
}
