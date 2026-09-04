package main

import "fmt"

type MenuItem struct {
	Label  string
	Action func() bool
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
		keepGoing := m.Items[choice-1].Action()
		if !keepGoing {
			return
		}
	}
}

func showHeader(title string) {
	fmt.Printf("\n<-------- %s -------->\n", title)
}

func handleBack() bool {
	return false
}
