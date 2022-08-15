package coll

type DictKeyable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | bool | uintptr
}

type Dict[keyType DictKeyable, valType any] map[keyType]valType

func (dict *Dict[keyType, valType]) Get(k keyType) valType {
	return (*dict)[k]
}

func (dict *Dict[keyType, valType]) GetOrDefault(k keyType, defaultValue valType) valType {
	v, found := (*dict)[k]
	if !found {
		return defaultValue
	}
	return v
}

func (dict *Dict[keyType, valType]) Set(k keyType, v valType) {
	(*dict)[k] = v
}

func (dict *Dict[keyType, valType]) Contains(k keyType) bool {
	_, found := (*dict)[k]
	return found
}

func (dict *Dict[keyType, valType]) Len() int {
	return len(*dict)
}

func (dict *Dict[keyType, valType]) Keys() []keyType {
	ret := make([]keyType, dict.Len())
	i := 0
	for k := range *dict {
		copyK := k
		ret[i] = copyK
		i++
	}
	return ret
}

func (dict *Dict[keyType, valType]) Values() []valType {
	ret := make([]valType, dict.Len())
	i := 0
	for k := range *dict {
		ret[i] = dict.Get(k)
		i++
	}
	return ret
}
