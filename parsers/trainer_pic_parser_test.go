package parsers

import (
    "testing"
    "reflect"
)

func TestParseTrainerPics(t *testing.T) {
    t.Parallel()

    input1 := "testdata/trainer_pic_testcase.txt"
    input2 := "testdata/trainer_frontpic_testcase.txt"

    pics := ParseTrainerPics(input1, input2)

    expectedTrainerPics := map[string]string{
        "TRAINER_PIC_HIKER": "graphics/trainers/front_pics/hiker.png",
        "TRAINER_PIC_AQUA_GRUNT_M": "graphics/trainers/front_pics/aqua_grunt_m.png",
        "TRAINER_PIC_POKEMON_BREEDER_F": "graphics/trainers/front_pics/pokemon_breeder_f.png",
    }

    if !reflect.DeepEqual(pics, expectedTrainerPics) {
        t.Errorf("Trainer Pics were not equal")
    }
}
