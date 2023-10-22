package ui

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
    "reflect"
    "strconv"
)

var trainers []*data_objects.Trainer
var trainerParties []*data_objects.TrainerParty
var list *widget.List
var content *fyne.Container
var grid *fyne.Container
var trainerInfo *fyne.Container
var partyInfo *fyne.Container
var items  []string
var aiFlags []string


func RunApp() {
	myApp := app.New()
	myWindow := myApp.NewWindow("PokemonEmerald Decomp Trainer Editor")
    myWindow.Resize(fyne.NewSize(900, 600))


    searchBar := NewCustomEntry(0)
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
                if strings.Contains(strings.ToLower(trainer.TrainerClass), strings.ToLower(str)) ||
                   strings.Contains(strings.ToLower(trainer.TrainerName), strings.ToLower(str)) {
                    filteredList = append(filteredList, trainer)
                }
            }
            list = createList(filteredList)
            listContainer.Objects = []fyne.CanvasObject{list}
            listContainer.Refresh()
        }

    }

    folderDialog := dialog.NewFolderOpen(
        func(uri fyne.ListableURI, err error) {
            if err != nil {
                fmt.Println("Error Occurred")
                return
            }
            if uri == nil {
                return
            }
            path := uri.Path()
            loadAllData(path)
            if trainers != nil {
                list = createList(trainers)
                listContainer.Objects = []fyne.CanvasObject{list}
                listContainer.Refresh()
                center.Hide()
                grid.Show()
                list.Select(0)
            }
        },
        myWindow,
    )

    toolbar := widget.NewToolbar(
        widget.NewToolbarAction(theme.FileIcon(), func() {
            folderDialog.Show()
        }),
    )

    trainerInfo = container.NewVBox()
    partyInfo = container.NewVBox()

    mainContent := container.NewAppTabs(
        container.NewTabItem("Trainer", trainerInfo),
        container.NewTabItem("Party", partyInfo),
    )

    vbox := container.NewVSplit(searchBar, listContainer)
    vbox.SetOffset(0)
    hbox := container.NewHSplit(vbox, mainContent)
    hbox.SetOffset(0.3)
    grid = container.NewMax(hbox)
    grid.Hide()
    content =  container.NewBorder(toolbar, nil, nil, nil, grid, center)

    myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func loadAllData(path string) {
    trainerPath := buildTrainerPath(path)
    trainerPartyPath := buildTrainerPartiesPath(path)
    itemPath := buildItemPath(path)
    aiFlagsPath := buildAiFlagsPath(path)

    trainers = parsers.ParseTrainers(trainerPath)
    trainerParties = parsers.ParseTrainerParties(trainerPartyPath)
    items = parsers.ParseItems(itemPath)
    aiFlags = parsers.ParseAiFlags(aiFlagsPath)
}

func getTrainerParty(partyName string) *data_objects.TrainerParty {
    for _, party := range(trainerParties) {
        if party.Trainer == partyName {
            return party
        }
    }
    return nil
}

