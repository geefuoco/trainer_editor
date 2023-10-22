package parsers

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "github.com/geefuoco/trainer_editor/data_objects"
)

func ParseTrainerParties(filepath string) []*data_objects.TrainerParty{
    file, err := os.ReadFile(filepath)
    if err != nil {
        fmt.Println("Error when opening file: ", err)
        return nil
    }

    var parties []*data_objects.TrainerParty
    fileContents := string(file)

    trainerParties := strings.Split(fileContents, ";")
    for _, rawTrainerParty := range(trainerParties) {
        rawTrainerParty = rawTrainerParty[:len(rawTrainerParty)-1] // Remove the last }

        currentParty := &data_objects.TrainerParty{}
        currentMon := &data_objects.TrainerMon{}
        for _, line := range(strings.Split(rawTrainerParty, "\n")){ 
            line = strings.ReplaceAll(line, " ", "")
            // Incase of some windows BS
            line = strings.ReplaceAll(line, "\r", "")

            if strings.Contains(line, "TrainerMon") {
                start := strings.Index(line, "TrainerMon") + len("TrainerMon")
                if start == -1 {
                    panic("Error: Malformatted TrainerMon struct")
                }
                endOffset := 4
                trainer := line[start:len(line)-endOffset]
                currentParty.Trainer = trainer
            } else if len(line) == 1 && line[0] == '{' {
                // The least amount of fields a trainer can have is LVL and Species
                if currentMon.Lvl != 0 && currentMon.Species != "" {
                    currentParty.Party = append(currentParty.Party, currentMon)
                    currentMon = &data_objects.TrainerMon{}
                } 
            }else if strings.Contains(line, ".iv") {
                start := strings.IndexByte(line, '=')
                if start == -1 {
                    panic("Error: Malformatted TrainerMon struct")
                }
                startOffset := len("=TRAINER_PARTY_IVS(")
                endOffset := 2 // ),
                var ivs [6]uint64
                for i, iv := range strings.Split(line[start+startOffset:len(line)-endOffset], ",") {
                    if !(iv ==  "" || iv == "\n") {
                        value, err := strconv.ParseUint(strings.TrimSpace(iv), 10, 64)
                        if err == nil {
                            ivs[i] = value
                        }
                    } 
                }
                currentMon.Iv = ivs
            } else if strings.Contains(line, ".lvl") {
                start := strings.Index(line, "=")
                if start == -1 {
                    panic("Error: Malformatted TrainerMon struct")
                }
                startOffset := 1
                endOffset := 1
                lvl := strings.TrimSpace(line[start+startOffset:len(line)-endOffset])
                value, err := strconv.ParseUint(lvl, 10, 64)
                if err != nil {
                    fmt.Println("Error: Could not read level. Setting default")
                    currentMon.Lvl = 5
                } else {
                    currentMon.Lvl = value
                }
            } else if strings.Contains(line, ".species") {
                start := strings.IndexByte(line, '=')
                if start == -1 {
                    panic("Error: Malformatted TrainerMon struct")
                }
                startOffset := 1
                endOffset := 1
                species := line[start+startOffset:len(line)-endOffset]
                currentMon.Species = species
            } else if strings.Contains(line, ".ev") {
                start := strings.IndexByte(line, '=')
                if start == -1 {
                    panic("Error: Malformatted TrainerMon struct")
                }
                startOffset := len("=TRAINER_PARTY_EVS(")
                endOffset := 2 // ),
                var evs [6]uint64
                for i, ev := range strings.Split(line[start+startOffset:len(line)-endOffset], ",") {
                    if !(ev ==  "" || ev == "\n") {
                        value, err := strconv.ParseUint(strings.TrimSpace(ev), 10, 64)
                        if err == nil {
                            evs[i] = value
                        }
                    } 
                }
                currentMon.Ev = evs
            } else if strings.Contains(line, ".heldItem") {
                start := strings.IndexByte(line, '=')
                if start == -1 {
                    panic("Error: Malformatted TrainerMon struct")
                }
                startOffset := 1
                endOffset := 1
                heldItem := line[start+startOffset:len(line)-endOffset]
                currentMon.HeldItem = heldItem
            } else if strings.Contains(line, ".moves") {
                moves := [4]string{"MOVE_NONE", "MOVE_NONE", "MOVE_NONE", "MOVE_NONE"} 
                start := strings.IndexByte(line, '=')
                if start == -1 {
                    panic("Error: Malformatted Trainer struct")
                }
                startOffset := 2 // ={
                endOffset := 2 // },
                for i, item := range strings.Split(line[start+startOffset:len(line)-endOffset], ",") {
                    if !(item ==  "" || item == "\n") {
                        moves[i] = strings.TrimSpace(item)
                    }
                }
                currentMon.Moves = moves
            } else if strings.Contains(line, ".ability") {
                start := strings.IndexByte(line, '=')
                if start == -1 {
                    panic("Error: Malformatted TrainerMon struct")
                }
                startOffset := 1
                endOffset := 1
                ability := line[start+startOffset:len(line)-endOffset]
                currentMon.Ability = ability
            } else if strings.Contains(line, ".isShiny") {
                if strings.Contains(line, "TRUE") {
                    currentMon.IsShiny = true
                } else {
                    currentMon.IsShiny = false
                }
            }
        }
        if currentMon.Lvl != 0 && currentMon.Species != "" {
            currentParty.Party = append(currentParty.Party, currentMon)
            currentMon = &data_objects.TrainerMon{}
        } 
        if len(currentParty.Party) != 0 {
            parties = append(parties, currentParty)
        }
    }

    return parties
}
