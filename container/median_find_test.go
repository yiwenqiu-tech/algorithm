package container

import "testing"

func TestMedianHolder_getMedian(t *testing.T) {
	maxSize := 1000
	maxItem := 100000
	loopTime := 10000
	for i := 0; i < loopTime; i++ {
		input := getRandomArray(maxSize, maxItem)
		medianByExhaustive := getMedianByExhaustive(input)
		var medianHolder MedianHolder
		for index := range input {
			medianHolder.addNumber(input[index])
		}
		medianByHolder := medianHolder.getMedian()
		if medianByExhaustive != medianByHolder {
			t.Errorf("getMedianByExhaustive() = %v, medianHolder %v, array = %+v",
				medianByExhaustive, medianByHolder, input)
		}
	}
}
