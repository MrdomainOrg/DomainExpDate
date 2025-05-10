package irnicDate

import (
	"fmt"
	"time"
)

func parseDate(dateStr string) (time.Time, error) {
	// تعریف لیستی از قالب‌های زمانی مورد انتظار
	formats := []string{
		"2006-01-02", "2006-01-2", "2006-1-02", "2006-1-2", "2006/01/02", "2006/01/2", "2006/1/02", "2006/1/2",
	}

	// تلاش برای تجزیه تاریخ با هر یک از قالب‌ها
	for _, format := range formats {
		t, err := time.Parse(format, dateStr)
		if err == nil {
			return t, nil
		}
	}

	// در صورت عدم تطابق با هیچ قالبی، خطا برمی‌گرداند
	return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
}

func GetReleaseDate() string {
	// بارگذاری منطقه زمانی Asia/Tehran
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		fmt.Println("خطا در بارگذاری منطقه زمانی:", err)
		return ""
	}

	// دریافت تاریخ و زمان جاری در منطقه زمانی Asia/Tehran
	now := time.Now().In(loc)

	// کم کردن ۶۰ روز از تاریخ جاری
	pastDate := now.AddDate(0, 0, -60)

	// فرمت کردن تاریخ به فرمت YYYY-MM-DD
	return pastDate.Format("2006-01-02")
}
