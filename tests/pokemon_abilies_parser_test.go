package parser_tests

import (
    "reflect"
    "testing"
    "github.com/geefuoco/trainer_editor/parsers"
)

func TestPokemonAbilitiesParser(t *testing.T) {
    input := "test_cases/abilities_testcase.txt"

    expectedOutput := []string {
        "ABILITY_NONE",
        "ABILITY_STENCH",
        "ABILITY_DRIZZLE",
        "ABILITY_SPEED_BOOST",
        "ABILITY_BATTLE_ARMOR",
        "ABILITY_STURDY",
        "ABILITY_DAMP",
        "ABILITY_LIMBER",
        "ABILITY_SAND_VEIL",
        "ABILITY_STATIC",
        "ABILITY_VOLT_ABSORB",
        "ABILITY_WATER_ABSORB",
        "ABILITY_OBLIVIOUS",
    }

    actual := parsers.ParsePokemonAbilities(input) 

    if !reflect.DeepEqual(actual, expectedOutput) {
        t.Log("Expected\n")
        for _, v := range expectedOutput {
            t.Log(v+"\n")
        }
        t.Log("Actual\n")
        for _, v := range actual {
            t.Log(v+"\n")
        }
        t.Errorf("Abilities were not equal\n")
    }
}
