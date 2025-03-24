#+feature dynamic-literals

package main

import "core:fmt"
import "core:strings"

fibonacci :: proc(n : int) -> int{
    switch{
        case n < 1:
            return 0;
        case n == 1:
            return 1;
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

multi_res :: proc() -> (int, int){
    return 1, 2
}


add_string :: proc(a: string, b : string) -> string{
    data:=make([]string, 2)
    defer delete(data)
    data[0] = a;
    data[1] = b;
    return strings.join(data, "-")
}

add_int :: proc(a : int, b : int) -> int{
    return a + b;
}
add :: proc{add_string, add_int}


person :: struct{
    name : string
}

employee :: struct{
    id : int,
    using person:person
}

get_name :: proc(self : ^person) -> string{
    return self.name
}

main :: proc(){

    emp := employee{
        id = 123,
        name = "what",
        person = {
            name  = "huh"
        }
    }
    emp1 := employee{
        id = 123,
        name = "what",
    }

    fmt.println(emp)
    fmt.println(emp1)


    some_map := map[string]int{ "A" = 1, "B" = 2 }

    fmt.println("Hello World")
    x : int
    x=13


    for i:=0;i < 10; i+=1 {
        fmt.print(i)
    }

    for key, &value in some_map {
        fmt.println(key, "=", value)
    }

    fmt.println("FIB ", fibonacci(3))
    fmt.println("Multi ", multi_res())
    fmt.println("adding integers ", add(1, 2))
    fmt.println("adding strings", add("hello", "data"))
    
  
}