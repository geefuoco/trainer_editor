package main

import (
    // "image/color"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    // "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/container"
    "github.com/geefuoco/trainer_editor/data_objects"
    "github.com/geefuoco/trainer_editor/parsers"
    "strings"
)

// Filepath
var trainers = parsers.ParseTrainers("/home/gianl/projects/pokemon_decomps/pokeemerald-expansion/src/data/trainers.h")
var list *widget.List

func createList(listOfTrainers []*data_objects.Trainer) *widget.List {
    return widget.NewList(
        func() int {
            return len(listOfTrainers)
        },
        func() fyne.CanvasObject {
            return widget.NewLabel("template")
        },
        func(i widget.ListItemID, o fyne.CanvasObject) {
            displayText := listOfTrainers[i].TrainerClass + " " + listOfTrainers[i].TrainerName
            o.(*widget.Label).SetText(displayText)
        })
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("PokemonEmerald Decomp Trainer Editor")
    myWindow.Resize(fyne.NewSize(900, 600))


    searchBar := widget.NewEntry()
    searchBar.SetPlaceHolder("Search")

    list = createList(trainers)
    listContainer := container.NewMax(list)

    searchBar.OnChanged = func(str string) {
        var filteredList []*data_objects.Trainer
        if str == "" {
            list.Refresh() 
        } else {
            for _, trainer := range trainers {
                if strings.Contains(strings.ToLower(trainer.TrainerClass), str) || strings.Contains(strings.ToLower(trainer.TrainerName), str) {
                    filteredList = append(filteredList, trainer)
                }
            }
            list = createList(filteredList)
            listContainer.Objects = []fyne.CanvasObject{list}
            listContainer.Refresh()
        }

    }

    vbox := container.NewVSplit(searchBar, listContainer)
    vbox.SetOffset(0)

    // partyPanel := container.NewMax()
    content := container.New(layout.NewGridLayout(3), vbox, layout.NewSpacer())


	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
