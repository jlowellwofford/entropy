package entropy

func GetEntCnt() (int, error) {
	return getEntCnt()
}

func AddToEntCnt(add int) error {
	return addToEntCnt(add)
}

func AddEntropy(cnt int, buf []byte) error {
	return addEntropy(cnt, buf)
}

func ZapEntCnt() error {
	return zapEntCnt()
}

func ClearPool() error {
	return clearPool()
}

func ReseedCrng() error {
	return reseedCrng()
}
