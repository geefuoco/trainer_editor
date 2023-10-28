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
    "strconv"
    "strings"
)

var trainers []*data_objects.Trainer
var trainerParties []*data_objects.TrainerParty
var list *widget.List
var content *fyne.Container
var grid *fyne.Container
var picWrapper *fyne.Container
var trainerInfo *fyne.Container
var partyInfo *fyne.Container
var items  []string
var aiFlags []string
var trainerPicMap = make(map[string]string)
var trainerPics []string
var trainerClasses []string
var trainerEncounterMusics []string
var trainerPic *canvas.Image

var pokemonPicMap = make(map[string]string)
var pokemonPicWrapper *fyne.Container
var pokemonPic *canvas.Image
var moves []string
var species []string
var selectedMonIndex widget.ListItemID

const HEIGHT = 900
const WIDTH = 1200

func RunApp() {
	myApp := app.New()
	myWindow := myApp.NewWindow("PokemonEmerald Decomp Trainer Editor")
    myWindow.Resize(fyne.NewSize(WIDTH, HEIGHT))


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

    trainerInfo = container.NewMax()
    partyInfo = container.NewMax()

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
    encounterMusicPath := buildTrainerEncounterMusicPath(path)
    movesPath := buildMovesPath(path)
    
    pokemonSpeciesPath := buildPokemonSpeciesPath(path)
    pokemonSpritePath := buildPokemonSpritePath(path)

    pokemonPicMap = parsers.ParsePokemonPics(pokemonSpeciesPath, pokemonSpritePath)
    for k, v := range(pokemonPicMap) {
        pokemonPicMap[k] = path + "/" +  v
        species = append(species, k)
    }

    trainergFrontPicPath := buildgTrainerFrontPicPath(path)
    trainerTrainerSpritePath := buildTrainerSpritePath(path)

    trainerPicMap = parsers.ParseTrainerPics(trainergFrontPicPath, trainerTrainerSpritePath)
    for k, v := range(trainerPicMap) {
        trainerPics = append(trainerPics, k)
        trainerPicMap[k] = path + "/" +  v
    }
    trainers = parsers.ParseTrainers(trainerPath)
    for _, t := range(trainers) {
        if !sliceContains(trainerClasses, t.TrainerClass) {
            trainerClasses = append(trainerClasses, t.TrainerClass)
        }
    }
    trainerParties = parsers.ParseTrainerParties(trainerPartyPath)
    items = parsers.ParseItems(itemPath)
    aiFlags = parsers.ParseAiFlags(aiFlagsPath)
    trainerEncounterMusics = parsers.ParseTrainerEncounterMusic(encounterMusicPath)
    moves = parsers.ParseMoves(movesPath)
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
        selectedTrainer := listOfTrainers[id]
        selectedParty := getTrainerParty(selectedTrainer.GetPartyName())
        selectedMonIndex = 0
        if selectedParty != nil {
            trainerInfo.Objects = []fyne.CanvasObject{}
            updatedTrainerInfo := createTrainerInfo(selectedTrainer)
            trainerInfo.Add(updatedTrainerInfo)
            trainerInfo.Refresh()

            partyInfo.Objects = []fyne.CanvasObject{}
            updatedPartyInfo := createPartyInfo(selectedParty)
            partyInfo.Add(updatedPartyInfo)
            partyInfo.Refresh()
        } 
    }
    
    return list
}


func sliceContains(slice []string, item string) bool {
    for _, i:= range(slice) {
        if i == item {
            return true
        }
    }
    return false
}

const (
    PARTY_LIST = iota
    SPRITE
    STATS
    MOVES
    MAIN
    IVS
    EVS
)

