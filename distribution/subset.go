package distribution

import "math/rand"

func SubSet(backends []string, clientId int, subSetSize int) []string {
	subSetCount := len(backends) / subSetSize

	// Group clients into, each round use the same shuffle list
	round := clientId / subSetCount
	r := rand.New(rand.NewSource(int64(round)))
	r.Shuffle(len(backends), func(i, j int) {
		backends[i], backends[j] = backends[j], backends[i]
	})

	// the subset id corresponding to the current client
	subSetId := clientId % subSetCount

	start := subSetId * subSetSize
	return backends[start : start+subSetSize]
}
