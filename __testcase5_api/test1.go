package main

import (
	"encoding/xml"
	"fmt"
	// "io/ioutil"
	"io"
	"os"
)

type Post struct {
	XMLName 	xml.Name 	`xml:"post"`
	Id 			string 		`xml:"id,attr"`
	Content 	string 		`xml:"content"`
	Author 		Author 		`xml:"author"`
	Xml 		string 		`xml:",innerxml"`
	Comments    []Comment 	`xml:"comments>comment"`
}

type Comment struct {
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:author`
}

type Author struct {
	Id string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error open xml file:", err)
	}
	defer xmlFile.Close()

	// //1 无法高效处理以流方式传输的xml文件以及体积较大的xml文件
	// xmlData, err := ioutil.ReadAll(xmlFile)
	// if err != nil {
	// 	fmt.Println("err reading xml data", err)
	// 	return
	// }
	// var post Post
	// xml.Unmarshal(xmlData, &post)
	// fmt.Println(post)

	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error decoding xml into token:", err)
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
			}
		}
	}

}