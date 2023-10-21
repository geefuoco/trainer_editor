package parsers_test

import (
    "testing"
    "reflect"
    "github.com/geefuoco/trainer_editor/parsers"
    "github.com/geefuoco/trainer_editor/data_objects"
)

func TestParseTrainerPartyMultipleMons(t *testing.T) {
    // Create a test input string that simulates the contents of a file
    input := "trainer_party_multiple_testcase.txt"    

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
    input := "trainer_party_testcase.txt"    

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

func TestParseTrainers(t *testing.T) {
    // Create a test input string that simulates the contents of a file
    input := "trainer_testcase.txt"    

    // Call the function with the test input
    trainers := parsers.ParseTrainers(input)

    // Check the result
    if len(trainers) != 1 {
        t.Errorf("Expected 1 trainer, but got %d", len(trainers))
        return
    }

    expectedTrainer := data_objects.Trainer{
        TrainerClass:         "TRAINER_CLASS_PKMN_TRAINER_1",
        EncounterMusicGender: "TRAINER_ENCOUNTER_MUSIC_MALE",
        TrainerPic:           "TRAINER_PIC_HIKER",
        TrainerName:          "Trainer1",
        Items:                [4]string{"ITEM_NONE","ITEM_NONE", "ITEM_NONE", "ITEM_NONE"},
        DoubleBattle:         true,
        AiFlags:              []string{"AI_FLAG_A", "AI_FLAG_B"},
        Party:                "TRAINER_PARTY(myparty)",
    }

    if  !(trainers[0].TrainerClass == expectedTrainer.TrainerClass) {
        t.Errorf("TrainerClass was not equal")
    }
    if  !(trainers[0].EncounterMusicGender == expectedTrainer.EncounterMusicGender) {
        t.Errorf("EncounterMusicGender was not equal")
    }
    if  !(trainers[0].TrainerPic == expectedTrainer.TrainerPic) {
        t.Errorf("TrainerPic was not equal")
    }
    if  !(trainers[0].TrainerName == expectedTrainer.TrainerName) {
        t.Errorf("TrainerName was not equal")
    }
    if  !(len(trainers[0].Items) == len(expectedTrainer.Items)) {
        t.Errorf("Item Length was not equal")
    }
    for i := range(trainers[0].Items) {
        if  !(trainers[0].Items[i] == expectedTrainer.Items[i]) {
            t.Errorf("Item %d was not equal", i)
        }
    }
    if  !(trainers[0].DoubleBattle == expectedTrainer.DoubleBattle) {
        t.Errorf("DoubleBattle was not equal")
    }
    if  !(len(trainers[0].AiFlags) == len(expectedTrainer.AiFlags)) {
        t.Errorf("AiFlags length was not equal")
    }
    for i := range(trainers[0].AiFlags) {
        if  !(trainers[0].AiFlags[i] == expectedTrainer.AiFlags[i]) {
            t.Errorf("AiFlag %d was not equal", i)
        }
    }
    if  !(trainers[0].Party == expectedTrainer.Party) {
        t.Errorf("Party was not equal")
    }
}

func TestParseNoneTrainer(t *testing.T) {
    // Create a test input string that simulates the contents of a file
    input := "notrainer_testcase.txt"    

    // Call the function with the test input
    trainers := parsers.ParseTrainers(input)

    // Check the result
    if len(trainers) != 1 {
        t.Errorf("Expected 1 trainer, but got %d", len(trainers))
        return
    }

    expectedTrainer := data_objects.Trainer{
        TrainerClass:         "TRAINER_CLASS_PKMN_TRAINER_1",
        EncounterMusicGender: "TRAINER_ENCOUNTER_MUSIC_MALE",
        TrainerPic:           "TRAINER_PIC_HIKER",
        TrainerName:          "",
        Items:                [4]string{"ITEM_NONE","ITEM_NONE","ITEM_NONE","ITEM_NONE"},
        DoubleBattle:         false,
        AiFlags:              []string{},
        Party:                "",
    }

    if  !(trainers[0].TrainerClass == expectedTrainer.TrainerClass) {
        t.Logf(trainers[0].TrainerClass + "!=" + expectedTrainer.TrainerClass)
        t.Errorf("TrainerClass was not equal")
    }
    if  !(trainers[0].EncounterMusicGender == expectedTrainer.EncounterMusicGender) {
        t.Logf(trainers[0].EncounterMusicGender + "!=" + expectedTrainer.EncounterMusicGender)
        t.Errorf("EncounterMusicGender was not equal")
    }
    if  !(trainers[0].TrainerPic == expectedTrainer.TrainerPic) {
        t.Logf(trainers[0].TrainerPic + "!=" + expectedTrainer.TrainerPic)
        t.Errorf("TrainerPic was not equal")
    }
    if  !(trainers[0].TrainerName == expectedTrainer.TrainerName) {
        t.Logf(trainers[0].TrainerName + "!=" + expectedTrainer.TrainerName)
        t.Errorf("TrainerName was not equal")
    }
    if  !(len(trainers[0].Items) == len(expectedTrainer.Items)) {
        t.Logf("%d != %d", len(trainers[0].Items) , len(expectedTrainer.Items))
        t.Errorf("Item Length was not equal")
    }
    for i := range(trainers[0].Items) {
        if  !(trainers[0].Items[i] == expectedTrainer.Items[i]) {
            t.Logf(trainers[0].Items[i] + "!=" + expectedTrainer.Items[i])
            t.Errorf("Item %d was not equal", i)
        }
    }
    if  !(trainers[0].DoubleBattle == expectedTrainer.DoubleBattle) {
        t.Logf("%t != %t", trainers[0].DoubleBattle , expectedTrainer.DoubleBattle)
        t.Errorf("DoubleBattle was not equal")
    }
    if  !(len(trainers[0].AiFlags) == len(expectedTrainer.AiFlags)) {
        t.Logf("%d != %d", len(trainers[0].AiFlags) ,len(expectedTrainer.AiFlags))
        t.Errorf("AiFlags length was not equal")
    }
    for i := range(trainers[0].AiFlags) {
        if  !(trainers[0].AiFlags[i] == expectedTrainer.AiFlags[i]) {
            t.Logf(trainers[0].AiFlags[i] + "!=" + expectedTrainer.AiFlags[i])
            t.Errorf("AiFlag %d was not equal", i)
        }
    }
    if  !(trainers[0].Party == expectedTrainer.Party) {
        t.Logf(trainers[0].Party + "!=" + expectedTrainer.Party)
        t.Errorf("Party was not equal")
    }
}
