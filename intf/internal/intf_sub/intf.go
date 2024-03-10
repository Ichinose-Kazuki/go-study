package intf_sub

type Food interface {
    Eat()
}

type Hamburg struct {
    // Must be capitalized to be initialized from another package.
    Restaurant string
}


