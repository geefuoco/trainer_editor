package data_objects_tests

import (
    "testing"
    "github.com/geefuoco/trainer_editor/parsers"
    "github.com/geefuoco/trainer_editor/data_objects"
    "os"
)

func TestSaveTrainers(t *testing.T) {

    input := "test_cases/trainer_save_testcase.txt"

    actual := parsers.ParseTrainers(input)

    if len(actual) != 4 {
        t.Fatalf("Expected 4 trainers, found %d", len(actual))
    }

    err := data_objects.SaveTrainers("test_cases/test_save_file.c", actual)
    if err != nil {
        panic(err)
    }

    // Compare the two files

    f1, err := os.ReadFile(input)
    if err != nil {
        panic(err)
    }

    f2, err := os.ReadFile("test_cases/test_save_file.c")
    if err != nil {
        panic(err)
    }

    if string(f1) != string(f2) {
        t.Errorf("Files were different.")
    }
}
