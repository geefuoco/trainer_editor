package data_objects

import (
    "strconv"
    "strings"
)

type TrainerParty struct {
    Trainer string
    Party []*TrainerMon
}

type TrainerMon struct {
    Iv [6]uint64
    Lvl uint64
    Species string
    Ev [6]uint64
    HeldItem string
    Moves [4]string
    Ability string
    IsShiny bool
}

func TemplateMon() *TrainerMon {
    return &TrainerMon{
            Iv: [6]uint64{0, 0, 0, 0, 0, 0},
            Ev: [6]uint64{0, 0, 0, 0, 0, 0},
            Lvl: 1,
            Species: "SPECIES_UNKNOWN",
            HeldItem: "ITEM_NONE",
            Moves: [4]string{"MOVE_HIDDEN_POWER", "MOVE_NONE", "MOVE_NONE", "MOVE_NONE"},
            IsShiny: false,
    }
}

func (t* TrainerParty) String() string {
    var b strings.Builder
    b.WriteString("TrainerParty: \n")
    b.WriteString("  Trainer: " + t.Trainer + "\n")
    for _, mon := range(t.Party) {
        b.WriteString(mon.String())
    }
    return b.String()
}

func (t *TrainerMon) String() string {
    var b strings.Builder
    b.WriteString("TrainerMon: \n")
    if len(t.Iv) == 0 {
        b.WriteString("  Iv: {}\n")
    } else {
        b.WriteString("  Iv: {\n")
        for i := range(t.Iv) {
            value := strconv.FormatUint(t.Iv[i], 10)
            b.WriteString("    " + value + "\n")
        }
        b.WriteString("  }\n")
    }
    b.WriteString("  Lvl: ")
    value := strconv.FormatUint(t.Lvl, 10)
    b.WriteString(value + "\n")
    b.WriteString("  Species: " + t.Species + "\n")
    if len(t.Ev) == 0 {
        b.WriteString("  Ev: {}\n")
    } else {
        b.WriteString("  Ev: {\n")
        for i := range(t.Ev) {
            value := strconv.FormatUint(t.Ev[i], 10)
            b.WriteString("    " + value + "\n")
        }
        b.WriteString("  }\n")
    }
    b.WriteString("  HeldItem: " + t.HeldItem + "\n")
    if len(t.Moves) == 0 {
        b.WriteString("  Moves: {}\n")
    } else {
        b.WriteString("  Moves: {\n")
        for i := range(t.Moves) {
            value := t.Moves[i]
            b.WriteString("    " + value + "\n")
        }
        b.WriteString("  }\n")
    }
    b.WriteString("  Ability: " + t.Ability + "\n")
    b.WriteString("  IsShiny: ")
    if t.IsShiny {
        b.WriteString("true\n")
    } else {
        b.WriteString("false\n")
    }
    return b.String()
}