func createPartyInfo(trainerParty *data_objects.TrainerParty) *fyne.Container{
    content := container.NewMax()
    var movesForm *fyne.Container
    var infoForm *fyne.Container
    var grid *fyne.Container

    trainerMonList := trainerParty.Party

    monList := widget.NewList(func() int {
        return len(trainerMonList)
    },
    func() fyne.CanvasObject {
        return widget.NewLabel("template")
    },
    func(i widget.ListItemID, o fyne.CanvasObject) {
        mon := trainerMonList[i]
        label := o.(*widget.Label)
        label.SetText(mon.Species)
    })

    monList.OnSelected = func(i widget.ListItemID) {
        selectedMonIndex = i
        mon := trainerMonList[i]
        value, has := pokemonPicMap[mon.Species]
        if !has {
            fmt.Println("Error: '" + mon.Species + "' was not in the Pokemon Pic Map")
            pokemonPic = canvas.NewImageFromFile(pokemonPicMap["SPECIES_NONE"])
        } else {
            pokemonPic = canvas.NewImageFromFile(value)
        }
        pokemonPic.FillMode = canvas.ImageFillContain
        pokemonPic.SetMinSize(fyne.NewSize(64, 64))
        pokemonPic.Refresh()
        pokemonPicWrapper.Objects[0] = pokemonPic
        pokemonPicWrapper.Refresh()
        movesForm = createMovesForm(trainerMonList)
        infoForm = createPokemonInfoForm(trainerMonList)
        // Should always be in the bottom left (3rd) position
        grid.Objects[2] = movesForm
        grid.Objects[3] = infoForm
    }


    add := widget.NewButton("Append", func() {
        if (len(trainerMonList) + 1) <= 6 { 
            val := data_objects.TemplateMon()
            trainerMonList = append(trainerMonList, val)
            trainerParty.Party = trainerMonList
            monList.Refresh()
        }
    })

    remove := widget.NewButton("Remove", func() {
        if (len(trainerMonList)-1) >= 1 {
            trainerMonList = append(trainerMonList[:selectedMonIndex], trainerMonList[selectedMonIndex+1:]...)
            trainerParty.Party = trainerMonList
            var selectIdx int
            if selectedMonIndex == 0 {
                selectIdx = 0
            } else {
                selectIdx = selectedMonIndex-1
            }
            monList.Refresh()
            monList.OnSelected(selectIdx)
        }
    })

    buttons := container.New(layout.NewGridLayout(2), add, remove)
    monListContainer := container.NewMax(monList)
    leftPanel := container.NewVSplit(buttons, monListContainer)
    leftPanel.SetOffset(0.1)


    movesForm = createMovesForm(trainerMonList)
    infoForm = createPokemonInfoForm(trainerMonList)
    pokemonPicWrapper = container.NewMax(pokemonPic)
    grid = container.New(layout.NewGridLayout(2), leftPanel, pokemonPicWrapper, movesForm, infoForm)

    // Make sure only to select once all other containers have been defined
    monList.Select(0)
    content.Add(grid)
    return content
}

