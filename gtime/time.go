// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.7

package gtime

import (
    "github.com/xjh22222228/gosh/ginternal"
    "github.com/xjh22222228/gosh/gstrings"
    "github.com/xjh22222228/gosh/gtime/locale"
    "regexp"
    "strconv"
    "strings"
    "time"
)

type Gtime struct {
    // locale name
    Language locale.Language
}

var regex = regexp.MustCompile(
    `\[([^\]]+)]|Y{1,4}|M{1,4}|D{1,2}|d{1,4}|H{1,2}|h{1,2}|a|A|m{1,2}|s{1,3}|Z{1,2}`)

func SetLocale(language locale.Language) *Gtime {
    return &Gtime{
        Language: language,
    }
}

func (that Gtime) SetLocale(language locale.Language) *Gtime {
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

func (that Gtime) Format(t time.Time, format string) string {
    return formatTime(t, format, that.Language)
}

func Format(t time.Time, format string) string {
    return formatTime(t, format, locale.ENUS)
}

// | Format| Output           | Description              |
// | ------------------------ |------------------ |
// | YY   | 18                | Two-digit year     |
// | YYYY | 2018              | Four-digit year|
// | M    | 1-12              | The month, beginning at 1 |
// | MM   | 01-12             | The month, 2-digits |
// | MMM  | Jan-Dec           | The abbreviated month name |
// | MMMM | January-December  | The full month name |
// | D    | 1-31              | The day of the month |
// | DD   | 01-31             | The day of the month, 2-digits |
// | d    | 0-6               | The day of the week, with Sunday as 0 |
// | dd   | Su-Sa             | The min name of the day of the week |
// | ddd  | Sun-Sat           | The short name of the day of the week |
// | dddd | Sunday-Saturday   | The name of the day of the week |
// | H    | 0-23              | The hour |
// | HH   | 00-23             | The hour, 2-digits |
// | h    | 1-12              | The hour, 12-hour clock |
// | hh   | 01-12             | The hour, 12-hour clock, 2-digits |
// | m    | 0-59              | The minute |
// | mm   | 00-59             | The minute, 2-digits |
// | s    | 0-59              | The second |
// | ss   | 00-59             | The second, 2-digits |
// | sss  | 000-999           | The millisecond, 3-digits |
// | A    | AM PM             |  |
// | a    | am pm             |  |
func formatTime(t time.Time, format string, language locale.Language) string {
    intl := locale.GetLocale(language)
    y, m, d := t.Date()
    hour, minute, sec := t.Clock()
    
    year := strconv.Itoa(y)
    month := int(m)
    monthStr := strconv.Itoa(month)
    day := strconv.Itoa(d)
    weekday := int(t.Weekday())
    hourStr := strconv.Itoa(hour)
    hour12 := strconv.Itoa(hourTo12(hour))
    minuteStr := strconv.Itoa(minute)
    second := strconv.Itoa(sec)
    unixStr := strconv.FormatInt(t.Unix(), 10)
    a := ""

    if len(unixStr) < 3 {
        unixStr += strings.Repeat("0", 3 - len(unixStr))
    }

    if intl.GetTime == nil {
        if hour < 12 {
            a = "AM"
        } else {
            a = "PM"
        }
    } else {
        a = intl.GetTime(hour, minute)
    }

    matches := map[string]string{
        "YY": gstrings.Slice(year, -2),
        "YYYY": year,
        "M":    monthStr,
        "MM":   ginternal.AddZero(monthStr),
        "MMM":  intl.ShortMonths[month - 1],
        "MMMM": intl.Months[month - 1],
        "D":    day,
        "DD":   ginternal.AddZero(day),
        "d":    strconv.Itoa(weekday),
        "dd":   intl.WeekdaysMin[weekday],
        "ddd":  intl.WeekdaysShort[weekday],
        "dddd": intl.Weekdays[weekday],
        "H":    hourStr,
        "HH":   ginternal.AddZero(hourStr),
        "h":    hour12,
        "hh":   ginternal.AddZero(hour12),
        "m":    minuteStr,
        "mm":   ginternal.AddZero(minuteStr),
        "s":    second,
        "ss":   ginternal.AddZero(second),
        "sss":  gstrings.Slice(unixStr, -3),
        "A":    a,
        "a":    strings.ToLower(a),
    }

    v := regex.ReplaceAllStringFunc(format, func(s string) string {
        v, ok := matches[s]
        if !ok {
            return s
        }
        return v
    })

    return v
}

func hourTo12(n int) int {
    if n > 13 {
        return n - 12
    }
    return n
}
