package color

import (
	"fmt"
	"strings"
)

//绿色字体
func Green(str string, modifier ...interface{}) string {
	return cliColorRender(str, 32, 0)
}

//淡绿
func LightGreen(str string, modifier ...interface{}) string {
	return cliColorRender(str, 32, 1)
}

//青色/蓝绿色
func Cyan(str string, modifier ...interface{}) string {
	return cliColorRender(str, 36, 0)
}

//淡青色
func LightCyan(str string, modifier ...interface{}) string {
	return cliColorRender(str, 36, 1)
}

//红字体
func Red(str string, modifier ...interface{}) string {
	return cliColorRender(str, 31, 0)
}

//淡红色
func LightRed(str string, modifier ...interface{}) string {
	return cliColorRender(str, 31, 1)
}

//黄色字体
func Yellow(str string, modifier ...interface{}) string {
	return cliColorRender(str, 33, 0)
}

//黑色
func Black(str string, modifier ...interface{}) string {
	return cliColorRender(str, 30, 0)
}

//深灰色
func DarkGray(str string, modifier ...interface{}) string {
	return cliColorRender(str, 30, 1)
}

//浅灰色
func LightGray(str string, modifier ...interface{}) string {
	return cliColorRender(str, 37, 0)
}

//白色
func White(str string, modifier ...interface{}) string {
	return cliColorRender(str, 37, 1)
}

//蓝色
func Blue(str string, modifier ...interface{}) string {
	return cliColorRender(str, 34, 0)
}

//淡蓝
func LightBlue(str string, modifier ...interface{}) string {
	return cliColorRender(str, 34, 1)
}

//紫色
func Purple(str string, modifier ...interface{}) string {
	return cliColorRender(str, 35, 0)
}

//淡紫色
func LightPurple(str string, modifier ...interface{}) string {
	return cliColorRender(str, 35, 1)
}

//棕色
func Brown(str string, modifier ...interface{}) string {
	return cliColorRender(str, 33, 0)
}

func cliColorRender(str string, color int, weight int) string {
	var mo []string

	if weight > 0 {
		mo = append(mo, fmt.Sprintf("%d", weight))
	}
	if len(mo) <= 0 {
		mo = append(mo, "0")
	}
	return fmt.Sprintf("\033[%s;%dm"+str+"\033[0m", strings.Join(mo, ";"), color)
}
