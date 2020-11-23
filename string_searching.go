
// 1. string searching

package main
 
import (
  "fmt"
  "strings"
)
 
func main() {
  fmt.Println(strings.Compare("A", "B"))  // A < B
  fmt.Println(strings.Compare("B", "A"))  // B > A  
  fmt.Println(strings.Compare("Japan", "Australia"))  // J > A
  fmt.Println(strings.Compare("Australia", "Japan"))  // A < J
  fmt.Println(strings.Compare("Germany", "Germany"))  // G == G
  fmt.Println(strings.Compare("Germany", "GERMANY"))  // GERMANY > Germany
  fmt.Println(strings.Compare("", ""))
  fmt.Println(strings.Compare("", " ")) // Space is less
}

$go run main.go
-1
1
1
-1
0
1
0
-1

// 2.string index

package main
 
import (
  "fmt"
  "strings"
)
 
func main() {
  fmt.Println(strings.Index("Australia", "Aus"))
  fmt.Println(strings.Index("Australia", "aus"))
  fmt.Println(strings.Index("Australia", "A"))
  fmt.Println(strings.Index("Australia", "a"))
  fmt.Println(strings.Index("Australia", "Jap"))
  fmt.Println(strings.Index("Japan-124", "-"))
  fmt.Println(strings.Index("Japan-124", ""))
}

$go run main.go
0
-1
0
5
-1
5
0