func createPokemonInfoForm(trainerMonList []*data_objects.TrainerMon) *fyne.Container {
    form := container.New(layout.NewFormLayout())
    // Species 
    mon := trainerMonList[selectedMonIndex]
    label := widget.NewLabel("Species")
    speciesSelectBox := NewCompletionEntry(species)
    speciesSelectBox.SetText(mon.Species)
    speciesSelectBox.OnChanged = func(s string) {
        if len(s) == 0 {
            speciesSelectBox.HideCompletion()
            return
        }
        var filteredItems []string
        for _, item := range(species) {
            if strings.Contains(item, s) {
                filteredItems = append(filteredItems, item)
            }
        }
        if len(s) >= 5 {
            speciesSelectBox.SetOptions(filteredItems)
            speciesSelectBox.ShowCompletion()
        }
        if len(s) >= 9 {
            // Utilize the map that has species as keys
            value, has := pokemonPicMap[s]
            if has {
                mon.Species = s;
                pokemonPic = canvas.NewImageFromFile(value)
                pokemonPic.FillMode = canvas.ImageFillContain
                pokemonPic.SetMinSize(fyne.NewSize(64, 64))
                pokemonPic.Refresh()
                pokemonPicWrapper.Objects[0] = pokemonPic
                pokemonPicWrapper.Refresh()
            }
        }
    }
    form.Add(label)
    form.Add(speciesSelectBox)
    // Level
    label = widget.NewLabel("Level")
    levelEntry := widget.NewEntry()
    levelEntry.SetText(strconv.FormatUint(mon.Lvl, 10))
    levelEntry.OnChanged = func(s string) {
        lvl, err := strconv.ParseUint(s, 10, 64)
        if err != nil {
            fmt.Printf("Error: could not parse '%s' to int\n", s)
            return
        }
        if lvl > 0 && lvl <= 100 {
            mon.Lvl = lvl
        }
    }
    form.Add(label)
    form.Add(levelEntry)
    // HeldItem
    label = widget.NewLabel("Held Item")
    heldItemSelectBox := NewCompletionEntry(items)
    heldItemSelectBox.SetText(mon.HeldItem)
    heldItemSelectBox.OnChanged = func(s string) {
        if len(s) == 0 {
            heldItemSelectBox.HideCompletion()
            return
        }
        var filteredItems []string
        for _, item := range(items) {
            if strings.Contains(item, s) {
                filteredItems = append(filteredItems, item)
            }
        }
        if len(s) >= 5 {
            heldItemSelectBox.SetOptions(filteredItems)
            heldItemSelectBox.ShowCompletion()
        }
        if len(s) >= 9 {
            if sliceContains(items, s) {
                mon.HeldItem = s;
            }
        }
    }
    form.Add(label)
    form.Add(heldItemSelectBox)
    // Shiny
    label = widget.NewLabel("Shiny")
    check := widget.NewCheck("", func(val bool) {
        mon.IsShiny = val
    })
    check.Checked = mon.IsShiny
    form.Add(label)
    form.Add(check)

    // Abilitiy
    // label = widget.NewLabel("Ability")
    // selectBox = NewCompletionEntry(abilities)
    // form.Add(label)
    // form.Add(selectBox)
    // selectBox.SetText(mon.Ability)
    return form
}

func createMovesForm(trainerMonList []*data_objects.TrainerMon) *fyne.Container {
    form := container.New(layout.NewFormLayout())
    for j:=0; j < 4; j++ {
        label := widget.NewLabel("MOVE " + strconv.Itoa(j))
        selectBox := NewCompletionEntry(moves)
        form.Add(label)
        itemValue := trainerMonList[selectedMonIndex].Moves[j]
        selectBox.SetText(itemValue)
        selectBox.Id = j
        // Order matters here. Having this callback set earlier will cause 
        // A segfault, on behalf of the fyne/x library
        selectBox.OnChanged = func(s string) {
            if len(s) == 0 {
                selectBox.HideCompletion()
                return
            }
            var filteredItems []string
            for _, item := range(moves) {
                if strings.Contains(item, s) {
                    filteredItems = append(filteredItems, item)
                }
            }
            if len(s) >= 5 {
                selectBox.SetOptions(filteredItems)
                selectBox.ShowCompletion()
            }
            if len(s) >= len("MOVE_CUT") {
                if sliceContains(moves, s) {
                    trainerMonList[selectedMonIndex].Moves[selectBox.Id] = s;
                }
            }
        }
        form.Add(selectBox)
    }
    return form
}

func buildPokemonSpeciesPath(path string) string {
    buf := strings.Builder{}
    buf.WriteString(path)
    buf.WriteString("/src")
    buf.WriteString("/data")
    buf.WriteString("/pokemon_graphics")
    buf.WriteString("/front_pic_table.h")
    return buf.String()
}

func buildPokemonSpritePath(path string) string {
    buf := strings.Builder{}
    buf.WriteString(path)
    buf.WriteString("/src")
    buf.WriteString("/data")
    buf.WriteString("/graphics")
    buf.WriteString("/pokemon.h")
    return buf.String()
}

func buildMovesPath(path string) string {
    buf := strings.Builder{}
    buf.WriteString(path)
    buf.WriteString("/include")
    buf.WriteString("/constants")
    buf.WriteString("/moves.h")
    return buf.String()
}
