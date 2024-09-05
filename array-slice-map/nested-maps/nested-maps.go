package main

func main() {
	employees := map[string]map[string]int{
		"developer": {
			"Alice": 4000, "Bob": 4500,
		},
		"designer":  {
			"Charlie": 2500, "David": 3000,
		},
		"tester":    {
			"Eve": 2500,
		},
		"architect": {
			"Frank": 4500,
		},
		"manager":   {
			"Grace": 6000,
		},
		"director":  {
			"Heidi": 8000,
		},
		"ceo":       {
			"Ivan": 12000,
		},
	}

	println(len(employees))

	for job, salaries := range employees {
		for name, salary := range salaries {
			println(job, name, salary)
		}
	}

	delete(employees, "tester")

	println(len(employees))
}
