package list

import (
	"errors"
	"github.com/schbm/gotastructs/general"
)

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
