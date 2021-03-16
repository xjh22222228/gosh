// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.7

package gtime

import (
    "github.com/xjh22222228/gosh/gtime/glocale"
    "time"
)

type Time struct {
    // locale name
    Language glocale.Language
}

func SetLocale(language glocale.Language) *Time {
    return &Time{
        Language: language,
    }
}

func (that Time) SetLocale(language glocale.Language) *Time {
    return SetLocale(language)
}

// Determine whether a leap year
func IsLeapYear(year int) bool {
    return year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)
}

// Get the number of days in the current month.
func DaysInMonth(year int, month int) (days int) {
    if month != 2 {
        if month == 4 || month == 6 || month == 9 || month == 11 {
            days = 30

        } else {
            days = 31
        }
    } else {
        if IsLeapYear(year) {
            days = 29
        } else {
            days = 28
        }
    }
    return
}

// Returns the number of days in leap years and peace years
func DaysInYear(year int) int {
    if IsLeapYear(year) {
        return 366
    }
    return 365
}

// The toISOString() method returns a string in simplified extended ISO
// format(ISO 8601). https://en.wikipedia.org/wiki/ISO_8601
// which is always 24 or 27 characters long
// (YYYY-MM-DDTHH:mm:ss.sssZ or ±YYYYYY-MM-DDTHH:mm:ss.sssZ,
// respectively). The timezone is always zero UTC offset,
// as denoted by the suffix "Z".
func ToISOString(t time.Time) string {
    u := t.UTC()
    return Format(u, "YYYY-MM-DDTHH:mm:ss.sssZ")
}

func IsToday(t time.Time) bool {
    ny, nm, nd := time.Now().Date()
    y, m, d := t.Date()
    return ny == y && nm == m && nd == d
}

