package parser_tests

import (
    "testing"
    "reflect"
    "github.com/geefuoco/trainer_editor/parsers"
)

func TestParseTrainerPics(t *testing.T) {

    input1 := "test_cases/trainer_pic_testcase.txt"
    input2 := "test_cases/trainer_frontpic_testcase.txt"

    pics := parsers.ParseTrainerPics(input1, input2)

    expectedTrainerPics := map[string]string{
        "TRAINER_PIC_HIKER": "graphics/trainers/front_pics/hiker.png",
        "TRAINER_PIC_AQUA_GRUNT_M": "graphics/trainers/front_pics/aqua_grunt_m.png",
        "TRAINER_PIC_POKEMON_BREEDER_F": "graphics/trainers/front_pics/pokemon_breeder_f.png",
    }

    if !reflect.DeepEqual(pics, expectedTrainerPics) {
        t.Errorf("Trainer Pics were not equal")
    }
}
