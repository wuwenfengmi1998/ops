package models

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"time"
)

func Unix_to_str_time(nanos int64) string {
	t := time.Unix(0, nanos)
	return (t.Format("2006-01-02 15:04:05.999999999")) // 输出：2021-06-10 09:49:59.123456789
}

// 获取当前时间字符串
// 参数格式可选，默认"2006-01-02 15:04:05"
func Get_current_time_string(format ...string) string {
	// 默认格式
	layout := "2006_01_02-15_04_05.999999999"

	// 如果传入了格式参数则使用自定义格式
	if len(format) > 0 {
		layout = format[0]
	}

	return time.Now().Format(layout)
}

func Time_date_str_to_time(timestr string) time.Time {

	// 定义与字符串匹配的布局（注意必须使用Go的参考时间格式）
	layout := "2006-01-02"

	// 解析时间
	t, err := time.Parse(layout, timestr)
	if err != nil {
		var notime time.Time

		return notime
	}

	return t

}

// 判断文件是否存在
func File_exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

func Is_email_valid(email string) bool {
	// 正则表达式（覆盖 99% 常见邮箱格式）
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

func Valid_str_len(s string, min, max int, checkRune bool) bool {
	var length int
	if checkRune {
		length = len([]rune(s))
	} else {
		length = len(s)
	}

	if min < 0 || max < 0 || (max != 0 && max < min) {
		return false // 参数非法
	}

	return length >= min && (max == 0 || length <= max)

	// 示例
	//valid1 := Valid_str_len("hello", 3, 20, false) // 校验字节数
	//valid2 := Valid_str_len("你好", 1, 2, true)    // 校验字符数
}

func Rand_str_32() string {
	// 生成 32 字节 (256 位) 随机数据
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	// 转换为 16 进制字符串 (长度 64)
	cookie := hex.EncodeToString(b)
	return cookie
}

func Md5_str(str string) string {
	hashBytes2 := md5.Sum([]byte(str))
	hashString2 := hex.EncodeToString(hashBytes2[:]) // 注意数组转切片的[:]
	return hashString2
}

func Hash_user_pass(str string) string {
	switch Configs_user.Pass_hash_type {
	case "text":
		return str
	case "md5":
		return Md5_str(str)
	}

	return Get_current_time_string() + Rand_str_32() //如果转换失败返回当前时间，避免撞库
}

func Page_range(start_page int64, end_page int64, now_page int64, href_heard string) []map[string]string {
	var a []map[string]string

	for i := start_page; i <= end_page; i++ {
		var active = ""
		if i == now_page {
			active = "active"
		} else {
			active = ""
		}
		a = append(a, map[string]string{
			"page_href": fmt.Sprintf("%s%d", href_heard, i),
			"active":    active,
			"page":      fmt.Sprintf("%d", i),
		})
	}

	return a
}
