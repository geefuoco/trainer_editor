package data_objects

import (
    "strings"
)

type Trainer struct{
    TrainerClass string
    EncounterMusicGender string
    TrainerPic string
    TrainerName string
    Items []string
    DoubleBattle bool 
    AiFlags []string
    Party string
}

func (t *Trainer) String() string {
    var b strings.Builder
    b.WriteString("Trainer: \n")
    b.WriteString("  TrainerClass: ")
    b.WriteString(t.TrainerClass + "\n")
    b.WriteString("  EncounterMusicGender: ")
    b.WriteString(t.EncounterMusicGender+ "\n")
    b.WriteString("  TrainerPic: ")
    b.WriteString(t.TrainerPic + "\n")
    b.WriteString("  TrainerName: ")
    b.WriteString(t.TrainerName+ "\n")
    b.WriteString("  TrainerName: ")
    if len(t.Items) == 0 {
        b.WriteString("  Items: {}\n")
    } else {
        b.WriteString("  Items: {\n")
        for i := range(t.Items) {
            b.WriteString("    " + t.Items[i] + "\n")
        }
        b.WriteString("  }\n")
    }
    b.WriteString("  DoubleBattle: ")
    if t.DoubleBattle {
        b.WriteString("true\n")
    } else {
        b.WriteString("false\n")
    }
    if len(t.AiFlags) == 0 {
        b.WriteString("  AiFlags: {}\n")
    } else {
        b.WriteString("  AiFlags: {\n")
        for i := range(t.AiFlags) {
            b.WriteString("    " + t.AiFlags[i] + "\n")
        }
        b.WriteString("  }\n")
    }
    b.WriteString("  Party: ")
    b.WriteString(t.Party+ "\n")
    return b.String()
}
