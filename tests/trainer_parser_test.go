package parsers_test

import (
    "testing"
    "reflect"
    "github.com/geefuoco/trainer_editor/parsers"
    "github.com/geefuoco/trainer_editor/data_objects"
)

func TestParseTrainers(t *testing.T) {
    // Create a test input string that simulates the contents of a file
    input := "test_cases/trainer_testcase.txt"    

    // Call the function with the test input
    trainers := parsers.ParseTrainers(input)

    // Check the result
    if len(trainers) != 1 {
        t.Errorf("Expected 1 trainer, but got %d", len(trainers))
        return
    }

    expectedTrainer := &data_objects.Trainer{
        TrainerKey:           "TRAINER_SAWYER_1",
        TrainerClass:         "TRAINER_CLASS_HIKER",
        EncounterMusicGender: "TRAINER_ENCOUNTER_MUSIC_MALE",
        TrainerPic:           "TRAINER_PIC_HIKER",
        TrainerName:          "SAWYER",
        Items:                [4]string{"ITEM_NONE","ITEM_NONE", "ITEM_NONE", "ITEM_NONE"},
        DoubleBattle:         true,
        AiFlags:              []string{"AI_FLAG_A", "AI_FLAG_B"},
        Party:                "TRAINER_PARTY(sParty_Sawyer1)",
    }


    if !(reflect.DeepEqual(trainers[0], expectedTrainer)) {
        t.Log("Expected\n")
        t.Log(expectedTrainer)
        t.Log("Actual\n")
        t.Log(trainers[0].String())
        t.Error("Trainers were not Equal")
    }
}

func TestParseNoneTrainer(t *testing.T) {
    // Create a test input string that simulates the contents of a file
    input := "test_cases/notrainer_testcase.txt"    

    // Call the function with the test input
    trainers := parsers.ParseTrainers(input)

    // Check the result
    if len(trainers) != 1 {
        t.Errorf("Expected 1 trainer, but got %d", len(trainers))
        return
    }

    expectedTrainer := &data_objects.Trainer{
        TrainerKey:           "TRAINER_NONE",
        TrainerClass:         "TRAINER_CLASS_PKMN_TRAINER_1",
        EncounterMusicGender: "TRAINER_ENCOUNTER_MUSIC_MALE",
        TrainerPic:           "TRAINER_PIC_HIKER",
        TrainerName:          "",
        Items:                [4]string{"ITEM_NONE","ITEM_NONE","ITEM_NONE","ITEM_NONE"},
        DoubleBattle:         false,
        AiFlags:              []string{},
        Party:                "",
    }

    if !(reflect.DeepEqual(trainers[0], expectedTrainer)) {
        t.Error("Trainers were not Equal")
    }

}
