package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "math"
    "bufio"
    "log"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    count := 0

    for scanner.Scan() {
        line := scanner.Text()
        levels := strings.Fields(line)

        levelList := []int{}
        for _, level := range levels {
            levelInt, _ := strconv.Atoi(level)
            levelList = append(levelList, levelInt)
        }

        isUp := false
        isDown := false

        fmt.Println(levelList)

        for i, level := range levelList {
            if i == len(levelList) - 1 {
                fmt.Println("End of list, good report")
                count++
                break
            }
            if level == levelList[i+1] {
                fmt.Println("Same level")
                break
            }

            diff := level - levelList[i+1]
            fmt.Printf("\ndiff between %d and %d: %d", level, levelList[i+1], diff)
            if diff < 0 {
                fmt.Println("Going up")
                isUp = true
            } else {
                fmt.Println("Going down")
                isDown = true
            }

            if isUp && isDown {
                fmt.Println("Going up and down")
                break
            }

            absDiff := math.Abs(float64(diff))
            fmt.Printf("\nabsDiff: %f", absDiff)
            if absDiff > 3 || absDiff < 1 {
                fmt.Println("Too much difference")
                break
            } else {
                fmt.Println("Good difference")
            }
        }
    }
    fmt.Println(count)
}
