package main

import (
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func main() {
    str, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    // regex to match the pattern mul(XXX, YYY)
    re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
    matches := re.FindAllString(string(str), -1)

    count := 0

    //fmt.Println(matches)
    for _, match := range matches {
        //fmt.Println(match)
        // regex to extract the numbers from the pattern
        re2 := regexp.MustCompile(`[0-9]{1,3}`)
        nums := re2.FindAllString(match, -1)
        //fmt.Println(nums)
        num1, _ := strconv.Atoi(nums[0])
        num2, _ := strconv.Atoi(nums[1])
        //fmt.Println(num1, num2)
        result := int(num1) * int(num2)
        fmt.Println(result)
        count += result
    }
    fmt.Println(count)
}
