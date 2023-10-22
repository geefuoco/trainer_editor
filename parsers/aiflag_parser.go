package parsers

import (
    "os"
    "fmt"
    "strings"
)

func ParseAiFlags(filepath string) []string {
    file, err := os.ReadFile(filepath)
    if err != nil {
        fmt.Println("Error: Could not open file: " + filepath)
        return nil
    }

    fileContents := string(file)
    var flags []string

    for _, line := range(strings.Split(fileContents, "\n")) {

        if strings.Contains(line, "COUNT") {
            break
        }

        if strings.Contains(line, "STALL") {
            continue
        } else if strings.Contains(line, "SCREEN") {
            continue
        } else if strings.Contains(line, "AI_FLAG") {
            start := strings.Index(line, "AI_FLAG")
            var index uint = 2
            endOffset := getNthIndex(line, index, ' ')
            if endOffset == -1 {
                fmt.Printf("Could not find %d index of ' ' in line: %s\n", index, line)
                continue
            }
            flag := strings.TrimSpace(line[start:endOffset])
            flags = append(flags, flag)
        }
    }

    return flags
}

func getNthIndex(str string, index uint, query byte) int {
    var count uint
    for i := range(str) {
        if count > index {
            return -1
        }
        if str[i] == query{
            count += 1
            if count == index {
                return i
            }
        }
    }
    return -1
}
