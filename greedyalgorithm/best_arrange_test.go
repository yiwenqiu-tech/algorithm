package greedyalgorithm

import "testing"

func TestBestArrange(t *testing.T) {
	programSize := 20
	timeMax := 24 * 3600
	loopTime := 10000
	for i := 0; i < loopTime; i++ {
		ps := generatePrograms(programSize, timeMax)
		if BestArrangeByExhaustive(ps) != BestArrange(ps) {
			t.Errorf("BestArrangeByExhaustive() = %v, BestArrange %v",
				BestArrangeByExhaustive(ps), BestArrange(ps))
		}
	}
}
