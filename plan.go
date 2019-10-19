package furrytail

import (
	"errors"
	"strconv"
	"time"
)

//GetPlan returns feeding plan
func (a *Account) GetPlan(deviceID int) (plan *Plan, err error) {
	res, err := a.client.R().SetPathParams(map[string]string{
		"deviceId": strconv.Itoa(deviceID),
	}).SetResult(&planResponse{}).Get("/provider-feeder/device/feed/plan?deviceId={deviceId}")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(*apiError).Message)
	}
	data := res.Result().(*planResponse).Data
	var items []PlanItem
	for _, item := range data.PlanItems {
		hour, minute := convertToTime(item.Time, a.loc)
		items = append(items, PlanItem{
			Name: item.Name,
			Copies: item.Copies,
			Hour: hour,
			Minute: minute,
		})
	}
	return &Plan{
		Items: items,
		Days: convertDays(data.Repeat),
	}, nil
}

//GetNearestFeedings returns previous and next feeding times according to the plan
func (a *Account) GetNearestFeedings(plan Plan) (previous, next *time.Time) {
	if len(plan.Days) == 0 || len(plan.Items) == 0 {
		return
	}
	var feedings []time.Time
	now := time.Now().In(&a.loc)
	start := now.Add(-8*time.Hour*24)
	for i := 0; i <= 16; i++ {
		t := start.Add(time.Duration(i)*time.Hour*24)
		if hasWeekday(t.Weekday(), plan.Days) {
			for _, item := range plan.Items {
				feedings = append(feedings,
					time.Date(t.Year(), t.Month(), t.Day(), item.Hour, item.Minute, 0, 0, &a.loc))
			}
		}
	}
	for _, feeding := range feedings {
		cloned := feeding
		diff := now.Sub(feeding)
		if diff < 0 {
			if next == nil || now.Sub(*next) < diff {
				next = &cloned
			}
		}
		if diff > 0 {
			if previous == nil || now.Sub(*previous) > diff {
				previous = &cloned
			}
		}
	}
	return
}

func convertDays(repeat string) (days []time.Weekday) {
	for i, enabled := range repeat {
		day := i+1
		if day == 7 {
			day = 0
		}
		if string(enabled) == "1" {
			days = append(days, time.Weekday(day))
		}
	}
	return
}

type Plan struct {
	Items []PlanItem
	Days []time.Weekday
}

type PlanItem struct {
	Name string
	Copies int
	Hour int
	Minute int
}

type planResponse struct {
	apiResponse
	Data struct {
		PlanItems []struct {
			Copies int `json:"copies"`
			Name string `json:"name"`
			Time int `json:"time"`
		} `json:"planItems"`
		Repeat string `json:"repeat"`
	} `json:"data"`
}