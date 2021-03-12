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
    `\[([^\]]+)]|Y{1,4}|M{1,4}|D{1,2}|d{1,4}|H{1,2}|h{1,2}|a|A|m{1,2}|s{1,2}|Z{1,2}|SSS`)

func SetLocale(language locale.Language) *Gtime {
    return &Gtime{
        Language: language,
    }
}

func (that Gtime) Format(t time.Time, format string) string {
    return formatTime(t, format, that.Language)
}

func (that Gtime) SetLocale(language locale.Language) *Gtime {
    return SetLocale(language)
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
// | SSS  | 000-999           | The millisecond, 3-digits |
// | A    | AM PM             |  |
// | a    | am pm             |  |
func formatTime(t time.Time, format string, language locale.Language) string {
    intl := getLocale(language)
    y, m, d := t.Date()

    year := strconv.Itoa(y)
    month := int(m)
    day := strconv.Itoa(d)
    weekday := int(t.Weekday())
    hour := t.Hour()
    hourStr := strconv.Itoa(hour)
    hour12 := strconv.Itoa(hourTo12(hour))
    minute := t.Minute()
    minuteStr := strconv.Itoa(t.Minute())
    second := strconv.Itoa(t.Second())
    a := ""

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
        "M": strconv.Itoa(month),
        "MM": ginternal.AddZero(strconv.Itoa(month)),
        "MMM": intl.ShortMonths[month - 1],
        "MMMM": intl.Months[month - 1],
        "D": day,
        "DD": ginternal.AddZero(day),
        "d": strconv.Itoa(weekday),
        "dd": intl.WeekdaysMin[weekday],
        "ddd": intl.WeekdaysShort[weekday],
        "dddd": intl.Weekdays[weekday],
        "H": hourStr,
        "HH": ginternal.AddZero(hourStr),
        "h": hour12,
        "hh": ginternal.AddZero(hour12),
        "m": minuteStr,
        "mm": ginternal.AddZero(minuteStr),
        "s": second,
        "ss": ginternal.AddZero(second),
        "SSS": gstrings.Slice(strconv.FormatInt(t.Unix(), 10), -3),
        "A": a,
        "a": strings.ToLower(a),
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

func getLocale(l locale.Language) *locale.Locale {
    switch l {
    case locale.ENUS:
        return locale.EnUS
    case locale.ZHCN:
        return locale.ZhCN
    }
    return locale.EnUS
}

func hourTo12(n int) int {
    if n > 13 {
        return n - 12
    }
    return n
}
