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

func TestFooTrue(t *testing.T) {
    result := FooTrue()
    if !result {
        t.Fatalf("It is impossible, but returned bool is wrong!")
    }
    t.Logf("Everything as expected")
}

