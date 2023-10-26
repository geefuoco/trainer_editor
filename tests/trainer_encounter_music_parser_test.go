package parser_tests

import (
    "reflect"
    "testing"
    "github.com/geefuoco/trainer_editor/parsers"
)

func TestParseEncounterMusic(t *testing.T) {

    input := "test_cases/trainer_encounter_music_testcase.txt"

    expectedEncounterMusics := []string{
        "TRAINER_ENCOUNTER_MUSIC_MALE",
        "TRAINER_ENCOUNTER_MUSIC_FEMALE",
        "TRAINER_ENCOUNTER_MUSIC_GIRL",
        "TRAINER_ENCOUNTER_MUSIC_SUSPICIOUS",
        "TRAINER_ENCOUNTER_MUSIC_INTENSE",
        "TRAINER_ENCOUNTER_MUSIC_COOL",
        "TRAINER_ENCOUNTER_MUSIC_AQUA",
        "TRAINER_ENCOUNTER_MUSIC_MAGMA",
        "TRAINER_ENCOUNTER_MUSIC_SWIMMER",
        "TRAINER_ENCOUNTER_MUSIC_TWINS",
        "TRAINER_ENCOUNTER_MUSIC_ELITE_FOUR",
        "TRAINER_ENCOUNTER_MUSIC_HIKER",
        "TRAINER_ENCOUNTER_MUSIC_INTERVIEWER",
        "TRAINER_ENCOUNTER_MUSIC_RICH",
    }

    music := parsers.ParseTrainerEncounterMusic(input)

    if !reflect.DeepEqual(music, expectedEncounterMusics) {
        t.Errorf("Enounter music was not equal")
    }
}
