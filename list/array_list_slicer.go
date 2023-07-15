package list

import "github.com/schbm/gotastructs"

func (list *ArrayList) ToSlice() []gotastructs.Element {
	return list.elements
}
