package vec

func largerCapacity(capacity uint64) uint64 {
	return capacity + (capacity * 2) + 1
}

func convertToInt64(item interface{}) int64 {
	if val, ok := item.(int); ok {
		return int64(val)
	} else if val, ok := item.(int8); ok {
		return int64(val)
	} else if val, ok := item.(int16); ok {
		return int64(val)
	} else if val, ok := item.(int32); ok {
		return int64(val)
	} else {
		return item.(int64)
	}
}

func convertToFloat64(item interface{}) float64 {
	if val, ok := item.(float32); ok {
		return float64(val)
	} else {
		return item.(float64)
	}
}

func setToNumericType(v *Vec, item interface{}) {
	switch {
	case v.integer:
		v.data[v.Len()] = convertToInt64(item)
	case v.floating:
		v.data[v.Len()] = convertToFloat64(item)
	default:
		v.data[v.Len()] = item
	}
}
