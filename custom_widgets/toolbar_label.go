package custom_widgets

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/widget"
)

type ToolbarLabel struct {
    *widget.Label
}

func NewToolbarLabel(label string) widget.ToolbarItem {
    l := widget.NewLabelWithStyle(label, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
    l.MinSize()
    return &ToolbarLabel{l}
}

func (t *ToolbarLabel) ToolbarObject() fyne.CanvasObject {
    return t.Label
}
