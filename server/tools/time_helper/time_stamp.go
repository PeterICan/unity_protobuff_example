package time_helper

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const TimeFormat = "2006-01-02 15:04:05.9999"
const MailTimeFormat = "2006-01-02T15:04"
const TaiPeiTimeZone = "Asia/Taipei"
const CodeTimeFormat = "2006-01-02"

func MakeTimestampMillisecond() int64 {
	// 獲取當前時間的 Unix timestamp，單位為納秒
	now := time.Now().UnixNano()
	// 將 Unix timestamp 的單位轉換為毫秒，並回傳數值
	return now / int64(time.Millisecond)
}

func MakeTimestampSecond() int64 {
	// 獲取當前時間的 Unix timestamp，單位為納秒
	now := time.Now().UnixNano()
	// 將 Unix timestamp 的單位轉換為秒，並回傳數值
	return now / int64(time.Second)
}

func MakeTimestampMicrosecond() int64 {
	// 獲取當前時間的 Unix timestamp，單位為納秒
	now := time.Now().UnixNano()
	// 將 Unix timestamp 的單位轉換為微秒，並回傳數值
	return now / int64(time.Microsecond)
}

func LocalTime() time.Time {
	// 取得當前時間
	now := time.Now()
	// 將當前時間轉換為本地時間
	t, _ := TimeIn(now, "Local")
	// 回傳轉換後的本地時間
	return t
}

// SubTimeIntervalBySeconds 以傳入的time為基礎，加上間隔秒數來判斷這
/*
	這是一個判斷時間差是否超過指定時間的函式。
	函式傳入兩個參數，第一個參數 baseTime 是一個 time.Time 型別的變數，表示要進行比較的時間基準；第二個參數 addSeconds 是一個 int 型別的變數，表示要增加的秒數。
	函式中首先使用 CopyTime() 函式複製傳入的時間作為結束時間 endTime，然後將結束時間加上指定的秒數，得到最終的到期時間。
	接著，計算現在時間與到期時間的差值，單位為秒，如果時間差小於等於 0，則表示現在時間還在指定時間內，回傳 false；反之，如果時間差大於 0，則表示現在時間已經超過指定時間，回傳 true。
*/
func SubTimeIntervalBySeconds(baseTime time.Time, addSeconds int) bool {
	endTime := CopyTime(baseTime)
	//傳入的時間加上addSeconds就是到期時間
	endTime = endTime.Add(time.Duration(addSeconds) * time.Second)
	sub := time.Now().Sub(endTime).Seconds()
	if sub <= 0 {
		//時間的誤差小於intervalSeconds
		return false
	}
	//傳入的時間加上
	return true
}

// TaipeiTime 指定時區為Taipei，方便之後在不同地區創建雲端伺服器時，維護人員仍以台北時間來做基準查詢
func TaipeiTime() time.Time {
	t, _ := TimeIn(time.Now(), TaiPeiTimeZone)
	return t
}

func TaipeiTimeDayDifferenceToNow(startTime int64) int64 {
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	start := time.Unix(startTime, 0).In(zone)
	startDate := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	now, _ := TimeIn(time.Now(), TaiPeiTimeZone)
	nowDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, start.Location())

	duration := nowDate.Sub(startDate)
	return int64(duration.Hours() / 24)
}

func TaipeiTimeDayDifference(startTime, targetTime int64) int64 {
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	start := time.Unix(startTime, 0).In(zone)
	startDate := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	target := time.Unix(targetTime, 0).In(zone)
	targetDate := time.Date(target.Year(), target.Month(), target.Day(), 0, 0, 0, 0, target.Location())

	duration := targetDate.Sub(startDate)
	return int64(duration.Hours() / 24)
}

