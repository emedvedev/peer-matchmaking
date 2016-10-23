package main

import (
	"../stoneflow"
	sw "./go"
	"encoding/binary"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type Bandwidth struct {
	SrcIP uint32
	DstIP uint32
}

func SetupSFlow() *stoneflow.StoneFlow {
	sflow := stoneflow.CreateSFlow()
	go ConsumeSFlow(sflow)
	return sflow
}

func ConsumeSFlow(sflow *stoneflow.StoneFlow) {
	session, err := mgo.Dial("localhost")
	c := session.DB("peering").C("bandwith")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	for {
		frame := <-sflow.ReadIn
		log.Printf("Got a frame...Doing something with it: %v", frame)
		srcip := binary.BigEndian.Uint32(frame.SrcIP)
		dstip := binary.BigEndian.Uint32(frame.DstIP)
		log.Printf("Adding to mongo")
		err = c.Insert(&Bandwidth{srcip, dstip})
		if err != nil {
			log.Fatal(err)
		}

		result := Bandwidth{}
		err = c.Find(bson.M{"srcip": srcip}).One(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("DstIP:", result.DstIP)
	}
}

func main() {
	log.Printf("Server started")
	SetupSFlow()
	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
