package main

import (
	pb "github.com/daplho/addressbook/proto/tutorial"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	"log"
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
