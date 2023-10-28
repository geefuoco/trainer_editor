package parsers_test

import (
    "reflect"
    "testing"
    "github.com/geefuoco/trainer_editor/parsers"
)

func TestPokemonPicsParser(t *testing.T) {

    input1 := "test_cases/pokemon_species_testcase.txt"
    input2 := "test_cases/pokemon_pic_parser_testcase.txt"


    expectedOutput := map[string]string {
        "SPECIES_BULBASAUR": "graphics/pokemon/bulbasaur/anim_front.png",
        "SPECIES_IVYSAUR":"graphics/pokemon/ivysaur/anim_front.png",
        "SPECIES_VENUSAUR":"graphics/pokemon/venusaur/anim_front.png",
        "SPECIES_CHARMANDER":"graphics/pokemon/charmander/anim_front.png",
        "SPECIES_CHARMELEON":"graphics/pokemon/charmeleon/anim_front.png",
        "SPECIES_CHARIZARD":"graphics/pokemon/charizard/anim_front.png",
        "SPECIES_SQUIRTLE":"graphics/pokemon/squirtle/anim_front.png",
        "SPECIES_ROSELIA":"graphics/pokemon/roselia/anim_front.png",
    }

    actual := parsers.ParsePokemonPics(input1, input2)

    if !reflect.DeepEqual(actual, expectedOutput) {
        t.Logf("Actual\n")
        for k, v := range actual {
            t.Logf(k + "->" + v +"\n")
        }
        t.Logf("Expected\n")
        for k, v := range expectedOutput{
            t.Logf(k + "->" + v +"\n")
        }
        t.Errorf("Maps were not equal")
    }

}