func createList(listOfTrainers []*data_objects.Trainer) *widget.List {
    list := widget.NewList(
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

    list.OnSelected = func(id widget.ListItemID) {
        // Cache the trainer that is currently being viewed
        // if trainerInfo != nil {
            // Go through each of the Canvas Objects
        // }
        selectedTrainer := listOfTrainers[id]
        // ie sParty_Sawyer1
        selectedParty := getTrainerParty(selectedTrainer.GetPartyName())
        if selectedParty != nil {
            trainerInfo.Objects = []fyne.CanvasObject{}
            updatedTrainerInfo := createTrainerInfo(selectedTrainer)
            trainerInfo.Add(updatedTrainerInfo)
            trainerInfo.Refresh()
        } 
    }
    
    return list
}

func createTrainerInfo(trainer *data_objects.Trainer) *fyne.Container{ 
    content := container.New(layout.NewFormLayout())

    structValue := reflect.ValueOf(trainer).Elem()

    for i :=0; i < structValue.NumField(); i++ {
        field := structValue.Field(i)
        fieldName := structValue.Type().Field(i).Name
        if fieldName == "Party" {
            continue
        }

        var value string
        switch field.Interface().(type){ 
        case []string:
            // Chunk the ai flags
            // Arbitrary 3 chunks
            var NUM_COLUMNS int = 3
            chunks := len(aiFlags) / NUM_COLUMNS
            if len(aiFlags) % NUM_COLUMNS != 0 {
                chunks++
            }
            label := widget.NewLabel(fieldName)
            content.Add(label)
            checkGroupHolder := container.New(layout.NewGridLayout(NUM_COLUMNS))
            for j:=0; j < len(aiFlags); j+= chunks {
                end := j + chunks
                if end > len(aiFlags) {
                    end = len(aiFlags)
                }
                check := widget.NewCheckGroup(aiFlags[j:end], func(opts []string) {
                    for _, opt := range(opts) {
                        if !sliceContains(trainer.AiFlags, opt) {
                            trainer.AiFlags = append(trainer.AiFlags, opt)
                        }

                    }
                })
                if trainer.AiFlags != nil {
                    check.SetSelected(trainer.AiFlags)
                }
                checkGroupHolder.Add(check)
            }
            content.Add(container.NewHScroll(checkGroupHolder))
            continue
        case [4]string:
            for j:=0; j < 4; j++ {
                label := widget.NewLabel(fieldName + " " + strconv.Itoa(j))
                selectBox := NewCompletionEntry(items)
                content.Add(label)
                itemValue := trainer.Items[j]
                selectBox.SetText(itemValue)
                selectBox.Id = j
                // Order matters here. Having this callback set earlier will cause 
                // A segfault, on behalf of the fyne/x library
                selectBox.OnChanged = func(s string) {
                    if len(s) == 0 {
                        selectBox.HideCompletion()
                        return
                    }
                    // Filter the items array
                    text := selectBox.Text
                    var filteredItems []string
                    for _, item := range(items) {
                        if strings.Contains(item, text) {
                            filteredItems = append(filteredItems, item)
                        }
                    }
                    if len(s) >= 9 {
                        if sliceContains(items, s) {
                            trainer.Items[selectBox.Id] = s;
                        }
                    }
                    selectBox.SetOptions(filteredItems)
                    selectBox.ShowCompletion()
                }
                content.Add(selectBox)
            }
            continue
        case string:
            value = field.Interface().(string)
        case bool:
            label := widget.NewLabel(fieldName)
            check := widget.NewCheck("", func(checked bool){
                fmt.Printf("%t\n", checked)
                trainer.DoubleBattle = checked
            })
            checked := field.Interface().(bool)
            check.Checked = checked
            content.Add(label)
            content.Add(check)
            continue
        default:
            value = "Unsupported Type"
        }


        label := widget.NewLabel(fieldName)
        entry := NewCustomEntry(i)
        entry.SetText(value)
        entry.OnChanged = func(s string) {
            if len(s) != 0 {
                structValue.Field(entry.Id).SetString(s)
            }
            fmt.Println(trainer)
        }

        content.Add(label)
        content.Add(entry)
    }

    return content
}

func sliceContains(slice []string, item string) bool {
    for _, i:= range(slice) {
        if i == item {
            return true
        }
    }
    return false
}

func createPartyInfo(trainerParty *data_objects.TrainerParty) *fyne.Container{
    content := container.New(layout.NewGridLayout(2))
    return content
}

func buildTrainerPartiesPath(path string) string {
    buf := strings.Builder{}
    buf.WriteString(path)
    buf.WriteString("/src")
    buf.WriteString("/data")
    buf.WriteString("/trainer_parties.h")
    return buf.String()
}

func buildTrainerPath(path string) string {
    buf := strings.Builder{}
    buf.WriteString(path)
    buf.WriteString("/src")
    buf.WriteString("/data")
    buf.WriteString("/trainers.h")
    return buf.String()
}

func buildItemPath(path string) string {
    buf := strings.Builder{}
    buf.WriteString(path)
    buf.WriteString("/include")
    buf.WriteString("/constants")
    buf.WriteString("/items.h")
    return buf.String()
}

func buildAiFlagsPath(path string) string {
    buf := strings.Builder{}
    buf.WriteString(path)
    buf.WriteString("/include")
    buf.WriteString("/constants")
    buf.WriteString("/battle_ai.h")
    return buf.String()
}

