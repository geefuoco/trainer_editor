package parsers

import (
    "strings"
    "os"
    "github.com/geefuoco/trainer_editor/logging"
)

func ParseMoves(filepath string) []string {
    file, err := os.ReadFile(filepath)
    if err != nil {
        logging.ErrorLog(" could not read file: " + filepath)
        return nil
    }

    fileContents := string(file)
    var moves []string
    for _, line := range strings.Split(fileContents, "\n") {
        if strings.Contains(line, "MOVE") {
            if strings.Contains(line, "MOVES_COUNT") {
                break
            }
            start := strings.Index(line, "MOVE")
            end := GetNthIndex(line, 2, ' ')
            if start == -1 || end == -1 {
                logging.WarnLog("Could not parse line: " + line)
                continue
            }
            moves = append(moves, line[start:end])
        }
    }
    return moves
}
