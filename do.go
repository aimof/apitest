package apitest

// DoTasks to test.
func DoTasks(tasks Tasks, kicker Kicker, selector CompSelector) ([]Result, error) {
	results := make([]Result, 0, len(tasks))
	for _, task := range tasks {

		r, err := do(task, kicker, selector)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, nil
}

func do(task Task, kicker Kicker, selector CompSelector) (Result, error) {
	got, err := kicker.Kick(task.Request, task.Retry)
	if err != nil {
		return Result{}, err
	}

	comparer := selector.Select(task.CompMode)
	match := comparer.Compare(got, task.Want)
	return Result{
		Name:  task.Name,
		Got:   got,
		Match: match,
	}, nil
}
