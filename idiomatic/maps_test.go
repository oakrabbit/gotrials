package idiomatic

import (
	"log"
	"testing"
)

func TestDumbProcedure(t *testing.T) {
	MapPlay()
	log.Println("Dumb Test")
}

func TestAddItem(t *testing.T) {
	m := MapPlay()
	m.addItem(5, "five")

	_, ok := m[5]

	if !ok {
		t.Fail()
	}
	log.Println("item:", m)
}

func TestDeleteItem(t *testing.T) {
	m := MapPlay()

	m.deleteItem(2)

	_, ok := m[2]
	if ok {
		t.Fail()
	}
	log.Println("item:", m)
}

func TestPurgeData(t *testing.T) {
	m := MapPlay()
	m.purge()

	if len(m) > 0 {
		t.Fail()
	}
	log.Println("item:", m)

}
