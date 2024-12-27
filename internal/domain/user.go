package domain

import "math"

type User struct {
	Weight        float64
	LunchInterval float64
	OfficeHours   int
}

func (u *User) DailyWaterIntake() float64 {
	return u.Weight * 35.0
}

func (u *User) DailyWaterIntakeLiters() float64 {
	return u.DailyWaterIntake() / 1000.0
}

func (u *User) DailyWaterIntakeInGlasses(glassML int) float64 {
	return u.DailyWaterIntake() / float64(glassML)
}

func (u *User) DailyWaterIntakeInGlassesPerOfficeHours(glassML int) int {
	return int(math.Ceil((u.DailyWaterIntake() / float64(glassML)) / (float64(u.OfficeHours) - u.LunchInterval)))
}
