package gtime

import (
    "github.com/stretchr/testify/assert"
    "github.com/xjh22222228/gosh/gtime/locale"
    "strconv"
    "testing"
    "time"
)

func TestFormat(t *testing.T) {
    _assert := assert.New(t)
    // 周日
    date := time.Date(2021, 3, 21, 21, 5, 5, 0, &time.Location{})

    // Chinese
    zh := SetLocale(locale.ZHCN)
    _assert.Equal(
        "21-2021-3-03-3月-三月-21-21-0-日-周日-星期日-21-21-9-09-5-05-5-05-晚上-晚上",
        zh.Format(date,
            "YY-YYYY-M-MM-MMM-MMMM-D-DD-d-dd-ddd-dddd-H-HH-h-hh-m-mm-s-ss-A-a"),
    )
    _assert.Equal(
        "2021-03-21 21:05:05",
        zh.Format(date,
            "YYYY-MM-DD HH:mm:ss"),
    )

    // ===================================================
    // =================== English =======================
    // ===================================================
    en := SetLocale(locale.ENUS)
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

    // SSS
    _assert.Equal(3, len(Format(date, "SSS")))
    _, err := strconv.Atoi(Format(date, "SSS"))
    _assert.Equal(nil, err)
}

func BenchmarkFormat(b *testing.B) {
    date := time.Date(2021, 3, 21, 21, 5, 5, 0, &time.Location{})
    for i := 0; i < b.N; i++ {
        zh := SetLocale(locale.ZHCN)
        zh.Format(date,
            "YY-YYYY-M-MM-MMM-MMMM-D-DD-d-dd-ddd-dddd-H-HH-h-hh-m-mm-s-ss-A-a")
    }
}

