package custom_widgets

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/theme"
)
// AutoComplete is a custom widget that combines Entry and Select for auto-completion
type AutoComplete struct {
    widget.Entry
    popup       *widget.PopUp
    optionList  *optionList
    itemHeight  float32
    Options     []string
    oldText     string
}

func NewAutoComplete(options []string) *AutoComplete {
    a := &AutoComplete{Options: options}
    a.ExtendBaseWidget(a)
    return a
}

func (a* AutoComplete) Refresh() {
    a.Entry.Refresh()
    if a.optionList != nil {
        a.optionList.SetOptions(a.Options)
    }
}

func (a* AutoComplete) SetOptions(options []string) {
    a.Options = options
    a.Refresh()
}

func (a* AutoComplete) HideCompletion() {
    if a.popup != nil {
        a.popup.Hide()
    }
}

func (a* AutoComplete) ShowCompletion() {
    if len(a.Options) == 0 {
        a.HideCompletion()
        return
    }
    if a.optionList == nil {
		a.optionList = newOptionList(a.Options, &a.Entry)
	} else {
		a.optionList.UnselectAll()
		a.optionList.selected = 0 
	}
	holder := fyne.CurrentApp().Driver().CanvasForObject(a)

	if a.popup == nil {
		a.popup = widget.NewPopUp(a.optionList, holder)
	}
	a.popup.Resize(a.maxSize())
	a.popup.ShowAtPosition(a.popupPos())
	holder.Focus(a.optionList)
}

func (a* AutoComplete) Move(pos fyne.Position) {
    a.Entry.Move(pos)
    if a.popup != nil {
        a.popup.Resize(a.maxSize())
        a.popup.Move(a.popupPos())
    }	
}

func (a* AutoComplete) popupPos() fyne.Position {
    start := fyne.CurrentApp().Driver().AbsolutePositionForObject(a) 
    return start.Add(fyne.NewPos(0, a.Size().Height))
}

func (a* AutoComplete) maxSize() fyne.Size {
    canvas := fyne.CurrentApp().Driver().CanvasForObject(a)
    if canvas == nil {
        return fyne.NewSize(0, 0)
    }

    if a.itemHeight == 0 {
        a.itemHeight = a.optionList.CreateItem().MinSize().Height
    }

    optionsDisplayLength := len(a.Options)  
    if len(a.Options) > 5 {
        optionsDisplayLength = 5
    }
    listHeight := float32(optionsDisplayLength)*(a.itemHeight+2*theme.Padding()+theme.SeparatorThicknessSize()) + 2*theme.Padding()
    canvasSize := canvas.Size()
    entrySize := a.Size()

    if canvasSize.Height > listHeight {
        return fyne.NewSize(entrySize.Width, listHeight)
    }
    return fyne.NewSize(entrySize.Width, canvasSize.Height-a.Position().Y-entrySize.Height-theme.InputBorderSize()-theme.Padding())
}

type optionList struct {
    widget.List
    entry       *widget.Entry
    selected    int
    navigating  bool
    items       []string
}


func newOptionList(items []string, entry *widget.Entry) *optionList {
    op := &optionList {
        entry:       entry,
        selected:    0,
        items:       items,
        navigating:  false,
    }

    op.List = widget.List {
        Length: func() int {
            return len(op.items)
        },
        CreateItem: func() fyne.CanvasObject{ 
            return widget.NewLabel("")
        },
        UpdateItem: func(i widget.ListItemID, o fyne.CanvasObject) {
            o.(*widget.Label).SetText(op.items[i])
        },
        OnSelected: func(i widget.ListItemID) {
            if !op.navigating && i >= 0 {
                op.entry.SetText(op.items[i])
            }
            op.navigating = false
        },
    }
    op.ExtendBaseWidget(op)
    return op
}

func (op *optionList) SetOptions(items []string) {
    op.Unselect(op.selected)
    op.items = items
    op.Refresh()
    op.selected = 0
}

func (op *optionList) FocusGained() {
}

func (op *optionList) FocusLost() {
}

func (op *optionList) TypedKey(event *fyne.KeyEvent) {
    switch event.Name{
    case fyne.KeyDown:
        if op.selected < len(op.items)-1 {
            op.selected += 1
        } else {
            op.selected = 0
        }
        op.navigating = true
        op.Select(op.selected)
    case fyne.KeyUp:
        if op.selected < len(op.items)-1 {
            op.selected -= 1
        } else {
            op.selected = len(op.items)-1
        }
        op.navigating = true
        op.Select(op.selected)
    case fyne.KeyReturn, fyne.KeyEnter:
        if op.selected != -1 {
            op.entry.TypedKey(event)
        }
    default:
        op.entry.TypedKey(event)
    }
}

func (op *optionList) TypedRune(r rune) {
    op.entry.TypedRune(r)
}
