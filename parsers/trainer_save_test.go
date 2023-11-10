package parsers

import (
    "testing"
    "github.com/geefuoco/trainer_editor/data_objects"
    "strings"
    "os"
)

func TestSaveTrainers(t *testing.T) {
    t.Parallel()

    input := "testdata/trainer_save_testcase.txt"

    actual := ParseTrainers(input)

    if len(actual) != 4 {
        t.Fatalf("Expected 4 trainers, found %d", len(actual))
    }

    err := data_objects.SaveTrainers("testdata/test_save_file.c", actual)
    if err != nil {
        panic(err)
    }

    // Compare the two files

    f1, err := os.ReadFile(input)
    if err != nil {
        panic(err)
    }

    f2, err := os.ReadFile("testdata/test_save_file.c")
    if err != nil {
        panic(err)
    }

    if strings.TrimSpace(string(f1)) != strings.TrimSpace(string(f2)) {
        t.Errorf("Files were different.")
    }
}
