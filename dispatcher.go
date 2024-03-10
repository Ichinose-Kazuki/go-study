package main

import (
    "os"
    "github.com/Ichinose-Kazuki/go-study/hello"
    "github.com/Ichinose-Kazuki/go-study/intf"
    "github.com/Ichinose-Kazuki/go-study/generics"
)

func main() {
    if len(os.Args) != 2 {
        panic("Usage: dispatcher <project name>")
    }
    projectName := os.Args[1]

    switch projectName {
    case "hello":
        hello.HelloMain()
    case "interface":
        intf.InterfaceMain()
    case "generics":
        generics.GenericsMain()
    }
}
