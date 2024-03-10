package intf

import . "github.com/Ichinose-Kazuki/go-study/intf/internal/intf_sub"

func InterfaceMain() {
    var food Food
    food = Hamburg{"Mansei"}
    food.Eat()
}
