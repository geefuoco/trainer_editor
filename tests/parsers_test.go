package parsers_test

import (
    "testing"
    "github.com/geefuoco/trainer_editor/parsers"
    "github.com/geefuoco/trainer_editor/data_objects"
)

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
        Items:                []string{},
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
        Items:                []string{},
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
