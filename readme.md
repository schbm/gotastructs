# gotastructs

Gotastructs is a collection of Go datastructures.
> The stuctures are not thoroughly tested and if you
wish to use them concurrently additional steps are required.

## V2
In the new version i switched from using wrappers to generics.

## V1

Version 1 is deprecated and does not work anymore.

### Containers

Every container implements the Iterable and Slicer interfaces. (in development)
As the name suggests they are able to return an iterator and a slice of their elements.

Currently, the following containers are implemented:
- DoublyLinkedList
- LinkedList
- ArrayList
- Hashset
- GeneralStack
- GeneralTree
- BinaryTree (in development)
- BinarySearchTree (in development)
- Queue
- CircularQueue (in development)
- Fibonacci Heap (in development)
- Graph (in development)
- DirectedGraph (in development)
- CircularBuffer (in development)

### Usage

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

As an example we will take a look at a simple string wrapper which implements the `Element` interface.
```go

package element

import "strings"

type WrappedString string

func (w *WrappedString) Equals(other any) bool {
	v, ok := other.(*WrappedString)
	if !ok {
		return false
	}
	return *w == *v
}

func (w *WrappedString) String() string {
	return string(*w)
}

func (w *WrappedString) Compare(other any) int8 {
	if w.Equals(other) {
		return 0
	}

	v, ok := other.(*WrappedString)
	if !ok {
		return 0
	}

	return int8(strings.Compare(string(*w), string(*v)))
}

func NewString(value string) *WrappedString {
	v := WrappedString(value)
	return &v
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

FilterList(func(el general.Element) bool {
    v, ok := el.(*element.WrappedInt)
    if !ok {
        return true
    }
    if v.Value()%2 == 0 {
        return true
    }
    
    return false
}, list)

```

