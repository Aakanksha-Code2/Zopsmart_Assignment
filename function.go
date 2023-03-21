package patient

func validatename(name string) bool {
	if name == "" {
		return name != ""
	}

	return true
}

func validId(id int) bool {
	return id > 0
}

