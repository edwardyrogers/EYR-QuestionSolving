package data_structures

type LinkedNode struct {
	Val  interface{}
	Next *LinkedNode
}

type LinkedList struct {
	Head *LinkedNode
}

func (std *LinkedList) Insert(val interface{}) {
	if std.Head == nil {
		std.Head = &LinkedNode{val, nil}

		return
	}

	if std.Head.Next == nil {
		std.Head.Next = &LinkedNode{val, nil}

		return
	}

	next := std.Head.Next

	for next.Next != nil {
		next = next.Next
	}

	next.Next = &LinkedNode{val, nil}
}

func (std *LinkedList) Retrieve(val interface{}) interface{} {
	var resultSet map[string]interface{}

	if std.Head != nil {
		index := 0

		if std.Head.Val == val {
			resultSet = map[string]interface{}{
				"idx": index,
				"val": std.Head.Val,
			}
		}

		if resultSet == nil {
			next := std.Head.Next

			for next != nil {
				index++

				if next.Val == val {
					resultSet = map[string]interface{}{
						"idx": index,
						"val": next.Val,
					}
				}

				next = next.Next
			}
		}
	}

	return resultSet
}

func (std *LinkedList) FilterByVal(val interface{}) interface{} {
	var resultSet []interface{}

	if std.Head != nil {
		index := 0

		if std.Head.Val == val {
			resultSet = append(resultSet, map[string]interface{}{
				"idx": index,
				"val": std.Head.Val,
			})
		}

		next := std.Head.Next

		for next != nil {
			index++

			if next.Val == val {
				resultSet = append(resultSet, map[string]interface{}{
					"idx": index,
					"val": next.Val,
				})
			}

			next = next.Next
		}

	}

	return resultSet
}

func (std *LinkedList) Update(oldVal interface{}, newVal interface{}) bool {
	result := false

	if std.Head != nil {

		if std.Head.Val == oldVal {
			std.Head.Val = newVal

			result = std.Head.Val == newVal
		}

		if !result {
			next := std.Head.Next

			for next != nil {
				if next.Val == oldVal {
					next.Val = newVal

					result = next.Val == newVal
					break
				}

				next = next.Next
			}
		}

	}

	return result
}

func (std *LinkedList) Delete(val interface{}) bool {
	result := false

	if std.Head != nil {

		if std.Head.Val == val {
			if std.Head.Next != nil {
				std.Head = std.Head.Next
			} else {
				std.Head = nil
			}

			result = true
		}

		if !result {
			prev := std.Head
			next := std.Head.Next

			for next != nil {

				if next.Val == val {
					if next.Next != nil {
						prev.Next = next.Next
					} else {
						prev.Next = nil
					}

					result = true
					break
				}

				prev = next
				next = next.Next
			}
		}
	}

	return result
}

func (std *LinkedList) Size() int {
	var count []interface{}

	if std.Head != nil {
		count = append(count, std.Head)

		next := std.Head.Next

		for next != nil {
			count = append(count, std.Head.Val)
			next = next.Next
		}
	}

	return len(count)
}

func (std *LinkedList) Reverse() {
	if std.Head == nil {
		return
	}

	var curr *LinkedNode = std.Head
	var prev *LinkedNode = nil
	var next *LinkedNode = nil

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	std.Head = prev
}
