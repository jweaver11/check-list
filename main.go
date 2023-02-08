package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor   int              // which to-do list item our curser is pointing at
	choices  []string         // items on the to-do list
	selected map[int]struct{} // which to-do items are selected
}

// Function that creates and returns our initial model.
// Could easily defined elsewhere too as a variable
func initialModel() model {
	return model{
		//Our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

var newTask string

func (m model) NewTask(newTask string) tea.Model {
	fmt.Println("\n\n\nPlease enter your new task:")

	fmt.Println(newTask)

	//choices[4] = newTask
	return model{
		choices: []string{"Eat ass", "sk8 fast", newTask},
	}
}

// Can return a command (tea.Cmd) but for now means "no command"
func (m model) Init() tea.Cmd {
	// returns 'nil', which means "no I/O" right now
	return nil
}

// Called when "things happen"
// Makes changes to updated model and returns updated model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Declares msg as a switch of msg.(type)
	switch msg := msg.(type) {

	// Checks if input is a key press
	case tea.KeyMsg:
		// Gets the actual key pressed
		switch msg.String() {

		// Quits program
		case "ctrl+c", "q":
			fmt.Println("\n\nOk bye bye!")
			return m, tea.Quit

		// Moves curser up
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		//Moves curser down
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		//Toggle selected state for item the curser is pointing at
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		//Add a new item to the list
		case "n":
			fmt.Println("you pressed n")
			NewTask("This is a new task")

		}
	}

	return m, nil
}

// Renders the view model and returns it as a string
func (m model) View() string {
	// Header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {
		// Is the cursor pointing at our choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor
		}
		// Is this choice selected?
		checked := " " // Not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // Selected
		}
		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	// The footer
	s += "Press 'space' to select a task and arrows to navigate"
	helpBar := "\n\n\n\n\n\nFor real tho bruv"

	//helpView := m.help.View(m.keys)

	// Sends the UI for rendering
	return s + helpBar
}

func main() {
	// Sets P as a new tea program using the intial model function
	// The underlying UI uses our already created functions
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
