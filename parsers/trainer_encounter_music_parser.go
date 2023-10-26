package parsers

import (
    "os"
    "fmt"
    "strings"
)

func ParseTrainerEncounterMusic(input string) []string {
    file, err := os.ReadFile(input)
    if err != nil {
        fmt.Println("Error: could not open file: " + input)
        return nil
    }

    fileContents := string(file)
    var music []string
    for _, line := range(strings.Split(fileContents, "\n")) {
        if strings.Contains(line, "TRAINER_ENCOUNTER_MUSIC") {
            start := strings.Index(line, "TRAINER_ENCOUNTER_MUSIC")
            end := GetNthIndex(line, 2, ' ')
            if start > end{
                fmt.Println("Error: could not parse line: \n" + line) 
                continue
            }
            item := strings.TrimSpace(line[start:end])
            music = append(music, item)
        }
    }
    return music
}
