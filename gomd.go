package gomd

import (
	// "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name  string
	Phone string
}
type Collection struct {
	DB   string
	Name string
}

const (
	URL = "127.0.0.1:27017"
)

func insert(p Person, col Collection) {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB(col.DB).C(col.Name)
	err = c.Insert(p)
	if err != nil {
		log.Fatal(err)
	}
}

func findByName(name string, col Collection) Person {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(col.DB).C(col.Name)
	result := Person{}
	err = collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func update(p Person, col Collection) {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	collection := session.DB(col.DB).C(col.Name)
	err = collection.Update(bson.M{"name": p.Name}, bson.M{"$set": bson.M{"phone": p.Phone}})
	if err != nil {
		log.Fatal(err)
	}
}

func deleteData(p Person, col Collection) {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	collection := session.DB(col.DB).C(col.Name)
	err = collection.Remove(bson.M{"name": p.Name})
}
