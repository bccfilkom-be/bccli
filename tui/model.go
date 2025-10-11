package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type step int

const (
	stepPackage step = iota
	stepFramework
	stepDriver
	stepFinish
)

type model struct {
	step              step
	projectName       textinput.Model
	cursor            int
	frameworkChoices  []string
	selectedFramework int
	driverChoices     []string
	selectedDriver    int
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "project name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 60

	return model{
		step:             stepPackage,
		projectName:      ti,
		frameworkChoices: frameWorkChoises,
		driverChoices:    databaseDriverChoises,
	}
}
