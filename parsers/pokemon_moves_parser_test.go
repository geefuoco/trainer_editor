package parsers

import (
    "reflect"
    "testing"
)

func TestPokemonMovesParser(t *testing.T) {
    t.Parallel()

    input := "testdata/pokemon_moves_testcase.txt"

    expectedOutput := []string {
        "MOVE_NONE",
        "MOVE_POUND",
        "MOVE_KARATE_CHOP",
        "MOVE_DOUBLE_SLAP",
        "MOVE_COMET_PUNCH",
        "MOVE_MEGA_PUNCH",
        "MOVE_PAY_DAY",
        "MOVE_FIRE_PUNCH",
        "MOVE_ICE_PUNCH",
        "MOVE_THUNDER_PUNCH",
        "MOVE_SCRATCH",
        "MOVE_VISE_GRIP",
        "MOVE_GUILLOTINE",
        "MOVE_RAZOR_WIND",
        "MOVE_SWORDS_DANCE",
        "MOVE_CUT",
        "MOVE_GUST",
        "MOVE_WING_ATTACK",
        "MOVE_WHIRLWIND",
        "MOVE_FLY",
        "MOVE_BIND",
        "MOVE_MATCHA_GOTCHA",
        "MOVE_SYRUP_BOMB",
        "MOVE_IVY_CUDGEL",
    }

    actual := ParseMoves(input)

    if !reflect.DeepEqual(actual, expectedOutput) {
        t.Logf("Expected\n")
        for _, k := range expectedOutput {
            t.Logf(k+"\n")
        }
        t.Logf("Actual\n")
        for _, k := range actual {
            t.Logf(k+"\n")
        }
        t.Errorf("Moves were not equal\n")
    }

}
