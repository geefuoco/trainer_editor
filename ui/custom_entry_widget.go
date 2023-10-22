package ui

import (
	"fyne.io/fyne/v2/widget"
)

// CompletionEntry is an Entry with options displayed in a PopUpMenu.
type CustomEntry struct {
	widget.Entry
    Id            int
}

// NewCompletionEntry creates a new CompletionEntry which creates a popup menu that responds to keystrokes to navigate through the items without losing the editing ability of the text input.
func NewCustomEntry(id int) *CustomEntry {
	c := &CustomEntry{Id: id}
	c.ExtendBaseWidget(c)
	return c
}
