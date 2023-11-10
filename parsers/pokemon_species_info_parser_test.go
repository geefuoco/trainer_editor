package parsers

import (
    "testing"
    "reflect"
    "github.com/geefuoco/trainer_editor/data_objects"
    "github.com/geefuoco/trainer_editor/logging"
    "os"
)

func TestParsePokemonSpeciesInfo(t *testing.T) {
    t.Parallel()
    logging.EnableLogging()
    input := "testdata/pokemon_species_info_testcase.txt"

    f, err := os.ReadFile(input)
    if err != nil {
        t.Fatalf("Could not read file: "+input)
        return
    }
    fileContents := string(f)

    expected := []*data_objects.PokemonSpeciesInfo{
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
        {
            Species: "SPECIES_BULBASAUR",
            BaseHp: 45,
            BaseAtk: 49,
            BaseDef: 49,
            BaseSpd: 45,
            BaseSpAtk: 65,
            BaseSpDef: 65,
            Types: [2]string{"TYPE_GRASS", "TYPE_POISON"},
        },
        {
            Species: "SPECIES_IVYSAUR",
            BaseHp: 60,
            BaseAtk: 62,
            BaseDef: 63,
            BaseSpd: 60,
            BaseSpAtk: 80,
            BaseSpDef: 80,
            Types: [2]string{"TYPE_GRASS", "TYPE_POISON"},
        },
        {
            Species: "SPECIES_VENUSAUR",
            BaseHp: 80,
            BaseAtk: 82,
            BaseDef: 83,
            BaseSpd: 80,
            BaseSpAtk: 100,
            BaseSpDef: 100,
            Types: [2]string{"TYPE_GRASS", "TYPE_POISON"},
        },
        {
            Species: "SPECIES_CHARMANDER",
            BaseHp: 39,
            BaseAtk: 52,
            BaseDef: 43,
            BaseSpd: 65,
            BaseSpAtk: 60,
            BaseSpDef: 50,
            Types: [2]string{"TYPE_FIRE", "TYPE_FIRE"},
        },
    }

    actual := ParsePokemonSpeciesInfo(fileContents)

    missingSpeciesInfo := []string{}
    for i:=0; i < len(expected); i++ {
        expectedMon := expected[i]
        currentMon := getMon(actual, expectedMon.Species)
        if currentMon == nil {
            missingSpeciesInfo = append(missingSpeciesInfo, expectedMon.Species)
        } else {
            if !reflect.DeepEqual(currentMon, expectedMon) {
                t.Errorf("%s was not equal", currentMon.Species)
            }
        }
    }
    if len(actual) != len(expected) {
        t.Errorf("Length mismatch %d vs expected %d", len(actual), len(expected))
        t.Log("Missing: ")
        for _, sp := range actual {
            t.Log(sp)
        }
        return
    }

}

func containsMon(slice []*data_objects.PokemonSpeciesInfo, species string) bool {
    for _, x:= range slice {
        if x.Species == species {
            return true
        }
    }
    return false
}

func getMon(slice []*data_objects.PokemonSpeciesInfo, species string) *data_objects.PokemonSpeciesInfo{
    for _, x:= range slice {
        if x.Species == species {
            return x
        }
    }
    return nil
}
