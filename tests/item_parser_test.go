package parser_tests

import (
    "testing"
    "reflect"
    "github.com/geefuoco/trainer_editor/parsers"
)

func TestItemParser(t *testing.T) {
    input := "test_cases/items_testcase.txt"


    expectedItems := []string{
        "ITEM_NONE",
        "ITEM_POKE_BALL",
        "ITEM_GREAT_BALL",
        "ITEM_ULTRA_BALL",
        "ITEM_MASTER_BALL",
        "ITEM_PREMIER_BALL",
        "ITEM_HEAL_BALL",
        "ITEM_NET_BALL",
        "ITEM_NEST_BALL",
        "ITEM_DIVE_BALL",
        "ITEM_DUSK_BALL",
        "ITEM_TIMER_BALL",
        "ITEM_QUICK_BALL",
        "ITEM_REPEAT_BALL",
        "ITEM_LUXURY_BALL",
        "ITEM_LEVEL_BALL",
    }

    items := parsers.ParseItems(input)

    t.Logf("Actual Items:\n")
    for _, item := range(items) {
        t.Logf("%s\n", item)
    }
    
    if !(reflect.DeepEqual(expectedItems, items)) {
        t.Errorf("Items did not match")
    }

}
