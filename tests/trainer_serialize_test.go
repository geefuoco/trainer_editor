package data_objects_tests

import (
    "testing"
    "github.com/geefuoco/trainer_editor/parsers"
    "os"
)

func TestSerializeTrainer(t *testing.T) {

    input := "test_cases/trainer_testcase.txt"

    file, err := os.ReadFile(input)
    if err != nil {
        t.Fatalf("Could not read file %s", input)
        return
    }

    expected := string(file)

    actual := parsers.ParseTrainers(input)[0]

    if actual.String() != expected {
        t.Log("Expected\n")
        t.Log(expected)
        t.Log("Actual\n")
        t.Log(actual.String())

        for i:=0; i < len(expected); i++ {
            t.Logf("%d: %c | %c", i, expected[i], actual.String()[i])
        }
        t.Errorf("Trainers did not match")
    }
}
