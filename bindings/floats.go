package bindings

import (
	"fmt"
	"math"
)

func jsonToFloat(j interface{}) float32 {
	switch j.(type) {
	case string:
		if j == "NaN" {
			return float32(math.NaN())
		} else if j == "Infinity" {
			return float32(math.Inf(1))
		} else if j == "-Infinity" {
			return float32(math.Inf(-1))
		}
	case float64:
		return float32(j.(float64))
	default:
		panic(fmt.Sprintf("unexepcted type %T: %v", j, j))
		return 0
	}

	return 0
}

func jsonToDoubl(j interface{}) float64 {
	switch j.(type) {
	case string:
		if j == "NaN" {
			return math.NaN()
		} else if j == "Infinity" {
			return math.Inf(1)
		} else if j == "-Infinity" {
			return math.Inf(-1)
		}
	case float64:
		return j.(float64)
	default:
		panic(fmt.Sprintf("unexepcted type %T: %v", j, j))
		return 0
	}

	return 0
}

func jsonToFloats(j []interface{}) []float32 {
	ret := make([]float32, len(j))
	for i, inter := range j {
		switch inter.(type) {
		case string:
			if inter == "NaN" {
				ret[i] = float32(math.NaN())
			} else if inter == "Infinity" {
				ret[i] = float32(math.Inf(1))
			} else if inter == "-Infinity" {
				ret[i] = float32(math.Inf(-1))
			}
		case float64:
			ret[i] = float32(inter.(float64))
		default:
			panic(fmt.Sprintf("unexepcted type %T: %v", inter, inter))
		}
	}

	return ret
}

func jsonToDoubls(j []interface{}) []float64 {
	ret := make([]float64, len(j))
	for i, inter := range j {
		switch inter.(type) {
		case string:
			if inter == "NaN" {
				ret[i] = math.NaN()
			} else if inter == "Infinity" {
				ret[i] = math.Inf(1)
			} else if inter == "-Infinity" {
				ret[i] = math.Inf(-1)
			}
		case float64:
			ret[i] = inter.(float64)
		default:
			panic(fmt.Sprintf("unexepcted type %T: %v", inter, inter))
		}
	}

	return ret
}
