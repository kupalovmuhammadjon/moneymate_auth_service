package main

import (
    "fmt"
    "log"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    passwords := []string{
        "1",
        "Admin2024$",
        "UserPass456@",
        "Secure789#",
        "MyPassword01!",
        "Password2024$",
        "User2024@!",
        "SecurePass2024#",
        "NewPassword123!",
        "Access2024$",
    }

    for _, password := range passwords {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            log.Fatalf("Failed to hash password: %v", err)
        }
        fmt.Printf("Plain: %s\nHashed: %s\n\n", password, hashedPassword)
    }
}
