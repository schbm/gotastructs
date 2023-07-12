package gotastructs

import "errors"

func FilterList(filter func(element Element) bool, list List) error {
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

	for _, i := range filterIndexes {
		err := list.Remove(i)
		if err != nil {
			return err
		}
	}

	return nil
}
