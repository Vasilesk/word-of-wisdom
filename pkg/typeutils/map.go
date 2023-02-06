package typeutils

type Mapper[K comparable, V any] interface {
	Map() map[K]V
}

func NewMapper[K comparable, V any](m map[K]V) Mapper[K, V] {
	return mapHolder[K, V](m)
}

type mapHolder[K comparable, V any] map[K]V

func (mp mapHolder[K, V]) Map() map[K]V {
	return mp
}
