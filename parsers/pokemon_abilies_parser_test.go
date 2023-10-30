package parsers

import (
    "reflect"
    "testing"
)

func TestPokemonAbilitiesParser(t *testing.T) {
    t.Parallel()
    input := "testdata/abilities_testcase.txt"

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

    actual := ParsePokemonAbilities(input) 

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
