package main
 
import (
    "fmt"
)

func arabicToRoman(arg int)(string) {

    figure := []int{1000,100,10,1}
    romanDigitA := []string{
        1 : "I",
        10 : "X",
        100 : "C",
        1000 : "M",
    }
    romanDigitB := []string{
        1 : "V",
        10 : "L",
        100 : "D",
        1000 : "MMMMM",
    }

    if arg < 0 || arg > 4000 {
        return "ROMAN_OUT_OF_RANGE"
    }

    romanSlice := []string{}
    x := ""

    for _, f := range figure {
        digit, i, v := int(arg / f), romanDigitA[f], romanDigitB[f]
        fmt.Println(digit)
        fmt.Println(i)
        fmt.Println(v)
        switch digit {
            case 1: romanSlice = append(romanSlice, string(i))
            case 2: romanSlice = append(romanSlice, string(i) + string(i))
            case 3: romanSlice = append(romanSlice, string(i) + string(i) + string(i))
            case 4: romanSlice = append(romanSlice, string(i) + string(v))
            case 5: romanSlice = append(romanSlice, string(v))
            case 6: romanSlice = append(romanSlice, string(v) + string(i))
            case 7: romanSlice = append(romanSlice, string(v) + string(i)+ string(i))
            case 8: romanSlice = append(romanSlice, string(v) + string(i)+ string(i)+ string(i))
            case 9: romanSlice = append(romanSlice, string(i) + string(x))
        }
        arg -= digit * f
        fmt.Println(arg)
        x = i
    }
    ret := ""
    for _, e := range romanSlice {
        ret += e
    }
    return ret
}
 
func main() {
    fmt.Print("Enter Arabic Numerals: ")
    var arabicNumeral int
    fmt.Scanln(&arabicNumeral)
    fmt.Printf("%s: %s\n", "Roman Numeral", arabicToRoman(arabicNumeral))
}