package chemistry

type flask map[string]int64

func newFlask() flask {
	return flask{}
}

func (f flask) Add(chemical string, amount int64) {
	f[chemical] = f[chemical] + amount
}

func (f flask) ContainsOnly(chemical string) bool {
	if _, contains := f[chemical]; contains {
		return len(f) == 1
	}
	return len(f) == 0
}
