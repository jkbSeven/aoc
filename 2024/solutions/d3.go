package solutions

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var Day_3 = Day{P1:d3p1, P2: d3p2}

func d3p1() {
    f, err := os.ReadFile("inputs.txt")
    if err != nil {
        panic(err)
    }

    fString := string(f)

    regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
    result := regex.FindAllStringSubmatch(fString, -1)

    answer := int64(0)
    for _, match := range result {
        var n1, n2 int64

        if n1, err = strconv.ParseInt(match[1], 10, 64); err != nil {
            panic(err)
        }

        if n2, err = strconv.ParseInt(match[2], 10, 64); err != nil {
            panic(err)
        }

        answer += n1 * n2
    }

    fmt.Printf("Result: %d", answer)
}

func d3p2() {
    f, err := os.ReadFile("inputs.txt")
    if err != nil {
        panic(err)
    }

    fString := string(f)

    doRegex := regexp.MustCompile(`do\(\)`)
    doResult := doRegex.FindAllStringIndex(fString, -1)

    dontRegex := regexp.MustCompile(`don't\(\)`)
    dontResult := dontRegex.FindAllStringIndex(fString, -1)

    mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
    mulResult := mulRegex.FindAllStringSubmatchIndex(fString, -1)

    validRanges := make([][]int, 0, 32)
    validRanges = append(validRanges, []int{0, dontResult[0][0]})

    for _, doMatch := range doResult {
        found := false

        var left, right int
        left = doMatch[1]

        for i, dontMatch := range dontResult {
            right = dontMatch[0]

            if left > right {
                continue
            }

            found = true
            dontResult = dontResult[i:]
            break
        }

        if !found {
            right = len(fString)
        }

        validRanges = append(validRanges, []int{left, right})
    }

    answer := int64(0)
    for _, mulMatch := range mulResult {
        inside := false
        for _, validRange := range validRanges {
            i1 := mulMatch[0]
            i2 := mulMatch[1]

            if i1 >= validRange[0] && i2 <= validRange[1] {
                inside = true
            }
        }

        if !inside {
            continue
        }

        n1String := fString[mulMatch[2]:mulMatch[3]]
        n2String := fString[mulMatch[4]:mulMatch[5]]

        var n1, n2 int64
        if n1, err = strconv.ParseInt(n1String, 10, 64); err != nil {
            panic(err)
        }

        if n2, err = strconv.ParseInt(n2String, 10, 64); err != nil {
            panic(err)
        }

        answer += n1 * n2
    }

    fmt.Printf("Result: %d", answer)
}

