package main

import (
    "flag"
    "2024/solutions"
)

func initDays() map[int]solutions.Day {
    days := make(map[int]solutions.Day)

    days[3] = solutions.Day_3

    return days
}

func main() {
    days := initDays()

    day_number := flag.Int("d", 3, "day number")
    part_number := flag.Int("p", 1, "part number")

    flag.Parse()

    day, ok := days[*day_number]

    if !ok {
        panic("Invalid day number")
    }

    if *part_number == 1 {
        day.P1()
    } else if *part_number == 2 {
        day.P2()
    } else {
        panic("Invalid part number")
    }
}
