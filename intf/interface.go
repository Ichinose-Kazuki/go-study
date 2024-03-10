package intf

import "github.com/Ichinose-Kazuki/go-study/intf/internal/intf_sub"

func InterfaceMain() {
    var food intf_sub.Food
    food = intf_sub.Hamburg{"Mansei"}
    food.Eat()
}
