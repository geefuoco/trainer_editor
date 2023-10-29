package data_objects_tests

import (
    "testing"
    "github.com/geefuoco/trainer_editor/parsers"
)

func TestTrainerMonString(t *testing.T) {

    input := "test_cases/trainer_party_multiple_testcase.txt"

    parties := parsers.ParseTrainerParties(input)

    expectedString := `static const struct TrainerMon sParty_Roxanne1[] = {
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

    if parties[0].String() != expectedString {
        t.Errorf("Strings were not equal")
        t.Log("Actual")
        t.Log(parties[0].String())
        t.Log("Expected")
        t.Log(expectedString)
    }
}
