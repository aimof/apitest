package apitest

func Do(s Scenario, kicker Kicker, comp Comparer) (bool, error) {
	for _, test := range s.Tests {
		resp, err := kicker.Kick(test.When)
		if err != nil {
			return false, err
		}
		match, err := comp.Match(resp, test.Then)
		resp.Body.Close()
		if err != nil {
			return false, err
		}
		if !match {
			return false, nil
		}
	}
	return true, nil
}
