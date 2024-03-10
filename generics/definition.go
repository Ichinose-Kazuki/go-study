package generics

type photoPrint struct {
    content string 
}

func (p photoPrint) String() string {
    return p.content
}

type digitalPhoto struct { 
    content string
}

func (d digitalPhoto) String() string {
    return d.content
}
