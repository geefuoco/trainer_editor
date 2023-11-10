package parsers

import (
    "reflect"
    "testing"
)

func TestAiFlagParser(t *testing.T) {
    t.Parallel()

    input := "testdata/aiflags_testcase.txt"

    expectedAiFlags := []string {
        "AI_FLAG_BASIC_GOOD_TRAINER",
        "AI_FLAG_CHECK_BAD_MOVE",
        "AI_FLAG_TRY_TO_FAINT",
        "AI_FLAG_CHECK_VIABILITY",
        "AI_FLAG_SETUP_FIRST_TURN",
        "AI_FLAG_RISKY",
        "AI_FLAG_PREFER_STRONGEST_MOVE",
        "AI_FLAG_PREFER_BATON_PASS",
        "AI_FLAG_HP_AWARE",
        "AI_FLAG_NEGATE_UNAWARE",
        "AI_FLAG_WILL_SUICIDE",
        "AI_FLAG_HELP_PARTNER",
        "AI_FLAG_PREFER_STATUS_MOVES",
        // These two are unfinished so we shouldnt allow them to be used in the editor
        // "AI_FLAG_STALL",
        // "AI_FLAG_SCREENER",
        // This flag was removed
        // "AI_FLAG_DOUBLE_BATTLE",
        "AI_FLAG_SMART_SWITCHING",
        "AI_FLAG_ACE_POKEMON",
        "AI_FLAG_OMNISCIENT",
    }

    aiFlags := ParseAiFlags(input)

    if !(reflect.DeepEqual(expectedAiFlags, aiFlags)) {
        t.Errorf("AI Flags were not equal")
    }
}
