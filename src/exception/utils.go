package exception

func or(special string, dfault string) string {
	if special != "" {
		return special
	}
	return dfault
}
