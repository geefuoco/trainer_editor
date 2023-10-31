package ui

import ( 
    "image/color"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/container"
    "github.com/geefuoco/trainer_editor/data_objects"
    "github.com/geefuoco/trainer_editor/parsers"
    "github.com/geefuoco/trainer_editor/logging"
    "github.com/geefuoco/trainer_editor/custom_widgets"
    "strings"
    "path/filepath"
    "time"
)

// For Debouncing autocomplete widgets
var isProcessing            bool
var processingText          string
var processingTimer         *time.Timer
var throttleInterval = 400*time.Millisecond
var isSaving                bool
//  Globals for storing data from parsers
var projectPath             string
var trainers                []*data_objects.Trainer
var trainerParties          []*data_objects.TrainerParty
var list                    *widget.List
var content                 *fyne.Container
var grid                    *fyne.Container
var picWrapper              *fyne.Container
var trainerInfo             *fyne.Container
var partyInfo               *fyne.Container
var items                   []string
var aiFlags                 []string
var trainerPics             []string
var trainerClasses          []string
var trainerEncounterMusics  []string
var trainerPic              *canvas.Image
var pokemonPicWrapper       *fyne.Container
var pokemonPic              *canvas.Image
var moves                   []string
var species                 []string
var abilities               []string
var selectedMonIndex        widget.ListItemID
var trainerPicMap = make(map[string]string)
var pokemonPicMap = make(map[string]string)

const HEIGHT = 900
const WIDTH = 1200


func RunApp() {
	myApp := app.New()
	myWindow := myApp.NewWindow("PokemonEmerald Decomp Trainer Editor")
    myWindow.Resize(fyne.NewSize(WIDTH, HEIGHT))


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
                logging.ErrorLog(err.Error())
                return
            }
            if uri == nil {
                return
            }
            path := uri.Path()
            projectPath = path
            loadAllData(path)
            if trainers != nil {
                list = createList(trainers)
                listContainer.Objects = []fyne.CanvasObject{list}
                listContainer.Refresh()
                center.Hide()
                grid.Show()
                list.Select(0)
            }
            if trainers == nil || len(trainers) == 0 {
                createModal(myWindow, "Error", "Not in pokeemerald directory").Show()
                logging.ErrorLog("Error: Could not populate the trainer list")
            }
        },
        myWindow,
    )

    toolbar := widget.NewToolbar(
        widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
            folderDialog.Show()
        }),
        widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
            // TODO
            // Might need a better way to validate that trainers are being
            // saved correctly
            if isSaving {
                return
            }
            if trainers != nil && len(trainers) > 0 && trainerParties != nil && len(trainerParties) > 0{
                if projectPath == "" {
                    logging.ErrorLog("Error: Did not find path to pokeemerald.")
                    createModal(myWindow, "Error", "Did not find path to pokeemerald directory").Show()
                    return
                }
                trainerPath := buildTrainerPath(projectPath)
                trainerPartyPath := buildTrainerPartiesPath(projectPath)

                isSaving = true
                err := data_objects.SaveAll(trainerPath, trainerPartyPath, trainers, trainerParties)
                if err != nil {
                    logging.ErrorLog(err.Error())
                    // Error Popup
                    createModal(myWindow, "Error", err.Error()).Show()
                } else {
                    isSaving = false
                    createModal(myWindow, "Saved", "Save successful").Show()
                }
            }
        }),
        widget.NewToolbarSpacer(),
        custom_widgets.NewToolbarLabel("Toggle Logging"),
        widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
            logging.EnableLogging()
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

func createModal(myWindow fyne.Window, title string, msg string) *widget.PopUp {
    var popup *widget.PopUp
    label := widget.NewLabel(title)
    message := widget.NewLabel(msg)
    closeButton := widget.NewButton("Close", func() {
        popup.Hide()
    })
    content := container.NewVBox(label, message, closeButton)
    popup = widget.NewModalPopUp(content, myWindow.Canvas())
    return popup
}

