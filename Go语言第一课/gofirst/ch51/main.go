package main

import (
	. "ch51/model"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	test := &Student{
		Name: "andyron",
		Male: true,
		Scores: []int32{
			100, 99, 98,
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}
}
