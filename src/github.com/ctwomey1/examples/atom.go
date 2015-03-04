package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Entry struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Updated   string `xml:"updated"`
	Published string `xml:"published"`
	Category  string `xml:"category"`
	Summary   string `xml:"summary"`
}
type Feed struct {
	XMLName  xml.Name `xml:"feed"`
	Title    string   `xml:"title"`
	Subtitle string   `xml:"subtitle"`
	Id       string   `xml:"id"`
	Updated  string   `xml:"updated"`
	Logo     string   `xml:"logo"`
	Icon     string   `xml:"icon"`
	Rights   string   `xml:"rights"`
	Entries  []Entry  `xml:"entry"`
}

func main() {
	dat, err := ioutil.ReadFile("mb-38_e.xml")
	if err != nil {
		panic(err)
	}

	var f Feed
	err2 := xml.Unmarshal(dat, &f)
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("Title: %q\n", f.Title)
	fmt.Printf("Subtitle: %q\n", f.Subtitle)
	fmt.Printf("Id: %q\n", f.Id)
	fmt.Printf("Updated: %q\n", f.Updated)
	fmt.Printf("Logo: %q\n", f.Logo)
	fmt.Printf("Icon: %q\n", f.Icon)
	fmt.Printf("Rights: %q\n", f.Rights)
	fmt.Printf("Entries: %v\n", f.Entries)
}
