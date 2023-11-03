package data_objects

import (
    "strconv"
    "strings"
    "reflect"
    "github.com/geefuoco/trainer_editor/logging"
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

type PokemonSpeciesInfo struct {
    Species    string
    BaseHp     uint64
    BaseAtk    uint64
    BaseDef    uint64
    BaseSpd    uint64
    BaseSpAtk  uint64
    BaseSpDef  uint64
    Types      [2]string
}

func TemplateSpeciesInfo() *PokemonSpeciesInfo {
    return &PokemonSpeciesInfo {
        Species: "SPECIES_NONE",
        BaseHp: 0,
        BaseAtk: 0,
        BaseDef: 0,
        BaseSpd: 0, 
        BaseSpAtk: 0,
        BaseSpDef: 0,
    }
}

func TemplateMon() *TrainerMon {
    return &TrainerMon{
            Iv: [6]uint64{0, 0, 0, 0, 0, 0},
            Ev: [6]uint64{0, 0, 0, 0, 0, 0},
            Lvl: 1,
            Species: "SPECIES_NONE",
            HeldItem: "ITEM_NONE",
            Moves: [4]string{"MOVE_NONE", "MOVE_NONE", "MOVE_NONE", "MOVE_NONE"},
            IsShiny: false,
    }
}

func (sp* PokemonSpeciesInfo) String() string {
    padding := "    "
    var buf strings.Builder
    buf.WriteString(padding)
    buf.WriteString("["+sp.Species+"] = \n")
    buf.WriteString(padding)
    buf.WriteString("{\n")
    buf.WriteString(padding)
    buf.WriteString(padding)
    buf.WriteString(".baseHP        = ")
    value := strconv.FormatUint(sp.BaseHp, 10)
    buf.WriteString(value+",\n")
    buf.WriteString(padding)
    buf.WriteString(padding)
    buf.WriteString(".baseAttack    = ")
    value = strconv.FormatUint(sp.BaseAtk, 10)
    buf.WriteString(value+",\n")
    buf.WriteString(padding)
    buf.WriteString(padding)
    buf.WriteString(".baseDefense   = ")
    value = strconv.FormatUint(sp.BaseDef, 10)
    buf.WriteString(value+",\n")
    buf.WriteString(padding)
    buf.WriteString(padding)
    buf.WriteString(".baseSpeed     = ")
    value = strconv.FormatUint(sp.BaseSpd, 10)
    buf.WriteString(value+",\n")
    buf.WriteString(padding)
    buf.WriteString(padding)
    buf.WriteString(".baseSpAttack  = ")
    value = strconv.FormatUint(sp.BaseSpAtk, 10)
    buf.WriteString(value+",\n")
    buf.WriteString(padding)
    buf.WriteString(padding)
    buf.WriteString(".baseSpDefense = ")
    value = strconv.FormatUint(sp.BaseSpDef, 10)
    buf.WriteString(value+",\n")
    buf.WriteString(padding)
    buf.WriteString(padding)
    buf.WriteString(".types = {")
    for i, x := range sp.Types {
        buf.WriteString(x)
        if i == 0 {
            buf.WriteString(", ")
        }
    }
    buf.WriteString("},\n")
    buf.WriteString(padding)
    buf.WriteString("},\n")
    return buf.String()
}

func (mon *TrainerMon) String() string {
    templateMon := TemplateMon()
    padding := "    "
    var buf strings.Builder
    buf.WriteString(padding)
    buf.WriteString("{\n")
    // IVs
    if !reflect.DeepEqual(templateMon.Iv, mon.Iv) {
        buf.WriteString(padding)
        buf.WriteString(".iv = TRAINER_PARTY_IVS(")
        for i, iv := range mon.Iv {
            value := strconv.FormatUint(iv, 10)
            buf.WriteString(value)
            if i != 5 {
                buf.WriteString(", ")
            }
        }
        buf.WriteString("),\n")
    }
    // LEVEL
    buf.WriteString(padding)
    buf.WriteString(".lvl = ")
    level := strconv.FormatUint(mon.Lvl, 10)
    buf.WriteString(level)
    buf.WriteString(",\n")
    // SPECIES
    buf.WriteString(padding)
    buf.WriteString(".species = ")
    if mon.Species == "SPECIES_NONE" {
        logging.WarnLog("Tried to set mon with SPECIES_NONE. Setting to default")
        buf.WriteString("SPECIES_UNOWN" + ",\n")
    } else {
        buf.WriteString(mon.Species + ",\n")
    }
    // EVS
    if !reflect.DeepEqual(templateMon.Ev, mon.Ev) {
        buf.WriteString(padding)
        buf.WriteString(".ev = TRAINER_PARTY_EVS(")
        for i, ev := range mon.Ev {
            value := strconv.FormatUint(ev, 10)
            buf.WriteString(value)
            if i != 5 {
                buf.WriteString(", ")
            }
        }
        buf.WriteString("),\n")
    }
    // HELD ITEM
    if mon.HeldItem != "" && mon.HeldItem != templateMon.HeldItem {
        buf.WriteString(padding)
        buf.WriteString(".heldItem = ")
        buf.WriteString(mon.HeldItem + ",\n")
    }
    // MOVES
    if !reflect.DeepEqual(templateMon.Moves, mon.Moves) && !reflect.DeepEqual(mon.Moves, [4]string{"","","",""}) {
        buf.WriteString(padding)
        buf.WriteString(".moves = {")
        for i, move := range mon.Moves {
            if move == "" {
                move = "MOVE_NONE"
            }
            buf.WriteString(move)
            if i != 3 {
                buf.WriteString(", ")
            }
        }
        buf.WriteString("},\n")
    }
    if mon.Ability != templateMon.Ability {
        buf.WriteString(padding)
        buf.WriteString(".ability = ")
        buf.WriteString(mon.Ability+",\n")
    }
    // SHINY
    if mon.IsShiny {
        buf.WriteString(padding)
        buf.WriteString(".isShiny = TRUE,\n")
    }
    buf.WriteString(padding)
    buf.WriteString("}")
    return buf.String()
}

func (mon *TrainerMon) CalculateEvTotal() int {
    var total int
    for _, x := range mon.Ev {
        total += int(x)
    }
    return total
}

func (mon *TrainerMon) CalculateIvTotal() int {
    var total int
    for _, x := range mon.Iv {
        total += int(x)
    }
    return total
}

func (t* TrainerParty) String() string {
    var b strings.Builder
    b.WriteString("static const struct TrainerMon ")
    b.WriteString(t.Trainer + "[] = {\n")
    for i, mon := range(t.Party) {
        b.WriteString(mon.String())
        if i < len(t.Party)-1 {
            b.WriteString(",\n")
        }
    }
    b.WriteString("\n")
    b.WriteString("};\n")
    return b.String()
}

