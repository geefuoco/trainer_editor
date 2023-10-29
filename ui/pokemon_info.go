package ui

import (
    "github.com/geefuoco/trainer_editor/data_objects"
    "github.com/geefuoco/trainer_editor/custom_widgets"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/canvas"
    "strconv"
    "time"
    "fmt"
)

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
        movesForm = createMovesForm(mon)
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


    currentMon := trainerMonList[selectedMonIndex]
    movesForm = createMovesForm(currentMon)
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
    speciesSelectBox := custom_widgets.NewAutoComplete(species)
    speciesSelectBox.SetText(mon.Species)
    speciesSelectBox.OnChanged = func(value string) {
        if len(value) < 4 || isProcessing {
            speciesSelectBox.HideCompletion()
            return
        }
        processingText = value
        processingTimer = time.AfterFunc(throttleInterval, func() {
            isProcessing = true
            defer func() {
                isProcessing = false
                processingText = ""
            }()

            filtered := filterOptions(species, value)
            speciesSelectBox.SetOptions(filtered)
            speciesSelectBox.ShowCompletion()
            // Utilize the map that has species as keys
            pic, has := pokemonPicMap[value]
            if has {
                mon.Species = value;
                pokemonPic = canvas.NewImageFromFile(pic)
                pokemonPic.FillMode = canvas.ImageFillContain
                pokemonPic.SetMinSize(fyne.NewSize(64, 64))
                pokemonPic.Refresh()
                pokemonPicWrapper.Objects[0] = pokemonPic
                pokemonPicWrapper.Refresh()
            }
        })
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
    heldItemSelectBox := custom_widgets.NewAutoComplete(items)
    heldItemSelectBox.SetText(mon.HeldItem)
    heldItemSelectBox.OnChanged = func(value string) {
        if len(value) < 4 || isProcessing {
            heldItemSelectBox.HideCompletion()
            return
        }
        processingText = value
        processingTimer = time.AfterFunc(throttleInterval, func() {
            isProcessing = true
            defer func() {
                isProcessing = false
                processingText = ""
            }()

            filtered := filterOptions(items, value)
            heldItemSelectBox.SetOptions(filtered)
            heldItemSelectBox.ShowCompletion()
            if SliceContains(items, value) {
                mon.HeldItem = value;
            }
        })
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
    abilityLabel := widget.NewLabel("Ability")
    abilitySelectBox := custom_widgets.NewAutoComplete(abilities)
    abilitySelectBox.SetText(mon.Ability)
    abilitySelectBox.OnChanged = func(value string) {
        if len(value) < 4 || isProcessing {
            abilitySelectBox.HideCompletion()
            return
        }
        processingText = value
        processingTimer = time.AfterFunc(throttleInterval, func() {
            isProcessing = true
            defer func() {
                isProcessing = false
                processingText = ""
            }()

            filtered := filterOptions(abilities, value)
            abilitySelectBox.SetOptions(filtered)
            abilitySelectBox.ShowCompletion()
            if SliceContains(abilities, value) {
                mon.Ability= value;
            }
        })
    }
    form.Add(abilityLabel)
    form.Add(abilitySelectBox)
    return form
}

func createMovesForm(mon *data_objects.TrainerMon) *fyne.Container {
    form := container.New(layout.NewFormLayout())
    // Item 0
    moveLabel0 := widget.NewLabel("MOVE 0")
    moveSelectBox0 := custom_widgets.NewAutoComplete(moves)
    form.Add(moveLabel0)
    moveValue := mon.Moves[0]
    moveSelectBox0.SetText(moveValue)
    moveSelectBox0.OnChanged = func(value string) {
        if len(value) < len("cut") || isProcessing {
            moveSelectBox0.HideCompletion()
            return
        }
        processingText = value
        processingTimer = time.AfterFunc(throttleInterval, func() {
            isProcessing = true
            defer func() {
                isProcessing = false
                processingText = ""
            }()

            filtered := filterOptions(moves, value)
            moveSelectBox0.SetOptions(filtered)
            moveSelectBox0.ShowCompletion()
            if SliceContains(moves, value) {
                mon.Moves[0] = value;
            }
        })
    }
    form.Add(moveSelectBox0)

    // Item 1
    moveLabel1 := widget.NewLabel("MOVE 1")
    moveSelectBox1 := custom_widgets.NewAutoComplete(moves)
    form.Add(moveLabel1)
    moveValue = mon.Moves[1]
    moveSelectBox1.SetText(moveValue)
    moveSelectBox1.OnChanged = func(value string) {
        if len(value) < len("cut") || isProcessing {
            moveSelectBox1.HideCompletion()
            return
        }
        processingText = value
        processingTimer = time.AfterFunc(throttleInterval, func() {
            isProcessing = true
            defer func() {
                isProcessing = false
                processingText = ""
            }()

            filtered := filterOptions(moves, value)
            moveSelectBox1.SetOptions(filtered)
            moveSelectBox1.ShowCompletion()
            if SliceContains(moves, value) {
                mon.Moves[1] = value;
            }
        })
    }
    form.Add(moveSelectBox1)

    // Item 2

    moveLabel2 := widget.NewLabel("MOVE 2")
    moveSelectBox2 := custom_widgets.NewAutoComplete(moves)
    form.Add(moveLabel2)
    moveValue = mon.Moves[2]
    moveSelectBox2.SetText(moveValue)
    moveSelectBox2.OnChanged = func(value string) {
        if len(value) < len("cut") || isProcessing {
            moveSelectBox2.HideCompletion()
            return
        }
        processingText = value
        processingTimer = time.AfterFunc(throttleInterval, func() {
            isProcessing = true
            defer func() {
                isProcessing = false
                processingText = ""
            }()

            filtered := filterOptions(moves, value)
            moveSelectBox2.SetOptions(filtered)
            moveSelectBox2.ShowCompletion()
            if SliceContains(moves, value) {
                mon.Moves[2] = value;
            }
        })
    }
    form.Add(moveSelectBox2)
    // Item 3
    moveLabel3 := widget.NewLabel("MOVE 3")
    moveSelectBox3 := custom_widgets.NewAutoComplete(moves)
    form.Add(moveLabel3)
    moveValue = mon.Moves[3]
    moveSelectBox3.SetText(moveValue)
    moveSelectBox3.OnChanged = func(value string) {
        if len(value) < len("cut") || isProcessing {
            moveSelectBox3.HideCompletion()
            return
        }
        processingText = value
        processingTimer = time.AfterFunc(throttleInterval, func() {
            isProcessing = true
            defer func() {
                isProcessing = false
                processingText = ""
            }()

            filtered := filterOptions(moves, value)
            moveSelectBox3.SetOptions(filtered)
            moveSelectBox3.ShowCompletion()
            if SliceContains(moves, value) {
                mon.Moves[3] = value;
            }
        })
    }
    form.Add(moveSelectBox3)
    return form
}

