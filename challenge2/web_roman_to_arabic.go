package main

import("fmt"
		"net/http"
        "html/template")

var tpl *template.Template

var m = map[rune]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/process", processHandler)
    http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "roman_to_arabic.gohtml", nil)
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	romanNumeral := r.FormValue("roman_number")
	fmt.Fprintf(w, "%s: %d\n", "Arabic Numerals Equivalent", romanToArabic(romanNumeral))
}

func romanToArabic(roman string) (arabic int) {
    lastDigit := 1000
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
