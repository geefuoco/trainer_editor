package ui

import (
    "github.com/geefuoco/trainer_editor/data_objects"
    "fmt"
    "reflect"
    "strconv"
    "strings"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/canvas"

)

func createTrainerInfo(trainer *data_objects.Trainer) *fyne.Container{ 
    form := container.New(layout.NewFormLayout())
    content := container.NewVBox()

    structValue := reflect.ValueOf(trainer).Elem()

    for i :=0; i < structValue.NumField(); i++ {
        field := structValue.Field(i)
        fieldName := structValue.Type().Field(i).Name
        if fieldName == "Party" {
            continue
        } else if fieldName == "TrainerClass" {
            label := widget.NewLabel(fieldName)
            entry:= widget.NewSelect(trainerClasses, func(value string) {
                trainer.TrainerClass = value
            })
            entry.SetSelected(trainer.TrainerClass)
            form.Add(label)
            form.Add(entry)
        } else if fieldName == "TrainerName" {
            label := widget.NewLabel(fieldName)
            entry:= NewCustomEntry(i)
            entry.SetText(field.Interface().(string))
            entry.OnChanged = func(s string) {
                if len(s) >= 3 {
                    structValue.Field(entry.Id).SetString(s)
                }
            }
            form.Add(label)
            form.Add(entry)
        } else if fieldName == "EncounterMusicGender" {
            label := widget.NewLabel(fieldName)
            value:= widget.NewSelect(trainerEncounterMusics, func(value string) {
                trainer.EncounterMusicGender = value
            })
            value.SetSelected(trainer.EncounterMusicGender)
            form.Add(label)
            form.Add(value)
        } else if fieldName == "TrainerPic" {
            label := widget.NewLabel(fieldName)
            entry := NewCompletionEntry(trainerPics)
            trainerPic = canvas.NewImageFromFile(trainerPicMap[trainer.TrainerPic])
            trainerPic.FillMode = canvas.ImageFillContain
            trainerPic.SetMinSize(fyne.NewSize(64, 64))
            entry.SetText(trainer.TrainerPic)
            entry.OnChanged = func(s string) {
                if len(s) == 0 {
                    entry.HideCompletion()
                    return
                }
                var filteredItems []string
                for _, item := range(trainerPics) {
                    if strings.Contains(item, s) {
                        filteredItems = append(filteredItems, item)
                    }
                }
                if len(s) >= 5 {
                    entry.SetOptions(filteredItems)
                    entry.ShowCompletion()
                }
                _, has := trainerPicMap[s]
                if has {
                    trainer.TrainerPic = s
                    trainerPic = canvas.NewImageFromFile(trainerPicMap[trainer.TrainerPic])
                    trainerPic.FillMode = canvas.ImageFillContain
                    trainerPic.SetMinSize(fyne.NewSize(64, 64))
                    trainerPic.Refresh()
                    picWrapper.Objects[1] = trainerPic
                    picWrapper.Refresh()
                }
            }
            labelWrapper := container.New(layout.NewFormLayout(), label, entry)
            picWrapper = container.NewVBox(labelWrapper, trainerPic)
            content.Add(picWrapper)
            continue
        } else if fieldName == "Items" {
            for j:=0; j < 4; j++ {
                label := widget.NewLabel(fieldName + " " + strconv.Itoa(j))
                selectBox := NewCompletionEntry(items)
                form.Add(label)
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
                    var filteredItems []string
                    for _, item := range(items) {
                        if strings.Contains(item, s) {
                            filteredItems = append(filteredItems, item)
                        }
                    }
                    if len(s) >= 5 {
                        selectBox.SetOptions(filteredItems)
                        selectBox.ShowCompletion()
                    }
                    if len(s) >= 9 {
                        if sliceContains(items, s) {
                            trainer.Items[selectBox.Id] = s;
                        }
                    }
                }
                form.Add(selectBox)
            }
        } else if fieldName == "AiFlags"{
            // Chunk the ai flags
            // Arbitrary 3 chunks
            var NUM_COLUMNS int = 3
            chunks := len(aiFlags) / NUM_COLUMNS
            if len(aiFlags) % NUM_COLUMNS != 0 {
                chunks++
            }
            label := widget.NewLabel(fieldName)
            form.Add(label)
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
            form.Add(container.NewHScroll(checkGroupHolder))
        } else if fieldName == "DoubleBattle"{
            label := widget.NewLabel(fieldName)
            check := widget.NewCheck("", func(checked bool){
                fmt.Printf("%t\n", checked)
                trainer.DoubleBattle = checked
            })
            checked := field.Interface().(bool)
            check.Checked = checked
            form.Add(label)
            form.Add(check)
        }
    }
    content.Add(form)
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

func buildgTrainerFrontPicPath(path string) string {
    buf := strings.Builder{}
    buf.WriteString(path)
    buf.WriteString("/src")
    buf.WriteString("/data")
    buf.WriteString("/graphics")
    buf.WriteString("/trainers.h")
    return buf.String()
}

func buildTrainerSpritePath(path string) string {
    buf := strings.Builder{}
    buf.WriteString(path)
    buf.WriteString("/src")
    buf.WriteString("/data")
    buf.WriteString("/trainer_graphics")
    buf.WriteString("/front_pic_tables.h")
    return buf.String()
}

func buildTrainerEncounterMusicPath(path string) string {
    buf := strings.Builder{}
    buf.WriteString(path)
    buf.WriteString("/include")
    buf.WriteString("/constants")
    buf.WriteString("/trainers.h")
    return buf.String()
}
