package util

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func ParseString(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Trim(str, "\r\n")
	str = strings.TrimSpace(str)
	return str
}

func ParseInt64(str string) int64 {
	str = ParseString(str)
	i, err := strconv.Atoi(str)
	if err != nil {
		i = -1
	}
	return int64(i)
}

func ParseInt(str string) int {
	str = ParseString(str)
	i, err := strconv.Atoi(str)
	if err != nil {
		i = 0
	}
	return i
}

func ParseOffset(offset string) int {
	offset = ParseString(offset)
	i, err := strconv.Atoi(offset)
	if err != nil {
		i = 0
	}
	return i
}

func ParseLimit(limit string) int {
	limit = ParseString(limit)
	i, err := strconv.Atoi(limit)
	if err != nil {
		i = 10
	}
	return i
}

func ParseFloat64(str string) float64 {
	str = ParseString(str)
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		i = 0
	}
	return i
}

var (
	timezone = ""
)

func ParseTime(str string) time.Time {
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	time.Local = loc
	t, err := dateparse.ParseLocal(str)
	if err != nil {
		t = time.Now()
	}
	return t
}

func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		return s
	}
	return output
}

func UrlEncode(s string) string {
	res := url.QueryEscape(s)
	return res
}

func UrlDecode(s string) string {
	res, err := url.QueryUnescape(s)
	if err != nil {
		return s
	}
	return res
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func GetPageSize(pageSize string) int64 {
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		pageSizeInt = 50
	}
	return pageSizeInt
}

func CurrentTime() time.Time {
	return time.Now()
}

func CurrentTimeMicro() int64 {
	microTime := int64(time.Now().UnixNano() / 1000)
	return microTime
}

func TimeToString(valueTime time.Time) string {
	return TimeToStringLayout(valueTime, "2006-01-02 15:04:05")
}

func TimeToStringLayout(valueTime time.Time, layout string) string {
	return valueTime.Format(layout)
}

func ParseFromStringToTime(timeStr string) time.Time {
	return ParseFromStringToTimeLayout(timeStr, "2006-01-02 15:04:05")
}

func ParseFromStringToTimeLayout(timeStr string, layout string) time.Time {
	date, _ := time.Parse(layout, timeStr)
	return date
}

func CheckStartEndDate(startDate, endDate string) (time.Time, time.Time, error) {
	startTime := time.Now()
	endTime := time.Now()
	if startDate != "" {
		startTime = ParseFromStringToTime(startDate)
		if startTime.IsZero() {
			return time.Time{}, time.Time{}, errors.New("start_time is invalid")
		}
	} else {
		startDate = TimeToStringLayout(startTime, "2006-01-02") + " 00:00:00"
		startTime = ParseFromStringToTime(startDate)
	}
	if endDate != "" {
		endTime = ParseFromStringToTime(endDate)
		if endTime.IsZero() {
			return time.Time{}, time.Time{}, errors.New("end_time is invalid")
		}
	} else {
		endDate = TimeToStringLayout(startTime, "2006-01-02") + " 23:59:59"
		endTime = ParseFromStringToTime(endDate)
	}
	if startTime.After(endTime) {
		return time.Time{}, time.Time{}, errors.New("start_date must be after end_date")
	}
	return startTime, endTime, nil
}

func GetStartEndCurrent() (time.Time, time.Time) {
	current := time.Now()
	startTime := ParseFromStringToTime(TimeToStringLayout(current, "2006-01-02") + " 00:00:00")
	endTime := ParseFromStringToTime(TimeToStringLayout(current, "2006-01-02") + " 23:59:59")
	return startTime, endTime
}

func ParseQueryArray(slice []string) []string {
	result := make([]string, 0)
	for _, v := range slice {
		if len(v) > 0 {
			result = append(result, v)
		}
	}
	return result
}

func RemoveDuplicate(array []string) []string {
	m := make(map[string]string)
	for _, x := range array {
		m[x] = x
	}
	result := make([]string, 0)
	for x := range m {
		result = append(result, x)
	}
	return result
}
