package examples

import "iter"

func (lst *List[T]) ShowAll() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := lst.head; i != nil; i = i.next {
			if !(yield(i.value)){
				return 
			}
		}
	}	
}

func GetFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		x, y := 0, 1

		for {
			if !(yield(x)){
				return
			}
			x, y = y, x + y
		}
	}
}