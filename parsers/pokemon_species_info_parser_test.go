package parsers

import (
    "testing"
    "reflect"
    "github.com/geefuoco/trainer_editor/data_objects"
)

func TestParsePokemonSpeciesInfo(t *testing.T) {

    input := "testdata/pokemon_species_info_testcase.txt"

    expected := []*data_objects.SpeciesInfo{
        {
            BaseHp: 35,
            BaseAttack: 55,
            BaseSpeed: 90,
            BaseSpAttack: 50,
            BaseDefence: 40,
            BaseSpDefence: 50,
            Types: [2]string{"TYPE_ELECTRIC", "TYPE_ELECTRIC"}
        },
    }

    actual := parsers.ParsePokemonSpeciesInfo(input)

    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("Species Infos were not equal")
    }
}
