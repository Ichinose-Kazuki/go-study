package main

import (
    "fmt"

    "github.com/Ichinose-Kazuki/go-study/hello"
    "github.com/Ichinose-Kazuki/go-study/intf"
    "github.com/Ichinose-Kazuki/go-study/generics"

    "github.com/manifoldco/promptui"
)

func main() {

prompt := promptui.Select{
        Label: "Select Mode",
        Items: []string{"hello", "interface", "generics"},
    }

    _, mode, err := prompt.Run()

    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        return
    }

    fmt.Printf("You choose %q:\n\n", mode)

    switch mode {
    case "hello":
        hello.HelloMain()
    case "interface":
        intf.InterfaceMain()
    case "generics":
        generics.GenericsMain()
    }
}
