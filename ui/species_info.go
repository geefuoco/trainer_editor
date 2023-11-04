package ui

import (
    "github.com/geefuoco/trainer_editor/data_objects"
    "github.com/geefuoco/trainer_editor/logging"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "strconv"
)

func createSpeciesInfo(speciesName string) *fyne.Container{
    content := container.NewVBox()
    label := widget.NewLabel("Base Stats")
    form := container.New(layout.NewFormLayout())
    var species *data_objects.PokemonSpeciesInfo
    for _, x := range speciesInfo {
        if speciesName == x.Species{ 
            species = x
            break
        }
    }
    if species == nil {
        logging.ErrorLog("Could not find species info with species: %s", speciesName)
        return nil
    }
    // HP
    hpLabel := widget.NewLabel("HP")
    hpLabelValue := widget.NewLabel(strconv.FormatUint(species.BaseHp, 10))
    form.Add(hpLabel)
    form.Add(hpLabelValue)
    // ATK
    atkLabel := widget.NewLabel("ATK")
    atkLabelValue := widget.NewLabel(strconv.FormatUint(species.BaseAtk, 10))
    form.Add(atkLabel)
    form.Add(atkLabelValue)
    // DEF
    defLabel := widget.NewLabel("DEF")
    defLabelValue := widget.NewLabel(strconv.FormatUint(species.BaseDef, 10))
    form.Add(defLabel)
    form.Add(defLabelValue)
    // SPD
    spdLabel := widget.NewLabel("SPD")
    spdLabelValue := widget.NewLabel(strconv.FormatUint(species.BaseSpd, 10))
    form.Add(spdLabel)
    form.Add(spdLabelValue)
    // SPATK
    spAtkLabel := widget.NewLabel("SPATK")
    spAtkLabelValue := widget.NewLabel(strconv.FormatUint(species.BaseSpAtk, 10))
    form.Add(spAtkLabel)
    form.Add(spAtkLabelValue)
    // SPDEF
    spDefLabel := widget.NewLabel("SPDEF")
    spDefLabelValue := widget.NewLabel(strconv.FormatUint(species.BaseSpDef, 10))
    form.Add(spDefLabel)
    form.Add(spDefLabelValue)
    // Types Label
    typesLabel := widget.NewLabel("Types")
    typeContainer := container.New(layout.NewFormLayout())

    type1 := widget.NewLabel(species.Types[0])
    type2 := widget.NewLabel(species.Types[1])
    typeContainer.Add(type1)
    typeContainer.Add(type2)

    form.Add(typesLabel)
    form.Add(typeContainer)

    content.Add(label)
    content.Add(form)
    return content
}

