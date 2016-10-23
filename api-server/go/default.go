package peering

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Default struct {
}

func GraphGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func RoutePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ToptalkersNGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	session, err := mgo.Dial("localhost")
	c := session.DB("peering").C("bandwith")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	pipe := c.Pipe([]bson.M{{"$group": bson.M{"_id": "$dstip", "count": bson.M{"$sum": 1}}}, bson.M{"$sort": bson.M{"count": -1}}, bson.M{"$limit": 10}})
	var result []bson.M
	pipe.All(&result)
	json, _ := bson.MarshalJSON(&result)

	w.Write(json)
}
