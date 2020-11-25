// 1. json in go
//encoding struct

package main 
   
import ( 
    "fmt"
    "encoding/json"
) 
   
type Human struct{ 
       
    Name string 
    Age int
    Address string 
} 
   

func main() { 
       
human1 := Human{"prasanthi", 23, "New Delhi"} 
 human_enc, err := json.Marshal(human1) 
       
    if err != nil { 
           
       
        fmt.Println(err) 
    } 
       
   
    fmt.Println(string(human_enc)) 
     
    human2 := []Human{ 
        {Name: "Rahul", Age: 23, Address: "New Delhi"}, 
        {Name: "Priyanshi", Age: 20, Address: "Pune"}, 
        
    } 
       
    
    human2_enc, err := json.Marshal(human2) 
        
        if err != nil { 
       
        
            fmt.Println(err) 
        } 
           
    
    fmt.Println() 
        fmt.Println(string(human2_enc)) 
} 
$go run main.go
{"Name":"prasanthi","Age":23,"Address":"New Delhi"}

[{"Name":"Rahul","Age":23,"Address":"New Delhi"},{"Name":"Priyanshi","Age":20,"Address":"Pune"}]

// 2. Decoding in json

package main 
   
import ( 
    "fmt"
    "encoding/json"
) 
   

type Human struct{ 
       
    
    Name string 
    Address string 
    Age int
} 
   

func main() { 
       
    
    var human1 Human 
       
     
    Data := []byte(`{ 
        "Name": "prasanthi",   
        "Address": "Hyderabad", 
        "Age": 21 
    }`) 
       
    
    err := json.Unmarshal(Data, &human1) 
       
    if err != nil { 
           
        
            fmt.Println(err) 
    } 
       
   
    fmt.Println("Struct is:", human1) 
    fmt.Printf("%s lives in %s.\n", human1.Name, human1.Address) 
       
   
    var human2 []Human 
       
    
    Data2 := []byte(` 
    [ 
        {"Name": "Vani", "Address": "Delhi", "Age": 21}, 
        {"Name": "Rashi", "Address": "Noida", "Age": 24}, 
        {"Name": "Rohit", "Address": "Pune", "Age": 25} 
    ]`) 
       
    
    err2 := json.Unmarshal(Data2, &human2) 
       
        if err2 != nil { 
       
        
            fmt.Println(err2) 
        } 
       
    
    for i := range human2{ 
       
        fmt.Println(human2[i]) 
    } 
} 

$go run main.go
Struct is: {prasanthi Hyderabad 21}
prasanthi lives in Hyderabad.
{Vani Delhi 21}
{Rashi Noida 24}
{Rohit Pune 25}



