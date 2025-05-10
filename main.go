package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func DateDifference(date time.Time) (int, error) {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		log.Fatalf("%s خطا در بارگذاری منطقه زمانی:", err)
	}
	date = date.In(loc)
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, loc)
	now := time.Now().In(loc)
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	diff := now.Sub(date)
	days := int(diff.Hours() / 24)
	return days, nil
}

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
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	// کم کردن ۶۰ روز از تاریخ جاری
	pastDate := now.AddDate(0, 0, -60)

	// فرمت کردن تاریخ به فرمت YYYY-MM-DD
	return pastDate.Format("2006-01-02")
}

func GetLockDate() string {
	// بارگذاری منطقه زمانی Asia/Tehran
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		fmt.Println("خطا در بارگذاری منطقه زمانی:", err)
		return ""
	}

	// دریافت تاریخ و زمان جاری در منطقه زمانی Asia/Tehran
	now := time.Now().In(loc)

	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	// کم کردن ۶۰ روز از تاریخ جاری
	pastDate := now.AddDate(0, 0, -30)

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
		exFlag   = flag.Bool("ex", false, "Date in YYYY/MM/DD or YYYY-MM-DD format")
	)
	flag.Parse()
	if *dateFlag != "" {
		newDate, errPars := parseDate(*dateFlag)
		if errPars != nil {
			log.Fatalf("%s خطا در تبدیل تاریخ:", errPars)
		}
		diff, err := DateDifference(newDate)
		if err != nil {
			log.Fatalf("%s خطا در محاسبه:", err)
		}
		diff = 60 - diff
		if diff <= 0 {
			log.Printf("Today")
		} else {
			log.Printf("After %d day(s)", diff)
		}
	} else if *exFlag {
		log.Printf("Lock date for today is %s", GetLockDate())
	} else {
		fmt.Printf("Release Date for today is: %s\n", GetReleaseDate())
	}
}
