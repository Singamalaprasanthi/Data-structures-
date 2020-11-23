// 1. Hashing

package main

import (
        "fmt"
        "hash/fnv"
)

func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}

func main() {
        fmt.Println(hash("Hello"))
        fmt.Println(hash("Hello."))
}

$go run go.main
4116459851
2770452991

// 2. Password Hashing example

package main

import (
    "fmt"

    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main() {
    password := "secret"
    hash, _ := HashPassword(password) // ignore error for the sake of simplicity

    fmt.Println("Password:", password)
    fmt.Println("Hash:    ", hash)

    match := CheckPasswordHash(password, hash)
    fmt.Println("Match:   ", match)
}

$go run go.main
Password: secret
Hash:     $2a$14$WT.s24nLJguThsqA4xMpuO82KIosLVZ9xGzzdZzaaBF8RnHfga.ji
Match:    true

// 3. example


package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	input := "Lorem Ipsum dolor sit Amet"
	md5 := md5.New()
	sha_256 := sha256.New()
	sha_512 := sha512.New()
	io.WriteString(md5, input)
	sha_256.Write([]byte(input))
	sha_512.Write([]byte(input))
	sha_512_256 := sha512.Sum512_256([]byte(input))
	hmac512 := hmac.New(sha512.New, []byte("secret"))
	hmac512.Write([]byte(input))

	
	fmt.Printf("md5:\t\t%x\n", md5.Sum(nil))

	
	fmt.Printf("sha256:\t\t%x\n", sha_256.Sum(nil))

	
	fmt.Printf("sha512:\t\t%x\n", sha_512.Sum(nil))

	
	fmt.Printf("sha512_256:\t%x\n", sha_512_256)

	
	fmt.Printf("hmac512:\t%s\n", base64.StdEncoding.EncodeToString(hmac512.Sum(nil)))
}

$go run go.main

md5:		4db45e622c0ae3157bdcb53e436c96c5
sha256:		eb7a03c377c28da97ae97884582e6bd07fa44724af99798b42593355e39f82cb
sha512:		5cdaf0d2f162f55ccc04a8639ee490c94f2faeab3ba57d3c50d41930a67b5fa6915a73d6c78048729772390136efed25b11858e7fc0eed1aa7a464163bd44b1c
sha512_256:	34c614af69a2550a4d39138c3756e2cc50b4e5495af3657e5b726c2ac12d5e60
hmac512:	GBZ7aqtVzXGdRfdXLHkb0ySp/f+vV9Zo099N+aSv+tTagUWuHrPeECDfUyd5WCoHBe7xkw2EdpyLWx+Ge4JQKg==

