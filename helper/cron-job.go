package helper

import (
	"time"
	"fmt"
	"app/model/domain"
	"github.com/go-co-op/gocron"
)

func SetSchedule(schedule *domain.Schedule) {

	
    local, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		local = time.UTC
	}

    time := fmt.Sprintf("%s:%s",schedule.Hour,schedule.Minute)
    
    c := gocron.NewScheduler(local)
    c.Every(schedule.Day).Day().At(time).Do(func(){
        SendEmailController(schedule.Email, schedule.Subject, schedule.Body)
		fmt.Println("Email Sent")
    })
    
    c.StartAsync()
}