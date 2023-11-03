package parsers

import (
    "github.com/geefuoco/trainer_editor/logging"
    "github.com/geefuoco/trainer_editor/data_objects"
    "os"
    "strings"
    "strconv"
)

func ParsePokemonSpeciesInfo(filepath string) []*data_objects.PokemonSpeciesInfo {
    file, err := os.ReadFile(filepath)
    if err != nil {
        logging.ErrorLog("Could not open file: "+filepath)
        return nil
    }


    fileContents := string(file)
    // Handle Edge cases with Minior by putting them in the slice, and skipping them during parsing 
    speciesInfos := []*data_objects.PokemonSpeciesInfo {
        {
            Species: "SPECIES_MINIOR",
            BaseHp: 60,
            BaseAtk: 60,
            BaseDef: 100,
            BaseSpd: 60,
            BaseSpAtk: 60,
            BaseSpDef: 100,
            Types: [2]string{"TYPE_ROCK", "TYPE_FLYING"},
        },
        {
            Species: "SPECIES_MINIOR_CORE",
            BaseHp: 60,
            BaseAtk: 100,
            BaseDef: 60,
            BaseSpd: 120,
            BaseSpAtk: 100,
            BaseSpDef: 60,
            Types: [2]string{"TYPE_ROCK", "TYPE_FLYING"},
        },
    }
    var speciesSet = make(map[string]bool)
    var skip bool
    currentSpeciesInfo := &data_objects.PokemonSpeciesInfo{}

    for _, line := range strings.Split(fileContents, "\n") {
        // logic for skipping species info if there is an error in parsing
        if skip && strings.Contains(line, ".base") {
            continue
        }
        // Extra Edge Cases for Pikachu
        if strings.Contains(line, "COSPLAY") || strings.Contains(line,"CAP_PIKACHU") {
            continue
        }

        // Edge Case for MINIOR

        if strings.Contains(line, "MINIOR") {
            continue
        }

        if strings.Contains(line, "SPECIES_FLAG") {
            continue
        }


        // This is a macro TBD how to parse it
        if strings.Contains(line, "SPECIES_INFO") {
            skip=false
            if strings.Contains(line, "#define") {
                // First make sure that the previous species is stored in the array
                if currentSpeciesInfo.Species != "" {
                    speciesInfos = append(speciesInfos, currentSpeciesInfo)
                    currentSpeciesInfo = &data_objects.PokemonSpeciesInfo{}
                }
                // Split the string
                split := strings.Split(line, " ")
                // first string is #define
                // Second string is XXXXX_SPECIES_INFO(arg)
                if len(split) < 2 {
                    logging.WarnLog("Could not parse macro: "+line)
                    continue
                }
                rawSpeciesInfo:= split[1]
                end := strings.Index(rawSpeciesInfo, "_SPECIES_INFO")
                if end == -1 {
                    logging.WarnLog("Could not parse line: " + line)
                    continue
                }
                species := rawSpeciesInfo[:end]
                speciesName := "SPECIES_"+species
                _, has := speciesSet[speciesName]
                if !has {
                    speciesSet[speciesName]=true
                    currentSpeciesInfo.Species = speciesName
                }
            } else {
                // First make sure that the previous species is stored in the array
                if currentSpeciesInfo.Species != "" {
                    speciesInfos = append(speciesInfos, currentSpeciesInfo)
                    currentSpeciesInfo = &data_objects.PokemonSpeciesInfo{}
                }


                // Macro is being using to assign to a species
                start := strings.IndexByte(line, '=')+1
                end := strings.Index(line, "_SPECIES_INFO")
                if start == -1 || end == -1 || start>end {
                    logging.WarnLog("Error when parsing macro assignment: "+line)
                    continue
                }
                species := strings.TrimSpace(line[start:end])
                speciesName := "SPECIES_"+species

                start =strings.IndexByte(line, '[')+1
                end = strings.IndexByte(line, ']')
                if start == -1 || end == -1 || start>end {
                    logging.WarnLog("Could not parse line: "+line)
                    continue
                }
                // i.e. SPECIES_SAWSBUCK_AUTUMN
                actualSpeciesName := line[start:end]
                // The edge case with certain macros
                if speciesName == actualSpeciesName {
                    continue
                }

                // Find the SPECIES_XXXX in the list that matches
                // The species info being used
                targetSpeciesInfo := &data_objects.PokemonSpeciesInfo{}
                for _, mon := range speciesInfos {
                    if strings.Contains(speciesName, mon.Species) {
                        *targetSpeciesInfo = *mon
                        targetSpeciesInfo.Species = actualSpeciesName
                        break
                    }
                }
                if targetSpeciesInfo.Species == ""  {
                    logging.ErrorLog("Could not find species: %s related to species info macro: %s", speciesName, actualSpeciesName)
                } else {
                    // Edge Cases for Types being assigned in the Macro
                    if strings.Contains(line, "TYPE_") {
                        start := strings.Index(line, "TYPE_")
                        var end int
                        if strings.Contains(line, "ARCEUS") || strings.Contains(line, "SILVALLY"){
                            end = strings.IndexByte(line, ')')
                        } else {
                            end = strings.IndexByte(line, ',') 
                        }
                        typing := line[start:end]

                        if strings.Contains(line, "ROTOM") {
                            targetSpeciesInfo.Types[0] = "TYPE_ELECTRIC"
                            targetSpeciesInfo.Types[1] = typing
                        } else if strings.Contains(line, "ORICORIO"){
                            targetSpeciesInfo.Types[0] = typing
                            targetSpeciesInfo.Types[1] = "TYPE_FLYING"
                        } else {
                            targetSpeciesInfo.Types[0] = typing
                            targetSpeciesInfo.Types[1] = typing
                        }
                    }
                    speciesSet[actualSpeciesName]=true
                    speciesInfos = append(speciesInfos, targetSpeciesInfo)
                }
            }
        } else if strings.Contains(line, "SPECIES") {
            skip=false
            // Skip the none species
            if strings.Contains(line, "SPECIES_NONE") {
                continue
            }
            start :=strings.IndexByte(line, '[')+1
            end := strings.IndexByte(line, ']')
            if start == -1 || end == -1 || start > end{
                logging.WarnLog("Could not parse line: "+line)
                continue
            }

            if currentSpeciesInfo.Species != "" {
                speciesInfos = append(speciesInfos, currentSpeciesInfo) 
            } 
            speciesName := line[start:end]
            _, has := speciesSet[speciesName]
            if !has {
                speciesSet[speciesName]=true
                currentSpeciesInfo = &data_objects.PokemonSpeciesInfo{
                    Species: speciesName,
                }
            }
        } else if strings.Contains(line, ".base") {
            // Don't parse the stats in the Macro
            if currentSpeciesInfo.Species == "" {
                continue
            }
            start := strings.IndexByte(line, '=')
            end := strings.IndexByte(line, ',')
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
        } else if strings.Contains(line, "PIKACHU_BASE_DEFENSES") {
            // Edge case with Pikachu's stats
            if currentSpeciesInfo.Species == "SPECIES_PIKACHU" {
                currentSpeciesInfo.BaseDef = 40
                currentSpeciesInfo.BaseSpDef = 50
            }
        } else if strings.Contains(line, "BUTTERFREE_SP_ATK") {
            if currentSpeciesInfo.Species == "SPECIES_BUTTERFREE" {
                currentSpeciesInfo.BaseSpAtk = 90
            }
        } else if strings.Contains(currentSpeciesInfo.Species, "PUMPKABOO") || strings.Contains(currentSpeciesInfo.Species, "GOURGEIST") {
        // Special Case for specific pokemon because of how the code is layed out
            currentSpeciesInfo.Types = [2]string{"TYPE_GHOST", "TYPE_GRASS"}
        } else if strings.Contains(currentSpeciesInfo.Species, "CASTFORM") {
            currentSpeciesInfo.Types = [2]string{"TYPE_NORMAL", "TYPE_NORMAL"}
        } else if strings.Contains(currentSpeciesInfo.Species, "SILVALLY") {
            currentSpeciesInfo.Types = [2]string{"TYPE_NORMAL", "TYPE_NORMAL"}
        } else if strings.Contains(currentSpeciesInfo.Species, "ARCEUS") {
            currentSpeciesInfo.Types = [2]string{"TYPE_NORMAL", "TYPE_NORMAL"}
        } else if strings.Contains(currentSpeciesInfo.Species, "ORICORIO") {
            currentSpeciesInfo.Types = [2]string{"TYPE_FIRE", "TYPE_FLYING"}
        }

    }
    if currentSpeciesInfo.Species != "" {
        speciesInfos = append(speciesInfos, currentSpeciesInfo) 
    }
    // Remove the ROTOM_FORM species info
    speciesInfos = removeFromSlice(speciesInfos, "SPECIES_ROTOM_FORM")

    return speciesInfos
}

func removeFromSlice(slice []*data_objects.PokemonSpeciesInfo, value string)[]*data_objects.PokemonSpeciesInfo {
    for i := range slice {
        if slice[i].Species == value {
            newSlice := append(slice[:i-1], slice[i+1:]...)
            return newSlice
        }
    }
    return slice
}
