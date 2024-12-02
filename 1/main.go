package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
    "sort"
    "math"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    list1 := []int{}
    list2 := []int{}
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        line := scanner.Text()
        //fmt.Println(line)
        items := strings.Fields(line)
        //fmt.Println(items)
        l1_item, _ := strconv.Atoi(items[0])
        //fmt.Println(l1_item)
        l2_item, _ := strconv.Atoi(items[1])
        //fmt.Println(l2_item)
        list1 = append(list1, l1_item)
        list2 = append(list2, l2_item)
    }

    sort.Ints(list1)
    sort.Ints(list2)
    fmt.Println(list1)
    fmt.Println(list2)

    sum := 0
    for i := 0; i < len(list1); i++ {
        //fmt.Println(list1[i], list2[i])
        diff := int(math.Abs(float64(list1[i] - list2[i])))
        sum += diff
    }
    fmt.Println(sum)
}
