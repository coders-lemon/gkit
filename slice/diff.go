package slice

import "errors"

type Number interface {
	~byte | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Max[T Number](vals []T) (T, error) {
	if len(vals) == 0 {
		var v T
		return v, errors.New("")
	}
	res := vals[0]
	for i := 1; i < len(vals); i++ {
		if res < vals[i] {
			res = vals[i]
		}
	}
	return res, nil
}
