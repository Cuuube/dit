package coll

type ViewExec[T any] func(T)
type MapExec[T any] func(T) T
type FilterExec[T any] func(T) bool

const (
	itemTypeView   = 0
	itemTypeMap    = 1
	ItemTypeFilter = 2
)

// 执行单元
type execItem[T any] struct {
	typ        int
	viewExec   ViewExec[T]
	mapExec    MapExec[T]
	filterExec FilterExec[T]
}

func NewStream[T any](data []T) *Stream[T] {
	return &Stream[T]{
		rawData:   data,
		execQueue: make([]execItem[T], 0),
	}
}

// 执行流
type Stream[T any] struct {
	rawData   []T
	execQueue []execItem[T]
}

// 添加一个查看算子
func (s *Stream[T]) View(exec ViewExec[T]) *Stream[T] {
	item := execItem[T]{typ: itemTypeView, viewExec: exec}
	s.execQueue = append(s.execQueue, item)
	return s
}

// 添加一个映射算子
func (s *Stream[T]) Map(exec MapExec[T]) *Stream[T] {
	item := execItem[T]{typ: itemTypeMap, mapExec: exec}
	s.execQueue = append(s.execQueue, item)
	return s
}

// 添加一个过滤算子
func (s *Stream[T]) Filter(exec FilterExec[T]) *Stream[T] {
	item := execItem[T]{typ: ItemTypeFilter, filterExec: exec}
	s.execQueue = append(s.execQueue, item)
	return s
}

// 计算出数据
func (s *Stream[T]) Execute() []T {
	ret := s.rawData

	// 遍历算子
	for _, exec := range s.execQueue {
		layerResult := make([]T, 0)

		// 使用每个算子，计算新的这一层数据
		for i := range ret {
			cur := ret[i]

			switch exec.typ {
			case itemTypeView:
				exec.viewExec(cur)
			case itemTypeMap:
				cur = exec.mapExec(cur)
			case ItemTypeFilter:
				// 跳过插入
				if !exec.filterExec(cur) {
					continue
				}
			}
			layerResult = append(layerResult, cur)
		}

		ret = layerResult
	}
	return ret
}

// // 计算出数据
// func (s *Stream[T]) Execute() []T {
// 	ret := make([]T, 0)
// 	for i := range s.rawData {
// 		cur := s.rawData[i]
// 		for _, exec := range s.execQueue {
// 			switch exec.typ {
// 			case itemTypeMap:
// 				cur = exec.mapExec(cur)
// 				continue
// 			case ItemTypeFilter:
// 				if !exec.filterExec(cur) {
// 					// 跳过插入
// 					goto IGNORE
// 				}
// 			}
// 		}

// 		ret = append(ret, cur)
// 	IGNORE:
// 	}
// 	return ret
// }

func (s *Stream[T]) Export(dst *[]T) {
	ret := s.Execute()
	*dst = ret
}
