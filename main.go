package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type status int

const (
	todo status = iota
	inProgress
	done
)

// custom item

type Task struct {
	status      status
	title       string
	description string
}

// implement the list.Item interface

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

// main model

type Model struct {
	list list.Model
}

func New() *Model {
	return &Model{}
}

// TODO: call this on tea.WindowSizeMsg
func (m *Model) initList(width, height int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "To Do"
	m.list.SetItems([]list.Item{
		Task{status: todo, title: "Write a blog post", description: "Write a blog post about Bubble Tea"},
		Task{status: todo, title: "Walk the dog", description: "Take the dog for a walk around the block"},
		Task{status: todo, title: "Buy groceries", description: "Buy milk, eggs, and bread"},
	})
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width, msg.Height)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	return m.list.View()
}

func main() {
	m := New()
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
