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
    if len(parties) != 2 {
        t.Errorf("Expected 2 trainerMon, but got %d", len(parties))
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

    if !reflect.DeepEqual(expectedParty, parties[0]) {
        t.Errorf("TrainerParty was not equal")
    }

    testMons = []*data_objects.TrainerMon {
        {
        Iv : [6]uint64{31, 31, 31, 31, 31, 31},
        Lvl : 32,
        Species : "SPECIES_GOLEM",
        HeldItem : "ITEM_NONE",
        Moves : [4]string{"MOVE_PROTECT", "MOVE_ROLLOUT", "MOVE_MAGNITUDE", "MOVE_EXPLOSION"},
        },
        {
        Iv :[6]uint64{31, 31, 31, 31, 31, 31},
        Lvl : 35,
        Species : "SPECIES_KABUTO",
        HeldItem : "ITEM_SITRUS_BERRY",
        Moves : [4]string{"MOVE_SWORDS_DANCE", "MOVE_ICE_BEAM", "MOVE_SURF", "MOVE_ROCK_SLIDE"},
        },
        {
        Iv : [6]uint64{31, 31, 31, 31, 31, 31},
        Lvl : 35,
        Species : "SPECIES_ONIX",
        HeldItem : "ITEM_NONE",
        Moves : [4]string{"MOVE_IRON_TAIL", "MOVE_EXPLOSION", "MOVE_ROAR", "MOVE_ROCK_SLIDE"},
        },
        {
        Iv : [6]uint64{31, 31, 31, 31, 31, 31},
        Lvl : 37,
        Species : "SPECIES_NOSEPASS",
        HeldItem : "ITEM_SITRUS_BERRY",
        Moves : [4]string{"MOVE_DOUBLE_TEAM", "MOVE_EXPLOSION", "MOVE_PROTECT", "MOVE_ROCK_SLIDE"},
        },
    }

    expectedParty = &data_objects.TrainerParty {
        Trainer: "sParty_Roxanne2",
        Party: testMons,
    }

    if !reflect.DeepEqual(parties[1], expectedParty) {
        t.Log("Expected")
        t.Log(expectedParty.String())
        t.Log("Actual")
        t.Log(parties[1].String())
        t.Errorf("Trainer Party was not equal\n")
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

