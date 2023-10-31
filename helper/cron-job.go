package helper

import (
	"time"
	"fmt"
	"strconv"
	"app/model/domain"
	"github.com/go-co-op/gocron"
)

func SetSchedule(schedule *domain.Schedule, medicine domain.Medicine) {

	
    local, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		local = time.UTC
	}

    time := fmt.Sprintf("%s:%s",schedule.Hour,schedule.Minute)
    day, _ := strconv.Atoi(schedule.Day)

    c := gocron.NewScheduler(local)
    c.Every(day).Day().At(time).Do(func(){
        SendEmailController(schedule.Email, schedule.Subject, medicine)
		fmt.Println("Email Sent")
    })
    
    c.StartAsync()
}