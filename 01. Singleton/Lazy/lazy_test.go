package Lazy

import (
    "testing"
    "time"
)

func TestGetInstance(t *testing.T) {
    for i := 0; i != 2048; i++ {
        go func() {
            instance := GetInstance()
            instance.Nothing()

        }()
    }

    time.Sleep(4 * time.Second)
}

func TestGetInstanceOnce(t *testing.T) {
    for i := 0; i != 2048; i++ {
        go func() {
            instance := GetInstanceOnce()
            instance.Nothing()

        }()
    }

    time.Sleep(40 * time.Second)
}
