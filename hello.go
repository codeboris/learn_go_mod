package hello

import "fmt"

var Version = "v1.0.0"

func Hello() {
    switch Version {
    case "v1.0.0":
        fmt.Println("Hello, v1.0.0")
    case "v1.1.0":
        fmt.Println("Hello, v1.1.0")
    default:
        fmt.Println("Unknown version")
    }
}
