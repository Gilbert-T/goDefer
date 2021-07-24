package deferNotes

func order(){
    defer fmt.Println("1st defer.")
    defer fmt.Println("2nd defer.")
    defer fmt.Println("3rd defer.")
    fmt.Println("hello defer!")
    return
}
