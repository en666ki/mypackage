package mypackage

import (
    "testing"
)

func TestGetInfo(t *testing.T) {
    result := GetInfo()
    if result != info {
        t.Fatalf("It is impossible, but returned constant is wrong!")
    }
    t.Logf("Everything as expected")
}
