package main

import (
    "io/ioutil"
    "fmt"
    "bufio"
    "os"
)

func main() {
    stdin := bufio.NewReader(os.Stdin)

    fmt.Print("Enter file path: ")
    filePath, _ := stdin.ReadString('\n')
    filePath = filePath[:len(filePath) - 2] //remove newline character from file path
    fileBytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println("Could not read file")
        return
    }

    fmt.Print("Enter password: ")
    password, _ := stdin.ReadString('\n')
    password = password[:len(password) - 2] //remove newline character from password
    passwordBytes := []byte(password)

    elongatedPasswordBytes := passwordBytes //passwordBytes must be extended to the length of the file for XOR to work
    for len(elongatedPasswordBytes) < len(fileBytes) {
        elongatedPasswordBytes = append(elongatedPasswordBytes, passwordBytes...)
    }

    encryptedBytes := make([]byte, len(fileBytes))
    for i, v := range fileBytes {
        encryptedBytes[i] = v ^ elongatedPasswordBytes[i]
    }

    ioutil.WriteFile(filePath, encryptedBytes, 0666)

}