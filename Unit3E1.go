// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-08
// Fileoverview: This program prompts the user to enter the names of the ten Canadian
// provinces, three territories, and their capital cities. It then outputs the
// information in two columns: Province/Territory and Capital.

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // variables
    var region1, capital1 string
    var region2, capital2 string
    var region3, capital3 string
    var region4, capital4 string
    var region5, capital5 string
    var region6, capital6 string
    var region7, capital7 string
    var region8, capital8 string
    var region9, capital9 string
    var region10, capital10 string
    var region11, capital11 string
    var region12, capital12 string
    var region13, capital13 string

    reader := bufio.NewReader(os.Stdin)

    // input
    fmt.Print("Enter the name of province #1: ")
    region1, _ = reader.ReadString('\n')
    region1 = strings.TrimSpace(region1)
    fmt.Print("Enter the capital city of " + region1 + ": ")
    capital1, _ = reader.ReadString('\n')
    capital1 = strings.TrimSpace(capital1)

    fmt.Print("Enter the name of province #2: ")
    region2, _ = reader.ReadString('\n')
    region2 = strings.TrimSpace(region2)
    fmt.Print("Enter the capital city of " + region2 + ": ")
    capital2, _ = reader.ReadString('\n')
    capital2 = strings.TrimSpace(capital2)

    fmt.Print("Enter the name of province #3: ")
    region3, _ = reader.ReadString('\n')
    region3 = strings.TrimSpace(region3)
    fmt.Print("Enter the capital city of " + region3 + ": ")
    capital3, _ = reader.ReadString('\n')
    capital3 = strings.TrimSpace(capital3)

    fmt.Print("Enter the name of province #4: ")
    region4, _ = reader.ReadString('\n')
    region4 = strings.TrimSpace(region4)
    fmt.Print("Enter the capital city of " + region4 + ": ")
    capital4, _ = reader.ReadString('\n')
    capital4 = strings.TrimSpace(capital4)

    fmt.Print("Enter the name of province #5: ")
    region5, _ = reader.ReadString('\n')
    region5 = strings.TrimSpace(region5)
    fmt.Print("Enter the capital city of " + region5 + ": ")
    capital5, _ = reader.ReadString('\n')
    capital5 = strings.TrimSpace(capital5)

    fmt.Print("Enter the name of province #6: ")
    region6, _ = reader.ReadString('\n')
    region6 = strings.TrimSpace(region6)
    fmt.Print("Enter the capital city of " + region6 + ": ")
    capital6, _ = reader.ReadString('\n')
    capital6 = strings.TrimSpace(capital6)

    fmt.Print("Enter the name of province #7: ")
    region7, _ = reader.ReadString('\n')
    region7 = strings.TrimSpace(region7)
    fmt.Print("Enter the capital city of " + region7 + ": ")
    capital7, _ = reader.ReadString('\n')
    capital7 = strings.TrimSpace(capital7)

    fmt.Print("Enter the name of province #8: ")
    region8, _ = reader.ReadString('\n')
    region8 = strings.TrimSpace(region8)
    fmt.Print("Enter the capital city of " + region8 + ": ")
    capital8, _ = reader.ReadString('\n')
    capital8 = strings.TrimSpace(capital8)

    fmt.Print("Enter the name of province #9: ")
    region9, _ = reader.ReadString('\n')
    region9 = strings.TrimSpace(region9)
    fmt.Print("Enter the capital city of " + region9 + ": ")
    capital9, _ = reader.ReadString('\n')
    capital9 = strings.TrimSpace(capital9)

    fmt.Print("Enter the name of province #10: ")
    region10, _ = reader.ReadString('\n')
    region10 = strings.TrimSpace(region10)
    fmt.Print("Enter the capital city of " + region10 + ": ")
    capital10, _ = reader.ReadString('\n')
    capital10 = strings.TrimSpace(capital10)

    fmt.Print("Enter the name of territory #1: ")
    region11, _ = reader.ReadString('\n')
    region11 = strings.TrimSpace(region11)
    fmt.Print("Enter the capital city of " + region11 + ": ")
    capital11, _ = reader.ReadString('\n')
    capital11 = strings.TrimSpace(capital11)

    fmt.Print("Enter the name of territory #2: ")
    region12, _ = reader.ReadString('\n')
    region12 = strings.TrimSpace(region12)
    fmt.Print("Enter the capital city of " + region12 + ": ")
    capital12, _ = reader.ReadString('\n')
    capital12 = strings.TrimSpace(capital12)

    fmt.Print("Enter the name of territory #3: ")
    region13, _ = reader.ReadString('\n')
    region13 = strings.TrimSpace(region13)
    fmt.Print("Enter the capital city of " + region13 + ": ")
    capital13, _ = reader.ReadString('\n')
    capital13 = strings.TrimSpace(capital13)

    // output
    fmt.Println("\nProvince/Territory\t\tCapital")
    fmt.Println("------------------------------------------")
    fmt.Println(region1 + "\t\t" + capital1)
    fmt.Println(region2 + "\t\t" + capital2)
    fmt.Println(region3 + "\t\t" + capital3)
    fmt.Println(region4 + "\t\t" + capital4)
    fmt.Println(region5 + "\t\t" + capital5)
    fmt.Println(region6 + "\t\t" + capital6)
    fmt.Println(region7 + "\t\t" + capital7)
    fmt.Println(region8 + "\t\t" + capital8)
    fmt.Println(region9 + "\t\t" + capital9)
    fmt.Println(region10 + "\t\t" + capital10)
    fmt.Println(region11 + "\t\t" + capital11)
    fmt.Println(region12 + "\t\t" + capital12)
    fmt.Println(region13 + "\t\t" + capital13)

    fmt.Println("\nDone.")
}
