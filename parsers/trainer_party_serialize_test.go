package parsers

import (
    "testing"
)

func TestTrainerMonString(t *testing.T) {
    t.Parallel()

    input := "testdata/trainer_party_multiple_testcase.txt"

    actual := ParseTrainerParties(input)[0]

    expected := `static const struct TrainerMon sParty_Roxanne1[] = {
    {
    .iv = TRAINER_PARTY_IVS(7, 0, 4, 5, 30, 20),
    .lvl = 18,
    .species = SPECIES_HONEDGE,
    .heldItem = ITEM_COLBUR_BERRY,
    .moves = {MOVE_AUTOTOMIZE, MOVE_SHADOW_SNEAK, MOVE_METAL_SOUND, MOVE_AERIAL_ACE},
    },
    {
    .lvl = 18,
    .species = SPECIES_SABLEYE,
    .moves = {MOVE_THUNDER_WAVE, MOVE_FAKE_OUT, MOVE_SHADOW_SNEAK, MOVE_AERIAL_ACE},
    .ability = ABILITY_PRANKSTER,
    },
    {
    .lvl = 18,
    .species = SPECIES_MISDREAVUS,
    .isShiny = TRUE,
    }
};
`

    if actual.String() != expected {
        t.Errorf("Strings were not equal")
        t.Log("Actual")
        t.Log(actual.String())
        t.Log("Expected")
        t.Log(expected)
        for i:=0; i < len(expected); i++ {
            t.Logf("%d: %c | %c", i, expected[i], actual.String()[i])
            if expected[i] != actual.String()[i] {
                t.Log("Did not Match here ^")
            }
        }
    }
}
