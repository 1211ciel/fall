package numutil

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// RandomFloat2 保留2位小数点
func RandomFloat2(min, max float64) float64 {
	if max == min {
		return max
	}
	val, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64()*(max-min)+min), 64)
	return val
}
func RandomFloatStr(minStr, maxStr string) float64 {
	min, err := strconv.ParseFloat(minStr, 64)
	if err != nil {
		panic(err)
	}
	max, err := strconv.ParseFloat(maxStr, 64)
	if err != nil {
		panic(err)
	}
	return RandomFloat2(min, max)
}

// Decimal2 保留两位小数点
func Decimal2(v float64) float64 {
	val, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", v), 64)
	return val
}

func RandomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

func StrToFloat(s string) float64 {
	if float, err := strconv.ParseFloat(s, 10); err != nil {
		panic(err)
	} else {
		return float
	}
	return 0
}

func NumToInt64(num interface{}) (res int64, err error) {
	switch num.(type) {
	case string:
		res, err = strconv.ParseInt(num.(string), 10, 64)
	case int, int8, int16, int32, uint, uint8, uint16, uint32, uint64:
		res = int64(num.(int))
	}
	return
}

func ToDuration(num interface{}) (time.Duration, error) {
	toInt64, err := NumToInt64(num)
	if err != nil {
		return 0, errors.New("解析参数失败")
	}
	return time.Duration(toInt64), nil
}

// GenerateRandomNumber 生成count个[start,end)结束的不重复的随机数  count 为需要生成的个数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn(end-start) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

// GetAfterPointNum 获取小数点后面的数字
func GetAfterPointNum(num float64) (numList []uint8) {
	sprintf := fmt.Sprintf("%04f", num)
	split := strings.Split(sprintf, ".")
	numStrList := strings.Split(split[1], "")
	for _, item := range numStrList {
		parseInt, _ := strconv.ParseInt(item, 10, 64)
		numList = append(numList, uint8(parseInt))
	}
	return
}
func StrToIntWithoutErr(str string) int64 {
	parseInt, _ := strconv.ParseInt(str, 10, 60)
	return parseInt
}
