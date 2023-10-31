package data_objects

import (
    "strings"
    "reflect"
    "os"
    "bufio"
    "github.com/geefuoco/trainer_editor/logging"
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

func SaveAll(trainerFilePath string, trainerPartiesFilePath string, trainers []*Trainer, trainerParties []*TrainerParty) error {
    err := SaveTrainers(trainerFilePath, trainers)
    if err != nil {
        return err
    }
    err = SaveTrainerParties(trainerPartiesFilePath, trainerParties)
    return err
}

func SaveTrainers(filepath string, trainers []*Trainer) error {
    file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0666)
    if err != nil {
        return err
    }
    defer file.Close()
    writer := bufio.NewWriter(file)
    _, err = writer.WriteString("const struct Trainer gTrainers[] = {\n")
    if err != nil {
        return err;
    }

    for i, trainer := range trainers {
        _, err = writer.WriteString(trainer.String())
        if err != nil {
            return err
        }
        if i != len(trainers)-1{
            _, err = writer.WriteString("\n")
            if err != nil {
                return err
            }
        }
    }
    _, err = writer.WriteString("};")
    if err != nil {
        return err
    }
    writer.Flush()
    logging.InfoLog("Saved %d trainers to %s", len(trainers), filepath)
    return nil
}

func SaveTrainerParties(filepath string, trainerParties []*TrainerParty) error {
    file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0666)
    if err != nil {
        return err
    }
    defer file.Close()
    writer := bufio.NewWriter(file)

    for _, trainerParty := range trainerParties {
        _, err = writer.WriteString(trainerParty.String())
        if err != nil {
            return err
        }
        _, err = writer.WriteString("\n")
        if err != nil {
            return err
        }
    }
    logging.InfoLog("Saved %d trainer parties to %s", len(trainerParties), filepath)
    return writer.Flush()
}

func (t* Trainer) GetPartyName() string {
    if t.Party == "NULL" {
        return t.Party
    }
    if len(t.Party) > 0 {
        return t.Party[len("TRAINER_PARTY("):len(t.Party)-1]
    }
    return ""
}

func (t *Trainer) String() string {
    var b strings.Builder

    templateItems := [4]string{"ITEM_NONE","ITEM_NONE","ITEM_NONE","ITEM_NONE"}
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
    if strings.Contains(t.EncounterMusicGender, "|") {
        split := strings.Split(t.EncounterMusicGender, "|")
        for i, sp := range split {
            b.WriteString(sp)
            if i != len(split)-1{ 
                b.WriteString(" | ")
            }
        }
        b.WriteString(",\n")
    } else {
        b.WriteString(t.EncounterMusicGender+ ",\n")
    }
    // Trainer Pic
    b.WriteString(padding)
    b.WriteString(padding)
    b.WriteString(".trainerPic = ")
    b.WriteString(t.TrainerPic + ",\n")
    // Trainer Name
    b.WriteString(padding)
    b.WriteString(padding)
    b.WriteString(".trainerName = _(\"")
    if strings.Contains(t.TrainerName, "&") {
        split := strings.Split(t.TrainerName, "&")
        for i, sp := range split {
            b.WriteString(sp)
            if i != len(split)-1{
                b.WriteString(" & ")
            }
        }
        b.WriteString("\"),\n")
    } else {
        b.WriteString(t.TrainerName+ "\"),\n")
    }
    // Items
    b.WriteString(padding)
    b.WriteString(padding)
    b.WriteString(".items = {")
    if !reflect.DeepEqual(t.Items, templateItems) {
        for i, v := range t.Items {
            if v != "" {
                b.WriteString(v)
                if i != 3 {
                    b.WriteString(", ")
                }
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
    // Special case for NONE trainer
    if t.TrainerKey == "TRAINER_NONE" {
        b.WriteString(".partySize = 0,\n")
        b.WriteString(padding)
        b.WriteString(padding)
    } 
    b.WriteString(".party = ")
    b.WriteString(t.Party+",\n")
    b.WriteString(padding)
    b.WriteString("},\n")
    return b.String()
}
