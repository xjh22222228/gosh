package gtime

import (
    "github.com/stretchr/testify/assert"
    "github.com/xjh22222228/gosh/gtime/glocale"
    "strconv"
    "strings"
    "testing"
    "time"
)

func TestFormat(t *testing.T) {
    _assert := assert.New(t)
    // 21-周日Sunday
    date := time.Date(2021, 3, 21, 21, 5, 5, 0, &time.Location{})

    // ===================================================
    // ===================== ZH_CN =======================
    // ===================================================
    zhcn := SetLocale(glocale.ZHCN)
    _assert.Equal(
        "21-2021-3-03-3月-三月-21-21-0-日-周日-星期日-21-21-9-09-5-05-5-05-晚上-晚上",
        zhcn.Format(date,
            "YY-YYYY-M-MM-MMM-MMMM-D-DD-d-dd-ddd-dddd-H-HH-h-hh-m-mm-s-ss-A-a"),
    )
    _assert.Equal(
        "2021-03-21 21:05:05",
        zhcn.Format(date,
            "YYYY-MM-DD HH:mm:ss"),
    )


    // ===================================================
    // ===================== ZH_HK =======================
    // ===================================================
    zhhk := SetLocale(glocale.ZHHK)
    _assert.Equal(
        "3月/三月/0/日/週日/星期日-PM-pm",
        zhhk.Format(date,
            "MMM/MMMM/d/dd/ddd/dddd-A-a"),
    )


    // ===================================================
    // ===================== ZH_Tw =======================
    // ===================================================
    zhtw := SetLocale(glocale.ZHHK)
    _assert.Equal(
        "3月/三月/0/日/週日/星期日-PM-pm",
        zhtw.Format(date,
            "MMM/MMMM/d/dd/ddd/dddd-A-a"),
    )



    // ===================================================
    // =================== English =======================
    // ===================================================
    en := SetLocale(glocale.ENUS)
    _assert.Equal(
        "2021-03-21 21:05:05 PM-pm",
        en.Format(date,
            "YYYY-MM-DD HH:mm:ss A-a"),
    )


    // ===================================================
    // ==================== Format =======================
    // ===================================================
    _assert.Equal(
        "2021-03-21 21:05:05 PM-pm",
        Format(date,
            "YYYY-MM-DD HH:mm:ss A-a"),
    )

    // sss
    _assert.Equal(3, len(Format(date, "sss")))
    _, err := strconv.Atoi(Format(date, "sss"))
    _assert.Equal(nil, err)
}

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

func BenchmarkFormat(b *testing.B) {
    date := time.Date(2021, 3, 21, 21, 5, 5, 0, &time.Location{})
    zh := SetLocale(glocale.ZHCN)

    for i := 0; i < b.N; i++ {
        zh.Format(date,
            "YY-YYYY-M-MM-MMM-MMMM-D-DD-d-dd-ddd-dddd-H-HH-h-hh-m-mm-s-ss-A-a")
    }
}

