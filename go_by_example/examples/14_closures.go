package examples

func Counter() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
