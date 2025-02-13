package timex

import "net/http"

// YearMonth is year month layout
const YearMonth = "2006-01"

// DateOnlyMonth
// Deprecated: Do not use. use DateOnlyMonth
const DateOnlyMonth = YearMonth

// MonthOnly
// Deprecated: Do not use. use DateOnlyMonth
const MonthOnly = YearMonth

const DateTimeMinute = "2006-01-02 15:04"

const DateTimeNano = "2006-01-02 15:04:05.999999999"

const UTCLayout = "2006-01-02T15:04:05.000Z"

const HttpTimeLayout = http.TimeFormat
