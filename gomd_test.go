package gomd

import (
	"fmt"
	"testing"
)

var col = Collection{"testDB", "contacts"}
var p = Person{"Jack_Green", "9987650"}

// C create/insert
func TestInsert(t *testing.T) {
	insert(p, col)
	pb := findByName(p.Name, col)
	if pb.Name != p.Name || pb.Phone != p.Phone {
		t.Error("insert failed")
	}
	fmt.Println("Insert Result")
	fmt.Println(pb)
}

// R read/find
func TestFindByName(t *testing.T) {
	p := findByName(p.Name, col)
	if p.Name == "" || p.Phone == "" {
		t.Error("find by name test failed")
	}
	fmt.Println("Find Result")
	fmt.Println(p)

}

// U update
func TestUpdate(t *testing.T) {
	p := Person{"Jack_Green", "121212122"}
	update(p, col)
	ub := findByName(p.Name, col)
	fmt.Println("Update Result")
	fmt.Println(ub)
}

// D delete
func TestDeleteData(t *testing.T) {
	deleteData(p, col)
}
