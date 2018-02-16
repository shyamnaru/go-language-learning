package main
 
import (
    "fmt"
)
 
var m = map[rune]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}

func romanToArabic(roman string) (arabic int) {
    lastDigit := 1000
    // _ will ignore index
    for _, r := range roman {
        digit := m[r]
        if lastDigit < digit {
            arabic -= 2 * lastDigit
        }
        lastDigit = digit
        arabic += digit
    }
 
    return arabic
}
 
func main() {
    fmt.Print("Enter Roman Numerals: ")
    var romanNumeral string
    fmt.Scanln(&romanNumeral)
    fmt.Printf("%s: %d\n", "Arabic Numerals Equivalent", romanToArabic(romanNumeral))
}