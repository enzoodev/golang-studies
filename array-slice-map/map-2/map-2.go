package main

func main() {
	jobsAndSalaries := map[string]int{
		"developer": 4000,
		"designer":  2500,
		"tester":    2500,
		"architect": 4500,
		"manager":   6000,
		"director":  8000,
		"ceo":       12000,
	}

	for job, salary := range jobsAndSalaries {
		println(job, salary)
	}
}