func TaipeiDateTime() time.Time {
	t, _ := TimeIn(time.Now(), TaiPeiTimeZone)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func TaipeiDateUnix() int64 {
	t, _ := TimeIn(time.Now(), TaiPeiTimeZone)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

func TaipeiDateUnixWithCodeDisplay() string {
	t, _ := TimeIn(time.Now(), TaiPeiTimeZone)
	return t.Format(CodeTimeFormat)
}

func TaipeiNextDateUnix() int64 {
	t, _ := TimeIn(time.Now().AddDate(0, 0, 1), TaiPeiTimeZone)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

func TaipeiMondayUnix() int64 {
	t, _ := TimeIn(time.Now(), TaiPeiTimeZone)
	if t.Weekday() == time.Monday {
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	} else {
		offset := int(time.Monday - t.Weekday())
		if offset > 0 {
			offset -= 7
		}
		newTime, _ := TimeIn(time.Now().AddDate(0, 0, offset), TaiPeiTimeZone)

		return time.Date(newTime.Year(), newTime.Month(), newTime.Day(), 0, 0, 0, 0, t.Location()).Unix()
	}
}

func TaipeiNextMondayUnix() int64 {
	return TaipeiMondayUnix() + 7*24*60*60
}

func TaipeiMonthUnix() int64 {
	t, _ := TimeIn(time.Now(), TaiPeiTimeZone)
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).Unix()
}

func TaipeiNextMonthUnix() int64 {
	t, _ := TimeIn(time.Now().AddDate(0, 1, 0), TaiPeiTimeZone)
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).Unix()
}

func ChicagoTime() time.Time {
	t, _ := TimeIn(time.Now(), "America/Chicago")
	return t
}

func TaipeiDate() string {
	t, _ := TimeIn(time.Now(), TaiPeiTimeZone)
	return t.Format("2006-01-02 15:04:05")
}

func TaipeiDate2() string {
	t, _ := TimeIn(time.Now(), TaiPeiTimeZone)
	return t.Format("20060102_150405")
}

func FormatData(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// NextResetTime 明日的重置時間
func NextResetTime() time.Time {
	//resetTime := ResetTime()
	//resetTime = resetTime.Add(time.Hour * 24)
	return ResetTime().AddDate(0, 0, 1)
}

// ResetTime 重置時間
func ResetTime() time.Time {
	now := TaipeiTime()
	yyyy, mm, dd := now.Date()
	todayResetTime := time.Date(yyyy, mm, dd, 04, 0, 0, 0, now.Location())
	return todayResetTime
}

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func TestTimeIn() {
	for _, name := range []string{
		"",
		"Local",
		"Asia/Shanghai",
		TaiPeiTimeZone,
	} {
		t, err := TimeIn(time.Now(), name)
		if err == nil {
			fmt.Println(t.Location(), t.Format("15:04"))
		} else {
			fmt.Println(name, "<time unknown>")
		}
	}
}

//proto Timestamp 參考資料：https://blog.csdn.net/qq_32828933/article/details/105773544

func GetProtoLocalTime() *timestamppb.Timestamp {
	timeProto := timestamppb.Now()
	fmt.Println(timeProto)
	return timeProto
}

func ProtoTimeToGoTime(timeProto *timestamppb.Timestamp) time.Time {
	return timeProto.AsTime()
}

func GoTimeToProtoTime(timeGo time.Time) *timestamppb.Timestamp {
	return timestamppb.New(timeGo)
}

func ParseTimeByFixedTimeZone(dataTime time.Time) (time.Time, error) {
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	t, err := time.ParseInLocation(TimeFormat, dataTime.Format(TimeFormat), zone)
	return t, err
}

func ParseTimeByFixedTimeZoneFromStr(dateTimeStr string) (time.Time, error) {
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	t, err := time.ParseInLocation(TimeFormat, dateTimeStr, zone)
	return t, err
}

func ParseMailTimeByFixedTimeZone(dataTime time.Time) (time.Time, error) {
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	t, err := time.ParseInLocation(MailTimeFormat, dataTime.Format(TimeFormat), zone)
	return t, err
}

func ParseMailTimeByFixedTimeZoneFromStr(dateTimeStr string) (time.Time, error) {
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	t, err := time.ParseInLocation(MailTimeFormat, dateTimeStr, zone)
	return t, err
}

func ParseCodeTimeByFixedTimeZoneFromStrToTime(dateTimeStr string) (time.Time, error) {
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	t, err := time.ParseInLocation(CodeTimeFormat, dateTimeStr, zone)
	return t, err
}

func ParseCodeTimeByFixedTimeZoneFromStr(dateTimeStr string) int64 {
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	t, err := time.ParseInLocation(CodeTimeFormat, dateTimeStr, zone)
	if err != nil {
		return 0
	}
	return t.Unix()
}

func ParseUnixTimeToTaipeiTomeZoneMailDisplay(unixTime int64) string {
	if unixTime == 0 {
		return ""
	}
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	t := time.Unix(unixTime, 0).In(zone)
	return t.Format(MailTimeFormat)
}

func ParseUnixTimeToTaipeiTomeZoneCodeDisplay(unixTime int64) string {
	if unixTime == 0 {
		return ""
	}
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	t := time.Unix(unixTime, 0).In(zone)
	return t.Format(CodeTimeFormat)
}

func ParseUnixTimeToTaipeiTomeZoneDisplay(unixTime int64) string {
	if unixTime == 0 {
		return ""
	}
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	t := time.Unix(unixTime, 0).In(zone)
	return t.Format(TimeFormat)
}

func ParseUnixTimeToTaipeiTomeZone(unixTime int64) time.Time {
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	t := time.Unix(unixTime, 0).In(zone)
	return t
}

func ParseUnixTimeToTaipeiTomeZonePtr(unixTime int64) *time.Time {
	zone, _ := time.LoadLocation(TaiPeiTimeZone)
	t := time.Unix(unixTime, 0).In(zone)
	return &t
}

func CopyTime(t time.Time) time.Time {
	//newTime := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	//return newTime
	return t.Add(0) //chatGPT給的版本，試用看看
}

// 填刷新時間0~23
const DailyRefreshHour = 0

// IsBeforeDailyRefreshTime 檢查傳進來的時間，是否在下次每日刷新的時間之前
func IsBeforeDailyRefreshTime(lastDailyRefreshTime time.Time) bool {
	nextTime := GetTodayDailyRefreshTime()
	if lastDailyRefreshTime.Before(nextTime) {
		return true
	}
	return false
}

// GetTodayDailyRefreshTime 取得今日的每日刷新時間
func GetTodayDailyRefreshTime() time.Time {
	var date time.Time
	resetTime := TaipeiTime()
	loc, err := time.LoadLocation(TaiPeiTimeZone)
	if err != nil {
		return TaipeiTime()
	}
	date = time.Date(resetTime.Year(), resetTime.Month(), resetTime.Day(), DailyRefreshHour, 0, 0, 0, loc)
	return date
}

// GetNextDailyRefreshTime 以現在時間來判斷，下次的每日刷新時間為何
func GetNextDailyRefreshTime() time.Time {
	var date time.Time
	resetTime := TaipeiTime()
	loc, err := time.LoadLocation(TaiPeiTimeZone)
	if err != nil {
		return TaipeiTime()
	}

	//如果還沒超過N點，就是今天的零晨N點
	if resetTime.Hour() < DailyRefreshHour {
		date = time.Date(resetTime.Year(), resetTime.Month(), resetTime.Day(), DailyRefreshHour, 0, 0, 0, loc)
	} else { //如果超過N點，那更新時間是明天的N點
		date = time.Date(resetTime.Year(), resetTime.Month(), resetTime.Day()+1, DailyRefreshHour, 0, 0, 0, loc)
	}
	return date
}

// ParseCornTime 解析時間字串
func ParseCornTime(timeDate time.Time) string {
	cronTime := fmt.Sprintf("%v %v %v %d *", timeDate.Minute(), timeDate.Hour(), timeDate.Day(), timeDate.Month())
	return cronTime
}

// ParseTimeString 解析時間字串
func ParseTimeString(timeStr string) (time.Time, error) {
	//ex: timeStr := "2023-03-15 14:30:00"
	layout := "2006-01-02 15:04:05"
	location, err := time.LoadLocation(TaiPeiTimeZone)
	if err != nil {
		// 載入時區出錯，處理錯誤
		return time.Time{}, err
	}
	parseTime, err := time.ParseInLocation(layout, timeStr, location)
	if err != nil {
		// 轉換出錯，處理錯誤
		return time.Time{}, err
	}
	return parseTime, nil
}

// ParseTimeFormatAndTime 解析時間格式與時間字串
func ParseTimeFormatAndTime(timeFormat, timeStr string) (time.Time, error) {
	layout := timeFormat
	location, err := time.LoadLocation(TaiPeiTimeZone)
	if err != nil {
		// 載入時區出錯，處理錯誤
		return time.Time{}, err
	}
	parseTime, err := time.ParseInLocation(layout, timeStr, location)
	if err != nil {
		// 轉換出錯，處理錯誤
		return time.Time{}, err
	}
	return parseTime, nil
}

// GetCycleStartTimeStamp 取得循環起始點的時間戳記
func GetCycleStartTimeStamp(cycleStartHour int32) int64 {
	var date time.Time
	resetTime := TaipeiTime()
	loc, err := time.LoadLocation(TaiPeiTimeZone)
	if err != nil {
		return resetTime.Unix()
	}
	//以台北時間00:00:00為基準進行循環起始點
	date = time.Date(resetTime.Year(), resetTime.Month(), resetTime.Day(), int(cycleStartHour), 0, 0, 0, loc)
	return date.Unix()
}

// CheckIsCrossDay 檢查傳入的兩個時間，是否已跨日
func CheckIsCrossDay(baseTime, targetTime time.Time) bool {
	if baseTime.YearDay() != targetTime.YearDay() {
		fmt.Printf("CheckIsCrossDay baseTime:%v targetTime:%v is cross day\n", baseTime, targetTime)
		return true
	}

	if baseTime.Year() != targetTime.Year() {
		fmt.Printf("CheckIsCrossDay baseTime:%v targetTime:%v is cross day\n", baseTime, targetTime)
		return true
	}
	fmt.Printf("CheckIsCrossDay baseTime:%v targetTime:%v is same day\n", baseTime, targetTime)
	return false
}

// TimeToYearDayFormat 取得年份與年天數的格式化字串 "YY_ddd"
func TimeToYearDayFormat(time time.Time) string {
	yearDay := time.YearDay()
	year := time.Year() % 100
	return fmt.Sprintf("%02d_%d", year, yearDay)
}

func TimeToYYYMMDDFormat(time time.Time) string {
	return time.Format("2006-01-02")
}
