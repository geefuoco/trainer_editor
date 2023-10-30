package ui

import (
    "github.com/geefuoco/trainer_editor/data_objects"
    "github.com/geefuoco/trainer_editor/custom_widgets"
    "time"
    "strings"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/layout"
)


func createTrainerInfo(trainer *data_objects.Trainer) *fyne.Container{ 
    form := container.New(layout.NewFormLayout())
    content := container.NewVBox()

    trainerClassLabel := widget.NewLabel("Trainer Class")
    trainerClassSelectBox:= custom_widgets.NewAutoComplete(trainerClasses)
    trainerClassSelectBox.SetText(trainer.TrainerClass)
    trainerClassSelectBox.OnChanged = func(value string) {
        if len(value) < 4 || isProcessing {
            trainerClassSelectBox.HideCompletion()
            return
        }
        processingText = value
        processingTimer = time.AfterFunc(throttleInterval, func() {
            isProcessing = true
            defer func() {
                isProcessing = false
                processingText = ""
            }()

            filtered := filterOptions(trainerClasses, value)
            trainerClassSelectBox.SetOptions(filtered)
            trainerClassSelectBox.ShowCompletion()
            if SliceContains(trainerClasses, value) {
                trainer.TrainerClass = value
            }
        })
    }
    form.Add(trainerClassLabel)
    form.Add(trainerClassSelectBox)

    // Trainer Name
    trainerNameLabel := widget.NewLabel("Trainer Name")
    trainerNameEntry:= widget.NewEntry()
    trainerNameEntry.SetText(trainer.TrainerName)
    trainerNameEntry.OnChanged = func(s string) {
        if len(s) >= 3 {
            trainer.TrainerName = s
        }
    }
    form.Add(trainerNameLabel)
    form.Add(trainerNameEntry)

    // Encounter Music Gender
    encounterMusicGenderLabel := widget.NewLabel("Trainer Encounter Music")
    encounterMusicGenderSelectBox:= widget.NewSelect(trainerEncounterMusics, func(value string) {
        trainer.EncounterMusicGender = value
    })
    encounterMusicGenderSelectBox.SetSelected(trainer.EncounterMusicGender)
    form.Add(encounterMusicGenderLabel)
    form.Add(encounterMusicGenderSelectBox)

    // Trainer Pic
    trainerPicLabel := widget.NewLabel("Trainer Pic")
    trainerPicSelectBox := custom_widgets.NewAutoComplete(trainerPics)
    trainerPic = canvas.NewImageFromFile(trainerPicMap[trainer.TrainerPic])
    trainerPic.FillMode = canvas.ImageFillContain
    trainerPic.SetMinSize(fyne.NewSize(64, 64))
    trainerPicSelectBox.SetText(trainer.TrainerPic)
    trainerPicSelectBox.OnChanged = func(value string) {
        if len(value) < 3 || isProcessing {
            trainerPicSelectBox.HideCompletion()
            return
        }
        processingText = value
        processingTimer = time.AfterFunc(throttleInterval, func() {
            isProcessing = true
            defer func() {
                isProcessing = false
                processingText = ""
            }()

            filtered := filterOptions(trainerPics, value)
            trainerPicSelectBox.SetOptions(filtered)
            trainerPicSelectBox.ShowCompletion()
            _, has := trainerPicMap[value]
            if has {
                trainer.TrainerPic = value
                trainerPic = canvas.NewImageFromFile(trainerPicMap[trainer.TrainerPic])
                trainerPic.FillMode = canvas.ImageFillContain
                trainerPic.SetMinSize(fyne.NewSize(64, 64))
                picWrapper.Objects[1] = trainerPic
                picWrapper.Refresh()
            }
        })
    }
    trainerPicLabelWrapper := container.New(layout.NewFormLayout(), trainerPicLabel, trainerPicSelectBox)
    picWrapper = container.NewVBox(trainerPicLabelWrapper, trainerPic)
    content.Add(picWrapper)

    // Item 0
    itemLabel0 := widget.NewLabel("ITEM 0")
    itemSelectBox0 := custom_widgets.NewAutoComplete(items)
    form.Add(itemLabel0)
    itemValue := trainer.Items[0]
    itemSelectBox0.SetText(itemValue)
    itemSelectBox0.OnChanged = func(value string) {
        if len(value) < len("cut") || isProcessing {
            itemSelectBox0.HideCompletion()
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
            itemSelectBox0.SetOptions(filtered)
            itemSelectBox0.ShowCompletion()
            if SliceContains(items, value) {
                trainer.Items[0] = value;
            }
        })
    }
    form.Add(itemSelectBox0)

    // Item 1
    itemLabel1 := widget.NewLabel("ITEM 1")
    itemSelectBox1 := custom_widgets.NewAutoComplete(items)
    form.Add(itemLabel1)
    itemValue = trainer.Items[1]
    itemSelectBox1.SetText(itemValue)
    itemSelectBox1.OnChanged = func(value string) {
        if len(value) < len("cut") || isProcessing {
            itemSelectBox1.HideCompletion()
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
            itemSelectBox1.SetOptions(filtered)
            itemSelectBox1.ShowCompletion()
            if SliceContains(items, value) {
                trainer.Items[1] = value;
            }
        })
    }
    form.Add(itemSelectBox1)

    // Item 2

    itemLabel2 := widget.NewLabel("ITEM 2")
    itemSelectBox2 := custom_widgets.NewAutoComplete(items)
    form.Add(itemLabel2)
    itemValue = trainer.Items[2]
    itemSelectBox2.SetText(itemValue)
    itemSelectBox2.OnChanged = func(value string) {
        if len(value) < len("cut") || isProcessing {
            itemSelectBox2.HideCompletion()
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
            itemSelectBox2.SetOptions(filtered)
            itemSelectBox2.ShowCompletion()
            if SliceContains(items, value) {
                trainer.Items[2] = value;
            }
        })
    }
    form.Add(itemSelectBox2)
    // Item 3
    itemLabel3 := widget.NewLabel("ITEM 3")
    itemSelectBox3 := custom_widgets.NewAutoComplete(items)
    form.Add(itemLabel3)
    itemValue = trainer.Items[3]
    itemSelectBox3.SetText(itemValue)
    itemSelectBox3.OnChanged = func(value string) {
        if len(value) < len("cut") || isProcessing {
            itemSelectBox3.HideCompletion()
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
            itemSelectBox3.SetOptions(filtered)
            itemSelectBox3.ShowCompletion()
            if SliceContains(items, value) {
                trainer.Items[3] = value;
            }
        })
    }
    form.Add(itemSelectBox3)
    
    // AI Flags
    // Chunk the ai flags
    // Arbitrary 3 chunks
    var NUM_COLUMNS int = 3
    chunks := len(aiFlags) / NUM_COLUMNS
    if len(aiFlags) % NUM_COLUMNS != 0 {
        chunks++
    }
    aiFlagsLabel := widget.NewLabel("AI Flags")
    form.Add(aiFlagsLabel)
    aiFlagsCheckGroup := container.New(layout.NewGridLayout(NUM_COLUMNS))
    for j:=0; j < len(aiFlags); j+= chunks {
        end := j + chunks
        if end > len(aiFlags) {
            end = len(aiFlags)
        }
        check := widget.NewCheckGroup(aiFlags[j:end], func(opts []string) {
            trainer.AiFlags = opts
        })
        if trainer.AiFlags != nil {
            check.SetSelected(trainer.AiFlags)
        }
        aiFlagsCheckGroup.Add(check)
    }
    form.Add(container.NewHScroll(aiFlagsCheckGroup))
    // Double Battle
    doubleBattleLabel := widget.NewLabel("Double Battle")
    doubleBattleCheck := widget.NewCheck("", func(checked bool){
        trainer.DoubleBattle = checked
    })
    doubleBattleCheck.Checked = trainer.DoubleBattle
    form.Add(doubleBattleLabel)
    form.Add(doubleBattleCheck)

    content.Add(form)
    return content 
}


func filterOptions(options []string, text string) []string {
    filtered := []string{}
    for _, option := range options {
        if strings.Contains(option, text) {
            filtered = append(filtered, option)
        }
    }
    return filtered
}
