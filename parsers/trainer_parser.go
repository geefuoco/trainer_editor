package parsers

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "github.com/geefuoco/trainer_editor/data_objects"
)

func ParseTrainers(filepath string) []*data_objects.Trainer {
    file, err := os.Open(filepath)
    defer file.Close()
    if err != nil {
        fmt.Println("Error when opening file: ", err)
        return nil
    }

    var trainers []*data_objects.Trainer
    currentTrainer := &data_objects.Trainer{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        line = strings.ReplaceAll(line, " ", "")
        // Strip the line of whitespaces
        if strings.Contains(line, "[TRAINER_") {
            if currentTrainer.TrainerKey != "" {
                trainers = append(trainers, currentTrainer)
                currentTrainer = &data_objects.Trainer{}
            }
            start := strings.Index(line, "TRAINER")
            end := len(line) - 2 // ]=
            currentTrainer.TrainerKey = line[start:end]
        }else if strings.Contains(line, ".trainerClass") {
            start := strings.IndexByte(line, '=')
            if start == -1 {
                panic("Error: Malformatted Trainer struct")
            }
            startOffset := 1
            trainerClass := line[start+startOffset:len(line)-1]
            currentTrainer.TrainerClass = trainerClass
        } else if strings.Contains(line, ".encounterMusic_gender") {
            start := strings.IndexByte(line, '=')
            if start == -1 {
                panic("Error: Malformatted Trainer struct")
            }
            startOffset := 1
            encounterMusic := line[start+startOffset:len(line)-1]
            currentTrainer.EncounterMusicGender = encounterMusic
        } else if strings.Contains(line, ".trainerPic") {
            start := strings.IndexByte(line, '=')
            if start == -1 {
                panic("Error: Malformatted Trainer struct")
            }
            startOffset:=1
            trainerPic:= line[start+startOffset:len(line)-1]
            currentTrainer.TrainerPic= trainerPic
        } else if strings.Contains(line, ".trainerName") {
            start := strings.IndexByte(line, '=')
            if start == -1 {
                panic("Error: Malformatted Trainer struct")
            }
            startOffset := 4 // =_("
            endOffset := 3 // "),
            trainerName:= line[start+startOffset:len(line)-endOffset]
            currentTrainer.TrainerName= trainerName
        } else if strings.Contains(line, ".items") {
            items := [4]string{"ITEM_NONE", "ITEM_NONE", "ITEM_NONE", "ITEM_NONE"}
            start := strings.IndexByte(line, '=')
            if start == -1 {
                panic("Error: Malformatted Trainer struct")
            }
            startOffset := 2 // ={
            endOffset := 2 // },
            for i, item := range strings.Split(line[start+startOffset:len(line)-endOffset], ",") {
                if !(item ==  "" || item == "\n") {
                    items[i] = strings.TrimSpace(item)
                }
            }
            currentTrainer.Items = items
        } else if strings.Contains(line, ".doubleBattle") {
            start := strings.IndexByte(line, '=')
            if start == -1 {
                panic("Error: Malformatted Trainer struct")
            }
            var doubleBattle bool
            if strings.Contains(line, "FALSE") {
                doubleBattle = false
            } else {
                doubleBattle = true
            }
            currentTrainer.DoubleBattle = doubleBattle
        } else if strings.Contains(line, ".aiFlags") {
            start := strings.IndexByte(line, '=')
            if start == -1 {
                panic("Error: Malformatted Trainer struct")
            }
            startOffset := 1
            endOffset := 1
            // Currently there is only 18 AI Flags Available
            aiFlags := make([]string, 0, 20)
            if strings.Contains(line, "AI_FLAG") {
                aiFlags = strings.Split(line[start+startOffset:len(line)-endOffset], "|")
            }
            currentTrainer.AiFlags = aiFlags
        } else if strings.Contains(line, ".party") && !strings.Contains(line, ".partySize"){
            start := strings.IndexByte(line, '=')
            if start == -1 {
                panic("Error: Malformatted Trainer struct")
            }
            startOffset := 1
            endOffset := 1
            var party string
            if strings.Contains(line, "NULL") {
                party = "NULL"
            } else {
                if strings.HasSuffix(line, ",") {
                    party = line[start+startOffset:len(line)-endOffset]
                } else {
                    party = line[start+startOffset:len(line)]
                }
            }
            currentTrainer.Party = party
        }
    }
    trainers = append(trainers, currentTrainer)

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file: ", err)
        return nil
    }
    return trainers
}
