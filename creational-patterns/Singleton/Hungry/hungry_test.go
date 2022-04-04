package Hungry

import "testing"

func TestGetInstance(t *testing.T) {
    foo := GetInstance()
    bar := GetInstance()

    t.Log(foo == bar)
}
