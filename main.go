package main

import (
	"flag"
	"fmt"
	"os"
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

type DateFlag struct {
	*time.Time
}

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
	}
	var (
		dateFlag = flag.String("d", "", "Date in YYYY/MM/DD or YYYY-MM-DD format")
	)
	flag.Parse()
	if *dateFlag != "" {

	} else {
		fmt.Printf("Release Date for today is: %s\n", GetReleaseDate())
	}
}
