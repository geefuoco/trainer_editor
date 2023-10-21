package parsers_test

import (
    "testing"
    "reflect"
    "github.com/geefuoco/trainer_editor/parsers"
    "github.com/geefuoco/trainer_editor/data_objects"
)

func TestParseTrainerPartyMultipleMons(t *testing.T) {
    // Create a test input string that simulates the contents of a file
    input := "test_cases/trainer_party_multiple_testcase.txt"    

    // Call the function with the test input
    parties := parsers.ParseTrainerParties(input)

    // Check the result
    if len(parties) != 1 {
        t.Errorf("Expected 1 trainerMon, but got %d", len(parties))
        return
    }

    testMons := []*data_objects.TrainerMon{
        {
            Iv: [6]uint64{7, 0, 4, 5, 30, 20},
            Lvl: 18,
            Species: "SPECIES_HONEDGE",
            HeldItem: "ITEM_COLBUR_BERRY",
            Moves: [4]string{"MOVE_AUTOTOMIZE", "MOVE_SHADOW_SNEAK", "MOVE_METAL_SOUND", "MOVE_AERIAL_ACE"},
            IsShiny: false,
        },
        {
            Lvl: 18,
            Species: "SPECIES_SABLEYE",
            HeldItem: "ITEM_NONE",
            Moves: [4]string{"MOVE_THUNDER_WAVE", "MOVE_FAKE_OUT", "MOVE_SHADOW_SNEAK", "MOVE_AERIAL_ACE"},
            Ability: "ABILITY_PRANKSTER",
            IsShiny: false,
        },
        {
            Lvl: 18,
            Species: "SPECIES_MISDREAVUS",
            IsShiny: true,
        },
    }

    expectedParty := &data_objects.TrainerParty{
        Trainer: "sParty_Roxanne1",
        Party: testMons,
    }

    if !(parties[0].Trainer == expectedParty.Trainer) {
        t.Errorf("Trainer was not equal")
    }

    for i := range(expectedParty.Party) {
        expectedPartyMon := expectedParty.Party[i]
        actualPartyMon := parties[0].Party[i]
        if !(reflect.DeepEqual(expectedPartyMon, actualPartyMon)) {
            t.Errorf("Party Mons did not match")
        }
    }

}

func TestParseTrainerParty(t *testing.T) {
    // Create a test input string that simulates the contents of a file
    input := "test_cases/trainer_party_testcase.txt"    

    // Call the function with the test input
    parties := parsers.ParseTrainerParties(input)

    // Check the result
    if len(parties) != 1 {
        t.Errorf("Expected 1 trainerMon, but got %d", len(parties))
        return
    }

    testMon := &data_objects.TrainerMon{
            Iv: [6]uint64{1, 2, 3, 4, 5, 6},
            Lvl: 25,
            Species: "SPECIES_MEWTWO",
            Ev: [6]uint64{10, 11, 21, 24, 90, 0},
            HeldItem: "ITEM_NONE",
            Moves: [4]string{"MOVE_SLASH", "MOVE_PSYCHIC", "MOVE_THUNDER", "MOVE_BLIZZARD"},
            Ability: "ABILITY_INTIMIDATE",
            IsShiny: false,
    }

    expectedParty := &data_objects.TrainerParty{
        Trainer: "sParty_Sawyer1",
        Party: []*data_objects.TrainerMon{
            testMon,
        },
    }

    if !(reflect.DeepEqual(parties[0], expectedParty)) {
        t.Errorf("Party Mons did not match")
    }
}

