package main

import (
	"io/ioutil"
	"fmt"
)

func main()  {
	p1 := Page{Title:"MyTitle",Body:[]byte("My Body")}
	p1.save()

	p2, _ := loadPage("MyTitle")
	fmt.Println(string(p2.Body))
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, error := ioutil.ReadFile(filename)

	if (error != nil) {
		return nil, error
	}

	return &Page{Title: title, Body: body}, nil
}
