package orderedmap

type (
	Elementer interface {
		Next() Elementer
		Prev() Elementer
	}

	Lister interface {
		Init() Lister
		Len() int
		Front() Elementer
		Back() Elementer
		Remove(e Elementer) interface{}
		PushFront(v interface{}) Elementer
		PushBack(v interface{}) Elementer
		InsertBefore(v interface{}, mark Elementer) Elementer
		InsertAfter(v interface{}, mark Elementer) Elementer
		MoveToFront(e Elementer)
		MoveToBack(e Elementer)
		MoveBefore(e, mark Elementer)
		MoveAfter(e, mark Elementer)
		PushBackList(other Lister)
		PushFrontList(other Lister)
	}
)
