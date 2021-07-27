// File: program2.go
// Name: D.Saravanan
// Date: 19/06/2021
// Program for string modification

package main

import "fmt"

func main() {

    var name string = "Institute of Mathematical Sciences"

    // double quotes

    /* finds the length of a string */
    fmt.Println(len("Mathematical Sciences"))

    /* access a particular character in the string */
    fmt.Println("Mathematical Sciences"[1])
    /* strings are indexed at 0, not 1 */
    /* [1] gives you the second element, not the first */
    /* 97 instead of 'a', because the character is represented by a byte */

    fmt.Println(name[1])

    /* concatenates two strings together */
    fmt.Println("Mathematical " + "Sciences")

    institute := "Institute of Theoretical Physics"

    // backticks
    
    /* finds the length of a string */
    fmt.Println(len(`Mathematical Sciences`))

    /* access a particular character in the string */
    fmt.Println(`Mathematical Sciences`[1])
    /* strings are indexed at 0, not 1 */
    /* [1] gives you the second element, not the first */
    /* 97 instead of 'a', because the character is represented by a byte */

    fmt.Println(institute[1])

    /* concatenates two strings together */
    fmt.Println(`Mathematical ` + `Sciences`)
}
