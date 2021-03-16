package gtime

import (
    "github.com/stretchr/testify/assert"
    "strings"
    "testing"
    "time"
)

func TestIsLeapYear(t *testing.T) {
    _assert := assert.New(t)
    _assert.True(IsLeapYear(1880))
    _assert.True(IsLeapYear(1884))
    _assert.True(IsLeapYear(1888))
    _assert.True(IsLeapYear(1892))
    _assert.True(IsLeapYear(1896))
    _assert.True(IsLeapYear(1904))
    _assert.True(IsLeapYear(1908))
    _assert.True(IsLeapYear(1912))
    _assert.True(IsLeapYear(1916))
    _assert.True(IsLeapYear(1920))
    _assert.True(IsLeapYear(1924))
    _assert.True(IsLeapYear(1928))
    _assert.True(IsLeapYear(1932))
    _assert.True(IsLeapYear(1936))
    _assert.True(IsLeapYear(1940))
    _assert.True(IsLeapYear(1944))
    _assert.True(IsLeapYear(1948))
    _assert.True(IsLeapYear(1952))
    _assert.True(IsLeapYear(1956))
    _assert.True(IsLeapYear(1960))
    _assert.True(IsLeapYear(1964))
    _assert.True(IsLeapYear(2020))
    _assert.False(IsLeapYear(2021))
}

func TestDaysInMonth(t *testing.T) {
    _assert := assert.New(t)
    // 2020
    _assert.Equal(31, DaysInMonth(2020, 1))
    _assert.Equal(29, DaysInMonth(2020, 2))

    // 2021
    _assert.Equal(31, DaysInMonth(2021, 1))
    _assert.Equal(28, DaysInMonth(2021, 2))
    _assert.Equal(31, DaysInMonth(2021, 3))
    _assert.Equal(30, DaysInMonth(2021, 4))
    _assert.Equal(31, DaysInMonth(2021, 5))
    _assert.Equal(30, DaysInMonth(2021, 6))
    _assert.Equal(31, DaysInMonth(2021, 7))
    _assert.Equal(31, DaysInMonth(2021, 8))
    _assert.Equal(30, DaysInMonth(2021, 9))
    _assert.Equal(31, DaysInMonth(2021, 10))
    _assert.Equal(30, DaysInMonth(2021, 11))
    _assert.Equal(31, DaysInMonth(2021, 12))
}

func TestDaysInYear(t *testing.T) {
    _assert := assert.New(t)
    _assert.Equal(366, DaysInYear(2020))
    _assert.Equal(365, DaysInYear(2021))
}

func TestToISOString(t *testing.T) {
    _assert := assert.New(t)
    date := time.Date(2021, 3, 21, 21, 5, 5, 0, &time.Location{})

    _assert.True(
        strings.Contains(ToISOString(date), "2021-03-21T21:05:05."),
        )
}

