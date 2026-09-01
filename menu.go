package main

import "fmt"

type MenuItem struct {
	Label  string
	Action func()
}

type Menu struct {
	Title string
	Items []MenuItem
}

func (m Menu) Show() {
	showHeader(m.Title)
	for i, item := range m.Items {
		fmt.Printf("[%d] %s\n", i+1, item.Label)
	}
}

func (m Menu) Run() {
	for {
		m.Show()
		choice := ReadInt("Option", 1, len(m.Items))
		m.Items[choice-1].Action()
	}
}

func showHeader(title string) {
	fmt.Printf("<-------- %s -------->\n", title)
}
