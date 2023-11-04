package parsers

import (
    "github.com/geefuoco/trainer_editor/logging"
    "github.com/geefuoco/trainer_editor/data_objects"
    "strings"
    _ "embed"
    "strconv"
)

//go:embed species_info.txt
var SpeciesFileContents string

func ParsePokemonSpeciesInfo(fileContents string) []*data_objects.PokemonSpeciesInfo {

    speciesInfos := []*data_objects.PokemonSpeciesInfo {
        {
            Species: "SPECIES_NONE",
            BaseHp: 0,
            BaseAtk: 0,
            BaseDef: 0,
            BaseSpd: 0,
            BaseSpAtk: 0,
            BaseSpDef: 0,
            Types: [2]string{"TYPE_NONE", "TYPE_NONE"},
        },
    }
    var skip bool
    currentSpeciesInfo := &data_objects.PokemonSpeciesInfo{}

    for _, line := range strings.Split(fileContents, "\n") {
        if skip && strings.Contains(line, "base") {
            continue
        }
        if strings.Contains(line, "SPECIES") {
            skip = false
            if strings.Contains(line, "SPECIES_NONE") {
                continue
            }
            start :=strings.IndexByte(line, '[')+1
            end := strings.IndexByte(line, ']')
            if start == -1 || end == -1 || start > end{
                logging.WarnLog("Could not parse line: "+line)
                skip=true
                continue
            }

            if currentSpeciesInfo.Species != "" {
                speciesInfos = append(speciesInfos, currentSpeciesInfo) 
            } 
            currentSpeciesInfo = &data_objects.PokemonSpeciesInfo{
                Species: strings.TrimSpace(line[start:end]),
            }
        } else if strings.Contains(line, ".base") {
            if currentSpeciesInfo.Species == "" {
                continue
            }
            var start int
            var end int
            if strings.Contains(line, "P_UPDATED_STATS") {
                start = strings.IndexByte(line, '?')+1
                end = strings.IndexByte(line, ':')
            } else {
                start = strings.IndexByte(line, '=')
                end = strings.IndexByte(line, ',')
            }
            if start == -1 || end == -1 || start > end {
                logging.WarnLog("Could not parse line: "+line)
                skip=true
                continue
            }
            start += 1
            if strings.Contains(line, "HP") {
                if currentSpeciesInfo.BaseHp != 0 {
                    continue
                }
                result := strings.TrimSpace(line[start:end])
                value, err := strconv.ParseUint(result, 10, 64)
                if err != nil {
                    logging.ErrorLog("Could not parse value to uint64: "+result)
                    skip=true
                    continue
                }
                currentSpeciesInfo.BaseHp = value
            } else if strings.Contains(line, "baseAttack") {
                if currentSpeciesInfo.BaseAtk != 0 {
                    continue
                }
                result := strings.TrimSpace(line[start:end])
                value, err := strconv.ParseUint(result, 10, 64)
                if err != nil {
                    logging.ErrorLog("Could not parse value to uint64: "+result)
                    skip=true
                    continue
                }
                currentSpeciesInfo.BaseAtk = value
            } else if strings.Contains(line, "baseDefense") {
                if currentSpeciesInfo.BaseDef != 0 {
                    continue
                }
                result := strings.TrimSpace(line[start:end])
                value, err := strconv.ParseUint(result, 10, 64)
                if err != nil {
                    logging.ErrorLog("Could not parse value to uint64: "+result)
                    skip=true
                    continue
                }
                currentSpeciesInfo.BaseDef = value
            } else if strings.Contains(line, "SpAttack") {
                if currentSpeciesInfo.BaseSpAtk != 0 {
                    continue
                }
                result := strings.TrimSpace(line[start:end])
                value, err := strconv.ParseUint(result, 10, 64)
                if err != nil {
                    logging.ErrorLog("Could not parse value to uint64: "+result)
                    skip=true
                    continue
                }
                currentSpeciesInfo.BaseSpAtk = value
            } else if strings.Contains(line, "Speed") {
                if currentSpeciesInfo.BaseSpd != 0 {
                    continue
                }
                result := strings.TrimSpace(line[start:end])
                value, err := strconv.ParseUint(result, 10, 64)
                if err != nil {
                    logging.ErrorLog("Could not parse value to uint64: "+result)
                    skip=true
                    continue
                }
                currentSpeciesInfo.BaseSpd  = value
            } else if strings.Contains(line, "baseSpDefense") {
                if currentSpeciesInfo.BaseSpDef != 0 {
                    continue
                }
                result := strings.TrimSpace(line[start:end])
                value, err := strconv.ParseUint(result, 10, 64)
                if err != nil {
                    logging.ErrorLog("Could not parse value to uint64: "+result)
                    skip=true
                    continue
                }
                currentSpeciesInfo.BaseSpDef = value
            } else {
                logging.WarnLog("Unreachable code: "+line)
            }
        } else if strings.Contains(line, ".types") {
            if currentSpeciesInfo.Species == "" {
                continue
            }

            start := strings.IndexByte(line, '{')
            end := strings.IndexByte(line, '}')
            if start == -1 || end == -1 || start > end {
                logging.WarnLog("Could not parse line: "+line)
                skip=true
                continue
            }
            start += 1
            rawTypes := strings.ReplaceAll(line[start:end], " ", "")
            copiedTypes := [2]string{}
            copy(copiedTypes[:], strings.Split(rawTypes, ","))
            currentSpeciesInfo.Types = copiedTypes
        } 
    }
    if currentSpeciesInfo.Species != "" {
        speciesInfos = append(speciesInfos, currentSpeciesInfo) 
    }
    return speciesInfos
}
