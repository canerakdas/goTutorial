package util

import (
	"gopkg.in/mgo.v2"
)

var session, err = mgo.Dial("localhost")
var Conn = session.DB("myDb").C("user")
