// 1. url parsing

package main

	import "fmt"
import "net"
import "net/url"

	func main() {

	    s := "postgres://user:pass@host.com:5432/path?k=v#f"

	    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

	    fmt.Println(u.Scheme)

	    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)


	    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

	    fmt.Println(u.Path)
    fmt.Println(u.Fragment)
}
$go run main.go
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f

// 2. example
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://www.website.com/person?name=prasanthi%20reddy&phone=%2B919999999999&phone=%2B628888888888")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Scheme: ", u.Scheme)
	fmt.Println("Host: ", u.Host)

	queries := u.Query()
	fmt.Println("Query Strings: ")
	for key, value := range queries {
		fmt.Printf("  %v = %v\n", key, value)
	}
	fmt.Println("Path: ", u.Path)
}

$go run main.go
Scheme:  https
Host:  www.website.com
Query Strings: 
  name = [prasanthi reddy]
  phone = [+919999999999 +628888888888]
Path:  /person

// 3. url decoding
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	encodedValue := "Hell%C3%B6+W%C3%B6rld%40Golang"
	decodedValue, err := url.QueryUnescape(encodedValue)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(decodedValue)
}

$go run main.go
Hellö Wörld@Golang

// 4. url encoding

package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	queryStr := "name=prasanthi%20reddy&phone=%2B9199999999&phone=%2B628888888888"
	params, err := url.ParseQuery(queryStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Query Params: ")
	for key, value := range params {
		fmt.Printf("  %v = %v\n", key, value)
	}
}
$go run main.go
Query Params: 
  name = [prasanthi reddy]
  phone = [+9199999999 +628888888888]
