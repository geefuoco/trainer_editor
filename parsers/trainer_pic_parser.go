package parsers

import (
    "os"
    "strings"
    "github.com/geefuoco/trainer_editor/logging"
)
func ParseTrainerPics(gTrainerFrontPicPath, trainerSpritePath string) map[string]string {
    frontPics := parseTrainerFrontPics(gTrainerFrontPicPath)
    picKeys := parseTrainerPicKeys(trainerSpritePath)

    output := make(map[string]string)
    for key, value := range(picKeys) {
        _, has := frontPics[value]
        if has {
            output[key] = frontPics[value]
        }
    }
    return output
}

func parseTrainerPicKeys(filepath string) map[string]string {
    file, err := os.ReadFile(filepath)
    if err != nil {
        logging.ErrorLog("could not read file: " + filepath)
        return nil
    }

    output := make(map[string]string)
    fileContents := string(file)

    for _, line := range(strings.Split(fileContents, "\n")) {
        if strings.Contains(line, "TRAINER_SPRITE") {
            start := strings.Index(line, "TRAINER_SPRITE(") + len("TRAINER_SPRITE(")
            split := strings.Split(line, ",")
            // key
            key := strings.TrimSpace("TRAINER_PIC_" + split[0][start:])
            // value
            value := strings.TrimSpace(split[1])
            output[key] = value
        }
    }
    return output
}

func parseTrainerFrontPics(filepath string) map[string]string{
    file, err := os.ReadFile(filepath)
    if err != nil {
        logging.ErrorLog("Error: could not read file: " + filepath)
        return nil
    }

    output := make(map[string]string)
    fileContents := string(file)

    for _, line := range(strings.Split(fileContents, "\n")) {
        if strings.Contains(line, ".4bpp") && strings.Contains(line, "INCBIN_U32") {
            // key
            start := strings.Index(line, "gTrainerFrontPic_")
            end := strings.IndexByte(line, '=') - 3 // []
            if start == -1 {
                logging.WarnLog(" could not parse line: " + line)
                continue
            }
            key := strings.TrimSpace(line[start:end])
            //value 
            start = strings.Index(line, "INCBIN_U32(\"") + len("INCBIN_U32(\"")
            end = len(line) - len(".4bpp.lz\");") // .4bpp.lz");
            value := strings.TrimSpace(line[start:end] + ".png")
            // trainer front pic -> pic path
            output[key]=value
        }
    }
    return output
}
