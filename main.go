package main

import (
	"io/ioutil"
	"log"
	"os"

	pb "github.com/daplho/addressbook/tutorial"
	"github.com/gogo/protobuf/proto"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

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
