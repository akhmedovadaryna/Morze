package main

import (
        "fmt"
        "strings"
)


// this is a comment

func main() {
    var text string =  ".-___.-_.-_.-___.-"
     morze := map[string]string{
             ".-"   : "A",
             "-..." : "B",
             "-.-." : "C",
             "-.."  : "D",
             "."    : "E",
             "..-." : "F",
             "--."  : "G",
             "...." : "H",
             ".."   : "I" ,
             ".---" : "J",
             "-.-"  : "K",
             ".-.." : "L",
             "--"   : "M",
             "-."   : "N",
             "---"  : "O" ,
             ".--." : "P",
             "--.-" : "Q",
             ".-."  : "R" ,
             "..."  : "S" ,
             "-"    : "T",
             "..-"  : "U" ,
             "...-" : "V",
             ".--"  : "W" ,
             "-..-" : "X" ,
             "-.--" : "Y",
             "--.." : "Z",
    }
        fmt.Println(morze["--.."])


    morz(morze, text)
}

func morz(morze map[string]string, text string)  {
    var final string
    keys := make([]string, 0, len(morze))
    for k := range morze {
        keys = append(keys, k)
    }

    fmt.Println(text)
        spl := strings.Split(text, "___")
        //fmt.Printf("%q\n", spl)
        for e:=0;e<len(spl);e = e+1 {     // words
                spl1 := strings.Split(spl[e], "_")
               // fmt.Printf("%q\n", spl1)
                for i:=0;i<len(spl1);i = i+1{ //word
                        fmt.Println(spl1[i])
                        for j:=0; j<len(keys);j=j+1{
                                if keys[j] == spl1[i]{
                                        final = strings.Join([]string{final,morze[keys[j]]}, "")
                                        fmt.Println(morze[keys[j]])
                                        break
                                }
                        }

                }
                final = strings.Join([]string{final," "}, "")
                fmt.Println("-----------------------------------------------------------------")
        }
    fmt.Println(final)




}

