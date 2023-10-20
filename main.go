package main

import (
    "image/color"
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/container"
    "github.com/geefuoco/trainer_editor/data_objects"
    "github.com/geefuoco/trainer_editor/parsers"
    "strings"
)

// Filepath
var trainers []*data_objects.Trainer
var list *widget.List
var content *fyne.Container

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

func buildTrainerPath(path string) string {
    buf := strings.Builder{}
    buf.WriteString(path)
    buf.WriteString("/src")
    buf.WriteString("/data")
    buf.WriteString("/trainers.h")
    return buf.String()
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("PokemonEmerald Decomp Trainer Editor")
    myWindow.Resize(fyne.NewSize(900, 600))


    searchBar := widget.NewEntry()
    searchBar.SetPlaceHolder("Search")

    textWidget := canvas.NewText("Open Pokemon Directory to begin", color.White)
    textWidget.TextSize = 30
    center := container.NewCenter(textWidget)

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
            list.Refresh()
        }

    }

    folderDialog := dialog.NewFolderOpen(
        func(uri fyne.ListableURI, err error) {
            if err != nil {
                fmt.Println("Error Occurred")
                return
            }
            path := uri.Path()
            fmt.Println(path)

            trainerPath := buildTrainerPath(path)
            fmt.Println(trainerPath)
            trainers = parsers.ParseTrainers(trainerPath)
            list = createList(trainers)
            listContainer.Objects = []fyne.CanvasObject{list}
            listContainer.Refresh()
            center.Hide()
        },
        myWindow,
    )

    toolbar := widget.NewToolbar(
        widget.NewToolbarAction(theme.FileIcon(), func() {
            folderDialog.Show()
        }),
    )

    vbox := container.NewVSplit(searchBar, listContainer)
    vbox.SetOffset(0)
    grid := container.New(layout.NewGridLayout(3), vbox, layout.NewSpacer(),layout.NewSpacer())
    content =  container.NewBorder(toolbar, nil, nil, nil, grid, center)

    myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
