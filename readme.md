# gotastructs

Godastructs is a collection of Go datastructures.
They may be useful as reference or starting point.
This is a personal exercise and reference packet.
The stuctures are not thoroughly tested and if you
wish to use them concurrently additional steps are required.

## Usage

Item or 'elements' of a datastructure have to implement the `type Element interface`.
Which currently consists of two further interfaces `type Stringer interface` and `type Comparable interface`.
```go

// Equals compares this Comparables to another Comparables.
// Returns true if they are equal, false otherwise.
// Two Comparables are equal if they have the same type and value.
// For custom types that implement Comparable, this method should be
// implemented in a way that makes sense for the type.
// As this probably will be called a lot of times, it should be
// implemented efficiently.
Equals(Comparable) bool

```

```go
// Compare compares this Comparable to another Comparable.
// Returns 0 if they are equal, a positive number if this Comparable is
// greater than the other Comparable, and a negative number if this
// Comparable is less than the other Comparable.
// For custom types that implement Comparable, this method should be
// implemented in a way that makes sense for the type.
// As this probably will be called a lot of times, it should be
// implemented efficiently.
// If two comparables are not of the same specific type, the return value
// is undefined. It is recommended to return 0 in this case.
Compare(Comparable) int8 
```

```go
// ToString returns a string representation of the object. 
String() string

```

### Example Element

As an example we will take a look at a simple int wrapper which implements the `Element` interface.
```go

type WrappedInt struct {
	value int
}

// implements IComparator
func (w *WrappedInt) Equals(other Comparable) bool {
	v, ok := other.(*WrappedInt)
	if !ok {
		return false
	}
	return w.value == v.value
}

func (w *WrappedInt) String() string {
	return strconv.Itoa(w.value)
}

func (w *WrappedInt) Compare(other Comparable) int8 {
	if w.Equals(other) {
		return 0
	}

	v, ok := other.(*WrappedInt)
	if !ok {
		return -1
	}

	if w.value > v.value {
		return 1
	}
	return -1
}

func NewInt(value int) *WrappedInt {
	return &WrappedInt{
		value: value,
	}
}


```

I am pretty sure there are better ways to do this, but this is the way I did it.
(I am not a big fan of the boilerplate code, but I guess it is necessary in this case.)
I have yet to find a documentation which shows the performance impact and discusses the
differences between using interfaces, interface{}, specific types and generics in Go.

## Containers

Every container implements the Iterable and Slicer interfaces.
As the name suggests they are able to return an iterator and a slice of their elements.

Currently the following containers are implemented:
- DoublyLinkedList
- LinkedList
- ArrayList
- GeneralStack (in development)
- GeneralTree (in development)
- GeneralBinaryTree (in development
- BinarySearchTree (in development)
- GeneralQueue (in development)
- CircularQueue (in development)
- Fibonacci Heap (in development)

## Filter, Map, Reduce

Currently the following functions are implemented:
- FilterList
- (...)

```go
FilterList(func(el Element) bool {
		v, ok := el.(*WrappedInt)
		if !ok {
			return true
		}
		if v.value%2 == 0 {
			return true
		}

		return false
	}, list)
```

