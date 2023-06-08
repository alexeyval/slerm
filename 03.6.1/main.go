package main

import "fmt"

func main() {
	fmt.Println(`Индетификаторы "pass":`)

	const (
		PassStatus = "pass"
		FailStatus = "fail"
	)

	type HealthCheck struct {
		ServiceID int
		status    string
	}

	GenerateCheck := func() []HealthCheck {
		checks := make([]HealthCheck, 5)
		for i := 0; i < 5; i++ {
			checks[i].ServiceID = i
			checks[i].status = FailStatus
			if i%2 == 0 {
				checks[i].status = PassStatus
			}
		}
		return checks
	}

	checks := GenerateCheck()
	for _, v := range checks {
		if v.status == PassStatus {
			fmt.Println(v.ServiceID)
		}
	}
}
