package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial("localhost")
	c := session.DB("peering").C("bandwith")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	/*	db.myCollection.aggregate(
		    {$group : {_id : "$user", "count" : {$sum : 1}}},
			    {$sort : {"count" : -1}},
				    {$limit : 10}
				)*/

	pipe := c.Pipe([]bson.M{{"$group": bson.M{"_id": "$dstip", "count": bson.M{"$sum": 1}}}, bson.M{"$sort": bson.M{"count": -1}}, bson.M{"$limit": 10}})
	var result []bson.M
	pipe.All(&result)

	fmt.Printf("%v\n", result)
}
