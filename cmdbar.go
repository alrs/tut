package main

import (
	"strings"

	"github.com/rivo/tview"
)

func NewCmdBar(app *App) *CmdBar {
	c := &CmdBar{
		app:   app,
		Input: tview.NewInputField(),
	}

	c.Input.SetFieldBackgroundColor(app.Config.Style.Background)
	c.Input.SetFieldTextColor(app.Config.Style.Text)

	return c
}

type CmdBar struct {
	app   *App
	Input *tview.InputField
}

func (c *CmdBar) GetInput() string {
	return strings.TrimSpace(c.Input.GetText())
}

func (c *CmdBar) ClearInput() {
	c.Input.SetText("")
}