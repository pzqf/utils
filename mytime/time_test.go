package mytime

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	fmt.Println("GetNowSeconds:", GetNowSeconds())
	fmt.Println("GetNowNano:", GetNowNano())
	fmt.Println("GetToday0ClockSeconds:", GetToday0ClockSeconds())

	now := time.Now()
	fmt.Println("GetFormatYMDHMS:", GetFormatYMDHMS(now))
	fmt.Println("SecondsToTime:", SecondsToTime(GetNowSeconds()))
	fmt.Println("GetTodaySeconds:", GetTodaySeconds())

	fmt.Println("GetLastWeek0ClockSeconds:", GetLastWeek0ClockSeconds())
	fmt.Println(Time2String(SecondsToTime(GetLastWeek0ClockSeconds())))

	fmt.Println("GetLastMonth0ClockSeconds:", GetLastMonth0ClockSeconds())
	fmt.Println(Time2String(SecondsToTime(GetLastMonth0ClockSeconds())))

	fmt.Println("GetNextMonth0ClockSeconds:", GetNextMonth0ClockSeconds())
	fmt.Println(Time2String(SecondsToTime(GetNextMonth0ClockSeconds())))

	fmt.Println("GetThisMonthRestDays:", GetThisMonthRestDays())

	fmt.Println(time.Now().AddDate(1, 0, 0))
}