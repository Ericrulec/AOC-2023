package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func reverse(str string) (result string) { 
    for _, v := range str { 
        result = string(v) + result 
    } 
    return
} 

func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    var sum int = 0;
    scanner := bufio.NewScanner(file)

    replacer:= strings.NewReplacer("one","1","two","2","three","3","four","4","five","5","six","6","seven","7","eight","8","nine","9")
    for scanner.Scan() {
        old_line:= scanner.Text()

        var first string;
        var last string;

        var line string;
        line = replacer.Replace(old_line)
        for _,char:=range line {
            if unicode.IsNumber(char){
                first = string(char)
                break
            }
        }
        var str string;
        for _,char := range reverse(old_line){
            str = string(char) + str
            str = replacer.Replace(str)
            if unicode.IsNumber(rune(str[0])){
                last = string(str[0])
                break
            }
            if unicode.IsNumber(char) && err != nil{
                last = string(char)
                break
            }
        }
        double  :=first + last
        fmt.Println("First: , Second:",first,last,double)
        num,_ := strconv.Atoi(double)
        sum += num
    }
    fmt.Println("Total sum:",sum)
}