func loadAllData(path string) {
    trainerPath := buildTrainerPath(path)
    trainerPartyPath := buildTrainerPartiesPath(path)
    itemPath := buildItemPath(path)
    aiFlagsPath := buildAiFlagsPath(path)
    encounterMusicPath := buildTrainerEncounterMusicPath(path)
    movesPath := buildMovesPath(path)
    abilitiesPath := buildAbilitiesPath(path)
    
    pokemonSpeciesPath := buildPokemonSpeciesPath(path)
    pokemonSpritePath := buildPokemonSpritePath(path)

    pokemonPicMap = parsers.ParsePokemonPics(pokemonSpeciesPath, pokemonSpritePath)
    for k, v := range(pokemonPicMap) {
        pokemonPicMap[k] = filepath.Join(path, v)
        species = append(species, k)
    }

    trainergFrontPicPath := buildgTrainerFrontPicPath(path)
    trainerTrainerSpritePath := buildTrainerSpritePath(path)

    trainerPicMap = parsers.ParseTrainerPics(trainergFrontPicPath, trainerTrainerSpritePath)
    for k, v := range(trainerPicMap) {
        trainerPics = append(trainerPics, k)
        trainerPicMap[k] = filepath.Join(path, v)
    }
    trainers = parsers.ParseTrainers(trainerPath)
    for _, t := range(trainers) {
        if !SliceContains(trainerClasses, t.TrainerClass) {
            trainerClasses = append(trainerClasses, t.TrainerClass)
        }
    }
    trainerParties = parsers.ParseTrainerParties(trainerPartyPath)
    items = parsers.ParseItems(itemPath)
    aiFlags = parsers.ParseAiFlags(aiFlagsPath)
    trainerEncounterMusics = parsers.ParseTrainerEncounterMusic(encounterMusicPath)
    moves = parsers.ParseMoves(movesPath)
    abilities = parsers.ParsePokemonAbilities(abilitiesPath)
    logging.InfoLog("Loaded %d trainers", len(trainers))
    logging.InfoLog("Loaded %d trainer parties", len(trainerParties))
}

func getTrainerParty(partyName string) *data_objects.TrainerParty {
    if partyName == "NULL"{
        return nil
    }
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
            displayText := listOfTrainers[i].TrainerClass + " " + listOfTrainers[i].TrainerKey
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


func SliceContains(slice []string, item string) bool {
    for _, i:= range(slice) {
        if i == item {
            return true
        }
    }
    return false
}

func buildPokemonSpeciesPath(path string) string {
    return filepath.Join(path, "src", "data", "pokemon_graphics", "front_pic_table.h")
}

func buildPokemonSpritePath(path string) string {
    return filepath.Join(path, "src", "data", "graphics", "pokemon.h")
}

func buildMovesPath(path string) string {
    return filepath.Join(path, "include", "constants", "moves.h")
}

func buildAbilitiesPath(path string) string {
    return filepath.Join(path, "include", "constants", "abilities.h")
}

func buildTrainerPartiesPath(path string) string {
    return filepath.Join(path, "src", "data", "trainer_parties.h")
}

func buildTrainerPath(path string) string {
    return filepath.Join(path, "src", "data", "trainers.h")
}

func buildItemPath(path string) string {
    return filepath.Join(path, "include", "constants", "items.h")
}

func buildAiFlagsPath(path string) string {
    return filepath.Join(path, "include", "constants", "battle_ai.h")
}

func buildgTrainerFrontPicPath(path string) string {
    return filepath.Join(path, "src", "data", "graphics", "trainers.h")
}

func buildTrainerSpritePath(path string) string {
    return filepath.Join(path, "src", "data", "trainer_graphics", "front_pic_tables.h")
}

func buildTrainerEncounterMusicPath(path string) string {
    return filepath.Join(path, "include", "constants", "trainers.h")
}
