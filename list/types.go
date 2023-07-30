package list

type List[T any] interface {
	Get(index int) (T, error)
	Append(ts ...T) error
	Add(index int, t T) error
	Set(index int, t T) error
	Delete(index int) (T, error)
	Len() int
	Cap() int
	Range(fn func(index int, t T) error) error
	asSlice() []T
}
