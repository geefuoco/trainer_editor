package parsers

import (
    "strings"
    "os"
    "github.com/geefuoco/trainer_editor/logging"
)

func ParsePokemonPics(speciesNamesFilePath, speciesSpritesFilePath string) map[string]string {
    frontPics := parsePokemonSprites(speciesSpritesFilePath)
    speciesNames := parseSpeciesNames(speciesNamesFilePath)

    output := make(map[string]string)
    for key, value := range(speciesNames) {
        _, has := frontPics[value]
        if has {
            output[key] = frontPics[value]
        }
    }

    return output
}

func parseSpeciesNames(filepath string) map[string]string {
    file, err := os.ReadFile(filepath)
    if err != nil {
        logging.ErrorLog("could not read file: " + filepath)
        return nil
    }

    fileContents := string(file)
    mons := make(map[string]string)
    for _, line := range strings.Split(fileContents, "\n") {
        if strings.Contains(line, "SPECIES_SPRITE") {
            if strings.Contains(line, "gMonFrontPicTableFemale"){ 
                // TODO
                // No good way to handle female sprites at the moment
                break
            }
            start := strings.Index(line, "SPECIES_SPRITE(") + len("SPECIES_SPRITE(")
            split := strings.Split(line, ",")
            if start == -1 {
                logging.WarnLog("Error: Could not parse line: " + line)
                continue
            }
            key := strings.TrimSpace("SPECIES_"+split[0][start:])
            value := strings.TrimSpace(split[1][:len(split[1])-1])

            if value[len(value)-1] == 'F' {
                // TODO 
                // Figure out how to handle female sprites
                continue
            }
            mons[key] = value
        }
    }
    return mons
}

func parsePokemonSprites(filepath string) map[string]string {
    file, err := os.ReadFile(filepath)
    if err != nil {
        logging.ErrorLog("could not read file: " + filepath)
        return nil
    }

    fileContents := string(file)
    pics :=make(map[string]string)
    for _, line := range strings.Split(fileContents, "\n") {
        if strings.Contains(line, ".4bpp") && strings.Contains(line, "gMonFrontPic") {
            if strings.Contains(line, "frontf") {
                continue
            }
            start := strings.Index(line, "gMonFrontPic_")
            end := strings.IndexByte(line, '=') - 3 // []
            if start == -1 {
                logging.WarnLog("Could not parse line: " + line)
                continue
            }
            key := strings.TrimSpace(line[start:end])
            // const u32 gMonFrontPic_CalyrexIceRider[] = INCBIN_U32("graphics/pokemon/calyrex/ice_rider/front.4bpp.lz");
            start = strings.IndexByte(line, '"') + len("\"")
            end = len(line) - len(".4bpp.lz\");")
            value := strings.TrimSpace(line[start:end]+".png")
            pics[key]=value
        }
    }
    return pics
}
