package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
	"github.com/atotto/clipboard"
)

func main() {
    random := rand.New(rand.NewSource(time.Now().UnixNano()))
    gen(random) 
}

func gen(random *rand.Rand) {
    file, _ := os.Open("wordsA.txt")
    defer file.Close()
    
    var chars = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
        "!", "?", "@", "#", "$", "%", "^", "&", "*", "-", "_", "+", "=", "/", "~"}
    var caps = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
    var lines []string
    scanner := bufio.NewScanner(file)
    for (scanner.Scan()) {
        lines = append(lines, scanner.Text())
    }
    
    password := ""
    rounds := 3 + random.Intn(3);  
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
