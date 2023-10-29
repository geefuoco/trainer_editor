package data_objects

import (
    "strings"
)

type Trainer struct{
    TrainerKey string
    TrainerClass string
    EncounterMusicGender string
    TrainerPic string
    TrainerName string
    Items [4]string
    DoubleBattle bool 
    AiFlags []string
    Party string
}

func (t* Trainer) GetPartyName() string {
    if len(t.Party) > 0 {
        return t.Party[len("TRAINER_PARTY("):len(t.Party)-1]
    }
    return ""
}

func (t *Trainer) String() string {
    var b strings.Builder
    padding := "    "
    // Trainer Key
    b.WriteString(padding)
    b.WriteString("[" + t.TrainerKey + "] =\n")
    b.WriteString(padding)
    b.WriteString("{\n")
    // Trainer Class
    b.WriteString(padding)
    b.WriteString(padding)
    b.WriteString(".trainerClass = ")
    b.WriteString(t.TrainerClass + ",\n")
    // Encounter Music
    b.WriteString(padding)
    b.WriteString(padding)
    b.WriteString(".encounterMusic_gender = ")
    b.WriteString(t.EncounterMusicGender+ ",\n")
    // Trainer Pic
    b.WriteString(padding)
    b.WriteString(padding)
    b.WriteString(".trainerPic = ")
    b.WriteString(t.TrainerPic + ",\n")
    // Trainer Name
    b.WriteString(padding)
    b.WriteString(padding)
    b.WriteString(".trainerName = _(\"")
    b.WriteString(t.TrainerName+ "\"),\n")
    // Items
    b.WriteString(padding)
    b.WriteString(padding)
    b.WriteString(".items = {")
    for i, v := range t.Items {
        if v != "" && v != "ITEM_NONE" {
            b.WriteString(v)
            if i != 3 {
                b.WriteString(",")
            }
        }
    }
    b.WriteString("},\n")
    // Double Battle
    b.WriteString(padding)
    b.WriteString(padding)
    b.WriteString(".doubleBattle = ")
    if t.DoubleBattle {
        b.WriteString("TRUE,\n")
    } else {
        b.WriteString("FALSE,\n")
    }
    // AI Flags
    b.WriteString(padding)
    b.WriteString(padding)
    if len(t.AiFlags) == 0 {
        b.WriteString(".aiFlags = 0,\n")
    } else {
        b.WriteString(".aiFlags = ")
        for i := range(t.AiFlags) {
            b.WriteString(t.AiFlags[i])
            if i < len(t.AiFlags)-1 {
                b.WriteString(" | ")
            }
        }
        b.WriteString(",\n")
    }
    b.WriteString(padding)
    b.WriteString(padding)
    b.WriteString(".party = ")
    b.WriteString(t.Party+",\n")
    b.WriteString(padding)
    b.WriteString("},\n")
    return b.String()
}
