/* RZFeeser | Alta3 Research
   switch - case and default  */

package main

import (
    "fmt"
    "runtime"
    "strings"
)

func main() {
    fmt.Println(runtime.Version())
    gove := runtime.Version()
           // init gover                
    switch {
    case strings.Contains(gove, "go1.17"):                 // if matches "go1.17", do the code below then STOP
        fmt.Println("You are using the latest version of GoLang")
    case strings.Contains(gove, "go1.16"):       // if matches "go1.16", "go1.16.5", OR "go1.15" 
        fmt.Println("You are using an older, but acceptable version of GoLang")
    default:                       // in all other cases run the code below
        fmt.Println("I cannot make a recommendation.")
    }
}

