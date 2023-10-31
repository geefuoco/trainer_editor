package parsers

import (
    "strings"
    "os"
    "github.com/geefuoco/trainer_editor/logging"
)

func ParseItems(filepath string) []string {
    file, err := os.ReadFile(filepath)
    if err != nil {
        logging.ErrorLog("when opening file: "+filepath)
        return nil
    }

    fileContents := string(file)
    var items []string

    for _, line := range(strings.Split(fileContents, "\n")) {
        if strings.Contains(line, "ITEMS_COUNT") {
            break;
        }
        if strings.Contains(line, "ITEM") {
            start := strings.Index(line, "ITEM")
            startOffset := start
            endOffset := strings.LastIndexByte(line, ' ')
            if startOffset > endOffset {
                logging.WarnLog("could not parse line: \n" + line) 
                continue
            }
            item := strings.TrimSpace(line[startOffset:endOffset])
            items = append(items, item)
        }
    }
    return items
}
