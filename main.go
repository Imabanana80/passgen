package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

//go:embed wordsA.txt
var wordsA string

func main() {
    random := rand.New(rand.NewSource(time.Now().UnixNano()))
    gen(random) 
}

func gen(random *rand.Rand) {
    var lines []string = strings.Split(strings.ReplaceAll(wordsA, "\r\n", "\n"), "\n")
    var chars = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
        "!", "?", "@", "#", "$", "%", "^", "&", "*", "-", "_", "+", "=", "/", "~"}
    var caps = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
    password := ""
    rounds := 4 + random.Intn(2);  
    for i := 0; i<rounds; i++ {
        password += lines[random.Intn(len(lines))]
        for o := 0; o < random.Intn(2); o++ {
            password += caps[random.Intn(len(caps))]
        }
        password += chars[random.Intn(len(chars))]
        for o := 0; o < random.Intn(3); o++ {
            password += chars[random.Intn(len(chars))]
        }
    }
    clipboard.WriteAll(password)
    fmt.Println("Copied to clipboard")
}
