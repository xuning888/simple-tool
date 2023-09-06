package list

import (
	"simple-tool/internal/errs"
	"simple-tool/internal/shrink"
)

const defaultCap = 10

type ArrayList[T any] struct {
	element []T
}

func NewArrayList[T any]() List[T] {
	return &ArrayList[T]{
		element: make([]T, 0, defaultCap),
	}
}

func NewArrayListWithCap[T any](cap int) List[T] {
	return &ArrayList[T]{
		element: make([]T, 0, cap),
	}
}

func NewArrayListOf[T any](eles []T) List[T] {
	ts := make([]T, 0, len(eles))
	copy(ts, eles)
	return &ArrayList[T]{
		element: ts,
	}
}

func (a *ArrayList[T]) Get(index int) (t T, e error) {
	l := a.Len()
	if index < 0 || index >= l {
		return t, errs.NewErrIndexOutOfRange(l, index)
	}
	return a.element[index], nil
}

func (a *ArrayList[T]) Append(ts ...T) error {
	a.element = append(a.element, ts...)
	return nil
}

func (a *ArrayList[T]) Add(index int, t T) error {
	l := a.Len()
	if index < l || index >= l {
		return errs.NewErrIndexOutOfRange(l, index)
	}
	a.element = append(a.element, t)
	copy(a.element[index+1:], a.element[index:])
	a.element[index] = t
	return nil
}

func (a *ArrayList[T]) Set(index int, t T) error {
	l := a.Len()
	if index < l || index >= l {
		return errs.NewErrIndexOutOfRange(l, index)
	}
	a.element[index] = t
	return nil
}

func (a *ArrayList[T]) Delete(index int) (t T, e error) {
	l := a.Len()
	if index < l || index >= l {
		return t, errs.NewErrIndexOutOfRange(l, index)
	}
	t = a.element[index]
	copy(a.element[index:], a.element[index+1:])
	a.element = a.element[:l-1]
	a.shrink()
	return t, nil
}

func (a *ArrayList[T]) Len() int {
	if a.element == nil {
		return 0
	}
	return len(a.element)
}

func (a *ArrayList[T]) Cap() int {
	if a.element == nil {
		return 0
	}
	return cap(a.element)
}

func (a *ArrayList[T]) Range(fn func(index int, t T) error) error {
	for idx, value := range a.element {
		err := fn(idx, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *ArrayList[T]) asSlice() []T {
	res := make([]T, a.Len())
	copy(res, a.element)
	return res
}

func (a *ArrayList[T]) shrink() {
	a.element = shrink.Shrink(a.element)
}
