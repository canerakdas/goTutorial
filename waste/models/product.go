package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db  *mgo.Database

type Product struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Price float64 `json:"price" bson:"price"`
	Description string `json:"description" bson:"description"`
	Categories []string `json:"categories" bson:"categories"`
	Colors []string `json:"colors" bson:"colors"`
	Sizes []string `json:"sizes" bson:"sizes"`
	Params map[string]string `json:"params" bson:"params"`
}

func Des(){

	// Models & structs need to db connection
	session, err := mgo.Dial("localhost")

	var result = make([]Product,1,5)

	var conn = session.DB("myDb").C("user")

	err = conn.Find(nil).Sort("-timestamp").All(&result)

	fmt.Printf("%v",result)

	err = conn.Find(bson.M{"_id":bson.ObjectIdHex("5c1a5c955dad5b64b232703e")}).Sort("-timestamp").All(&result)

	change := bson.M{"$set": bson.M{"title":"armut","price":22.90}}

	err = conn.Update(bson.M{"_id":bson.ObjectIdHex("5c1a5c955dad5b64b232703e")}, change)

	fmt.Printf("\n%v Armut: \n",result)

	err = conn.Find(bson.M{"_id":bson.ObjectIdHex("5c1a5c955dad5b64b232703e")}).Sort("-timestamp").One(&result)

	err = conn.Remove(bson.M{"_id":bson.ObjectIdHex("5c1a5e775dad5b650a6d93e8")})

	err = conn.Find(bson.M{"_id":bson.ObjectIdHex("5c1a5e775dad5b650a6d93e8")}).Sort("-timestamp").One(&result)

	fmt.Println(result)
	fmt.Println(err)
}