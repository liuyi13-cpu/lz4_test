package main

import "fmt"

/*
 * 输入：abcde_bcdefgh_abcdefghxxxxxxx
 * 输出：abcde_(5,4)fgh_(14,5)fghxxxxxxx
 * 格式：[token]literals(offset,match length)[token]literals(offset,match length)....
 */
func main() {
    var a = "abcde_bcdefgh_abcdefghxxxxxxx"
    fmt.Println(a)
    fmt.Println("Hello, World!")
}