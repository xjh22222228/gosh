package glocale

type Locale struct {
    // 星期一|星期二...
    Weekdays [7]string
    // 一|二...
    WeekdaysMin [7]string
    // 周一|周二...
    WeekdaysShort [7]string
    // 一月|二月...
    Months [12]string
    // 1月|2月...
    MonthsShort [12]string

    // 上午|下午...
    GetTime func(hour int, minute int) string
}

type Language int

const (
    ZHCN Language = 1 + iota
    ZHHK
    ZHTW
    ENUS
)

func GetLocale(l Language) *Locale {
    switch l {
    case ENUS:
        return IntlLocaleEnUS
    case ZHCN:
        return IntlLocaleZhCN
    case ZHHK:
        return IntlLocaleEnHK
    case ZHTW:
        return IntlLocaleEnTW
    }
    return IntlLocaleEnUS
}