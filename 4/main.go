package main

import (
    "fmt"
    "os"
    "bufio"
    //"strings"
    "log"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    rows, cols := 0, 0

    for scanner.Scan() {
        cols = len(scanner.Text())
        rows++
    }

    grid := make([][]string, rows)
    for i := 0; i < rows; i++ {
        grid[i] = make([]string, cols)
    }

    file.Close()

    file2, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    // fill the grid with the letters from the input file
    charScanner := bufio.NewScanner(file2)
    gridLineNum := 0

    for charScanner.Scan() {
        line := charScanner.Text()
        for i, char := range line {
            grid[gridLineNum][i] = string(char)
        }
        gridLineNum++
    }

    //fmt.Println(grid)
    file2.Close()

    count := 0

    // loop over grid
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == "X" {
                // check surrounding cells for the letter M
                if grid[i+1][j] == "M" {
                    if grid[i+2][j] == "A" {
                        if grid[i+3][j] == "S" {
                            count++ } } } } } }
    
    fmt.Println(count)

}
