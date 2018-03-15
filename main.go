package main

import (
	"io/ioutil"
	"log"

	pb "github.com/daplho/addressbook/tutorial"
	"github.com/gogo/protobuf/proto"
)

func main() {
	book := &pb.AddressBook{}

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book: ", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book: ", err)
	}
}
