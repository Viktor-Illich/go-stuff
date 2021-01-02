package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	type Address struct {
		City    string `xml:"address>city"`
		Country string `xml:"address>country"`
	}
	type Employee struct {
		XMLName   xml.Name `xml:"employee"`
		Id        int      `xml:"id,attr"`
		Uniq      int      `xml:"uniq,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Initials  string   `xml:"name>initials"`
		Height    float32  `xml:"height,omitempty"`
		Address
		Comment string `xml:",comment"`
	}

	r := &Employee{Id: 7, Uniq: 4, FirstName: "Viktor", LastName: "Lazhevskyi", Initials: "VLI"}
	r.Comment = "Dev"
	r.Address = Address{"SomeWhere 12", "12312, Greece"}

	output, err := xml.MarshalIndent(r, "  ", "    ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	output = []byte(xml.Header + string(output))
	os.Stdout.Write(output)
	os.Stdout.Write([]byte("\n"))
}
