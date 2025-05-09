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
