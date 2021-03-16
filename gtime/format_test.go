package gtime

import (
    "github.com/stretchr/testify/assert"
    "github.com/xjh22222228/gosh/gtime/glocale"
    "strconv"
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

func BenchmarkFormat(b *testing.B) {
    date := time.Date(2021, 3, 21, 21, 5, 5, 0, &time.Location{})
    zh := SetLocale(glocale.ZHCN)

    for i := 0; i < b.N; i++ {
        zh.Format(date,
            "YY-YYYY-M-MM-MMM-MMMM-D-DD-d-dd-ddd-dddd-H-HH-h-hh-m-mm-s-ss-A-a")
    }
}
