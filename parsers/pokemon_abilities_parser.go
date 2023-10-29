package parsers

import (
    "strings"
    "os"
    "fmt"
)

func ParsePokemonAbilities(filepath string) []string {
    file, err := os.ReadFile(filepath)
    if err != nil {
        fmt.Println("Error: could not read file: " + filepath)
        return nil
    }
    fileContents := string(file)
    var abilities []string
    for _, line := range strings.Split(fileContents, "\n") {
        if strings.Contains(line, "ABILITY_") {
            start := strings.Index(line, "ABILITY")
            end := GetNthIndex(line, 2, ' ')
            if start > end {
                fmt.Println("Could not parse line: " + line)
                continue
            }
            abilities = append(abilities, strings.TrimSpace(line[start:end]))
        }
    }
    return abilities
}
