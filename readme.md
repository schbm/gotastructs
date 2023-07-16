# gotastructs

Gotastructs is a collection of Go datastructures.
They may be useful as reference or starting point.
This is a personal exercise and reference packet.
The stuctures are not thoroughly tested and if you
wish to use them concurrently additional steps are required.

## Containers

Every container implements the Iterable and Slicer interfaces. (in development)
As the name suggests they are able to return an iterator and a slice of their elements.

Currently, the following containers are implemented:
- DoublyLinkedList
- LinkedList
- ArrayList
- GeneralStack
- GeneralTree
- BinaryTree (in development
- BinarySearchTree (in development)
- Queue
- CircularQueue (in development)
- Fibonacci Heap (in development)
- Graph (in development)
- DirectedGraph (in development)

## Usage

Items or 'elements' of a datastructure have to implement the `type Element interface`.
Which currently consists of three further interfaces `type Stringer interface`, `type Comparer interface`
and `type Equaler interface`.

```go

// Stringer is an interface for types that can be converted to a string.
type Stringer interface {
    // ToString returns a string representation of the object.
    String() string
}

```

```go

// Comparer describes a type that can be compared to another type of the
// same specific type.
type Comparer interface {
    // Compare compares this Comparable to another type.
    // Returns 0 if they are equal, a positive number if this Comparable is
    // greater than the other, and a negative number if this
    // Comparable is less than the other Comparable.
    // For custom types that implement Comparable, this method should be
    // implemented in a way that makes sense for the type.
    // As this probably will be called a lot of times, it should be
    // implemented efficiently.
    // If two comparables are not of the same specific type, the return value
    // is undefined. It is recommended to return 0 in this case.
    Compare(any) int8
}

```

```go

type Equaler interface {
    // Equals compares this Comparables to another Comparables.
    // Returns true if they are equal, false otherwise.
    // Two Comparables are equal if they have the same type and value.
    // For custom types that implement Comparable, this method should be
    // implemented in a way that makes sense for the type.
    // As this probably will be called a lot of times, it should be
    // implemented efficiently.
    Equals(any) bool
}

```

### Example Element

As an example we will take a look at a simple int wrapper which implements the `Element` interface.
```go

package element

type WrappedInt struct {
    value int
}

func (w *WrappedInt) Value() int {
    return w.value
}

// implements IComparator
func (w *WrappedInt) Equals(other any) bool {
    v, ok := other.(*WrappedInt)
    
    if !ok {
        return false
    }
	
    return w.value == v.value
}

func (w *WrappedInt) String() string {
    return strconv.Itoa(w.value)
}

func (w *WrappedInt) Compare(other any) int8 {
if w.Equals(other) {
return 0
}

v, ok := other.(*WrappedInt)
if !ok {
return 0
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

## Filter, Map, Reduce

Currently the following functions are implemented:
- list.Filter
- (...) (in development)

```go

package list

func FilterList(filter func(element general.Element) bool, list List) error {
	if list == nil {
		return errors.New("error list is nil")
	}
	iterator := list.Iterator()
	filterIndexes := make([]int, 0, list.Size())

	index := 0
	for iterator.HasNext() {
		v := iterator.Next()
		if filter(v) {
			filterIndexes = append(filterIndexes, index)
		}
		index++
	}

	for i := len(filterIndexes) - 1; i >= 0; i-- {
		err := list.Remove(filterIndexes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

```

