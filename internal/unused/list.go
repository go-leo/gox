package unused

// Node 结构体表示链表中的一个节点。
type Node struct {
	// 节点的值。
	Val int
	// 下一个节点的指针。
	Next *Node
}

// ReverseList 反转一个单链表。
// 参数 head 是指向链表头部的指针。
// 返回值是反转后链表的头节点指针。
func ReverseList(head *Node) *Node {
	// 初始化前一个节点指针为nil，当前节点指针从链表头开始。
	var prev *Node

	// 当前节点指针从链表头开始。
	cur := head
	// 遍历链表直到当前节点指针为nil。
	for cur != nil {
		// 保存当前节点的下一个节点。
		next := cur.Next
		// 将当前节点的Next指向前一个节点，实现反转。
		cur.Next = prev
		// 更新前一个节点和当前节点指针。
		prev = cur
		cur = next
	}

	// 返回反转后的链表头节点指针。
	return prev
}

// ReverseBetween 反转单链表中从第m个节点到第n个节点的部分。
// head是链表的头节点，m和n是需要反转的链表部分的起始和结束位置。
// 返回反转后的链表的头节点。
func ReverseBetween(head *Node, m int, n int) *Node {
	// 创建一个虚拟头节点dummy，简化边界情况处理。
	dummy := &Node{Next: head}
	// prev用于找到第m-1个节点，初始化为dummy。
	prev := dummy
	// 移动prev到第m-1个节点。
	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}
	// cur用于找到第m个节点。
	cur := prev.Next
	// 从第m个节点到第n个节点进行反转。
	for i := 0; i < n-m; i++ {
		// next暂存当前节点的下一个节点。
		next := cur.Next
		// 更新cur的Next指针，指向next的下一个节点，即cur的下一个节点变为next的下一个节点。
		cur.Next = next.Next
		// 将next的Next指针指向prev的下一个节点，即把next节点插入到prev后面。
		next.Next = prev.Next
		// 更新prev的Next指针为next，完成节点的反转。
		prev.Next = next
	}
	// 返回反转后的链表的头节点。
	return dummy.Next
}

// _Merge 合并两个已排序的链表为一个已排序的链表。（递归）
// pHead1 和 pHead2 分别是两个已排序链表的头节点。
// 该函数返回合并后的已排序链表的头节点。
func _Merge(pHead1 *Node, pHead2 *Node) *Node {
	// 如果其中一个链表为空，则直接返回另一个链表。
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	// 比较两个链表当前节点的值。
	// 值较小的节点成为合并链表的当前节点。
	// 然后递归地将较小值节点的下一个节点与另一个链表进行合并。
	if pHead1.Val < pHead2.Val {
		pHead1.Next = _Merge(pHead1.Next, pHead2)
		return pHead1
	} else {
		pHead2.Next = _Merge(pHead1, pHead2.Next)
		return pHead2
	}
}

// Merge 合并两个已排序的链表为一个已排序的链表。（迭代）
// pHead1 和 pHead2 分别是两个已排序链表的头节点。
// 该函数返回合并后的已排序链表的头节点。
func Merge(pHead1 *Node, pHead2 *Node) *Node {
	// 创建一个虚拟节点作为合并链表的起始点。
	dummy := &Node{}
	// tail 指向当前合并链表的最后一个节点。
	tail := dummy

	// 遍历两个链表，直到其中一个链表完全合并。
	for pHead1 != nil && pHead2 != nil {
		// 比较两个链表的当前节点，将值较小的节点添加到合并链表中。
		if pHead1.Val < pHead2.Val {
			tail.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			tail.Next = pHead2
			pHead2 = pHead2.Next
		}
		// 将 tail 指针移动到合并链表的最后一个节点。
		tail = tail.Next
	}

	// 如果 list pHead1 中还有剩余节点，直接将它们追加到合并链表中。
	if pHead1 != nil {
		tail.Next = pHead1
	} else {
		// 如果 list pHead2 中还有剩余节点，直接将它们追加到合并链表中。
		tail.Next = pHead2
	}

	// 返回合并链表的头节点，dummy.Next 指向合并链表的第一个节点。
	return dummy.Next
}

// MergeKLists 合并多个排序链表。
// 参数 lists 是一个 Node 类型链表的切片。
// 返回值是合并后的单链表头节点。
// 该函数通过迭代每个链表并使用 Merge 函数将它们逐一合并成一个链表。
func MergeKLists(lists []*Node) *Node {
	// prev 用于跟踪上一次合并操作的结果。
	var prev *Node
	// 遍历链表切片。
	for i := 0; i < len(lists); i++ {
		// 如果当前链表为空，则跳过，继续处理下一个链表。
		if lists[i] == nil {
			continue
		}
		// 将当前链表与 prev 合并，并更新 prev 为新的结果。
		prev = Merge(prev, lists[i])
	}
	// 返回合并后的链表头节点。
	return prev
}

// HasCycle 检查链表中是否存在环。
// 使用快慢指针策略，慢指针每次移动一步，快指针每次移动两步。
// 如果链表中存在环，快慢指针最终会相遇；如果链表无环，快指针会先到达链表尾部。
// 参数:
//
//	head (*Node): 链表的头节点。
//
// 返回值:
//
//	bool: 如果链表中存在环，则返回true；否则返回false。
func HasCycle(head *Node) bool {
	// 检查链表是否为空或只有一个节点，这样的链表不可能有环。
	if head == nil || head.Next == nil {
		return false
	}

	// 初始化快慢指针，慢指针从头节点开始，快指针从第二个节点开始。
	slow, fast := head, head

	// 遍历链表，直到快指针或其下一个节点为空。
	for fast != nil && fast.Next != nil {
		// 慢指针向前移动一步。
		slow = slow.Next
		// 快指针向前移动两步。
		fast = fast.Next.Next
		// 如果快慢指针相遇，说明链表中存在环。
		if slow == fast {
			return true
		}

	}
	// 如果快指针到达链表尾部，说明链表无环。
	return false
}

// EntryNodeOfLoop 寻找链表中环的入口节点。
//
// 为了找出链表中环的入口节点，可以采用快慢指针的方法，具体步骤如下：
// 首先判断链表是否有环：设置两个指针，一个快指针 fast 每次移动两步，一个慢指针 slow 每次移动一步。
// 若链表有环，这两个指针必然会在环内相遇。
// 若链表有环，计算环的入口节点：当快慢指针相遇后，将其中一个指针（如 slow）重新指向链表头节点，
// 另一个指针（如 fast）保持在相遇点，然后让两个指针都以每次一步的速度移动，它们再次相遇的节点就是环的入口节点。
//
// 如果链表中有环，返回环的入口节点；如果没有环，返回nil。
// 参数 pHead 是链表的头节点。
func EntryNodeOfLoop(pHead *Node) *Node {
	// 检查链表是否为空或只有一个节点，如果是，直接返回nil，因为这些情况下不可能有环。
	if pHead == nil || pHead.Next == nil {
		return nil
	}

	// 初始化两个指针，slow和fast都指向链表的头节点。
	// slow每次移动一步，fast每次移动两步。
	slow, fast := pHead, pHead

	// 遍历链表，直到fast指针到达链表尾部。
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// 如果slow和fast相遇，说明链表中有环。
		if slow == fast {
			// 将fast指针重新指向链表头部，slow保持在相遇点。
			fast = pHead

			// 两个指针再次以相同的速度移动，直到它们再次相遇。
			// 再次相遇的点就是环的入口节点。
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}

			// 返回环的入口节点。
			return slow
		}
	}

	// 如果fast指针到达链表尾部，说明链表中没有环，返回nil。
	return nil
}

// FindKthToTail 寻找链表中倒数第k个节点。
// pHead 是链表的头节点，k 是待查找的倒数位置。
// 如果链表为空或k小于等于0，则返回nil。
// 如果链表的长度小于k，也会返回nil。
func FindKthToTail(pHead *Node, k int) *Node {
	// 检查输入参数的有效性
	if pHead == nil || k <= 0 {
		return nil
	}

	// 初始化快指针
	fast := pHead

	// 快指针先前进k步
	for i := 0; i < k; i++ {
		// 如果k大于链表长度，则返回nil
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	// 初始化慢指针
	slow := pHead

	// 快慢指针同时前进，当快指针到达链表尾部时，慢指针指向倒数第k个节点
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 返回倒数第k个节点
	return slow
}

// FindFirstCommonNode 寻找两个链表的第一个公共节点。
// 该函数的主要思路是：使用两个指针p1和p2分别从两个链表的头节点开始遍历，如果p1遍历完链表 1 后，将其指向链表 2 的头节点继续遍历；
// 同理，如果p2遍历完链表 2 后，将其指向链表 1 的头节点继续遍历。
// 当p1和p2相等时，这个节点就是第一个公共节点；如果没有公共节点，最终p1和p2都会指向nil。
//
// 这个函数通过双指针技术实现，避免了额外的空间开销，且时间复杂度为O(n+m)，n和m分别是两个链表的长度。
// 参数:
//
//	pHead1 *Node: 第一个链表的头节点。
//	pHead2 *Node: 第二个链表的头节点。
//
// 返回值:
//
//	*Node: 返回第一个公共节点，如果没有公共节点，则返回nil。
func FindFirstCommonNode(pHead1 *Node, pHead2 *Node) *Node {
	// 检查输入的链表头节点是否为空，如果任意一个链表为空，则直接返回nil。
	if pHead1 == nil || pHead2 == nil {
		return nil
	}

	// 初始化两个指针p1和p2，分别指向两个链表的头节点。
	p1, p2 := pHead1, pHead2

	// 循环直到p1和p2指向同一个节点，该节点即为第一个公共节点。
	for p1 != p2 {
		// 如果p1到达链表末尾，则重置为第二个链表的头节点；否则，p1移动到下一个节点。
		if p1 == nil {
			p1 = pHead2
		} else {
			p1 = p1.Next
		}
		// 如果p2到达链表末尾，则重置为第一个链表的头节点；否则，p2移动到下一个节点。
		if p2 == nil {
			p2 = pHead1
		} else {
			p2 = p2.Next
		}
	}

	// 返回p1（或p2，因为此时p1==p2），如果两个链表没有公共节点，则p1和p2最终都会变成nil。
	return p1
}

// AddInList 将两个逆序存储的链表代表的数字相加，返回一个新的逆序存储的链表。
// 首先反转两个输入链表，以便从低位开始相加。
// 最后，反转结果链表并返回。
// head1 和 head2 分别是两个链表的头节点。
func AddInList(head1 *Node, head2 *Node) *Node {
	// 首先反转两个输入链表，以便从低位开始相加。
	head1 = ReverseList(head1)
	head2 = ReverseList(head2)

	// 创建一个虚拟头节点，简化链表操作。
	dummy := &Node{}
	var prev *Node
	var carry int
	// 遍历两个链表，直到两者都为空。
	// 使用 carry 来处理进位。
	for head1 != nil || head2 != nil {
		sum := carry
		// 如果当前节点不为空，则加上其值，并移动到下一个节点。
		if head1 != nil {
			sum = sum + head1.Val
			head1 = head1.Next
		}
		if head2 != nil {
			sum = sum + head2.Val
			head2 = head2.Next
		}
		// 计算当前位的值和进位。
		val := sum % 10
		carry = sum / 10
		// 如果 prev 为空，说明这是第一个节点，需要初始化它。
		if prev == nil {
			prev = &Node{Val: val}
			dummy.Next = prev
		} else {
			// 否则，将新的节点添加到链表末尾，并更新 prev。
			prev.Next = &Node{Val: val}
			prev = prev.Next
		}
	}
	// 最后，反转结果链表并返回，以满足逆序存储的要求。
	return ReverseList(dummy.Next)
}

// SortInList 使用归并排序对链表进行排序。
// 它首先将链表分成两半，然后分别对每半进行排序，最后将它们合并成一个有序链表。
// 参数 head 是待排序链表的头节点。
// 返回值是排序后的链表头节点。
func SortInList(head *Node) *Node {
	// 如果链表为空或只有一个节点，则无需排序，直接返回头节点。
	if head == nil || head.Next == nil {
		return head
	}

	// 使用快慢指针找到链表的中点。
	// slow 指针最终会指向链表的中点（如果链表长度为奇数）或中间两个节点的第一个（如果链表长度为偶数）。
	// fast 指针用于辅助找到中点，它每次移动两步。
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// mid 是链表的第二部分的起始节点。
	// 将链表从中间断开，形成两个独立的链表。
	mid := slow.Next
	slow.Next = nil

	// 递归地对左半部分和右半部分进行排序。
	left := SortInList(head)
	right := SortInList(mid)

	// 合并两个有序链表并返回合并后的头节点。
	return Merge(left, right)
}

// IsPail 检查链表是否为回文链表。
// 使用快慢指针找到链表的中间点，然后反转中间点之后的链表部分。
// 最后，比较前半部分和反转后的后半部分链表是否相等。
// 参数:
//
//	head *Node: 链表的头节点。
//
// 返回值:
//
//	bool: 如果链表是回文链表，则返回true；否则返回false。
func IsPail(head *Node) bool {
	// 如果链表为空或只有一个节点，则视为回文链表。
	if head == nil || head.Next == nil {
		return true
	}

	// 初始化快慢指针，slow每次移动一步，fast每次移动两步。
	slow, fast := head, head.Next
	// 快指针和快指针的下一节点不为空时，继续移动指针。
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// slow指针现在位于链表的中间点，mid为中间点后的第一个节点。
	mid := slow.Next
	// 将链表从中点切断，得到两个独立的链表部分。
	slow.Next = nil

	// 反转中间点后的链表部分。
	mid = ReverseList(mid)

	// 比较前半部分和反转后的后半部分链表是否相等。
	for head != nil && mid != nil {
		if head.Val != mid.Val {
			return false
		}
		head = head.Next
		mid = mid.Next
	}

	// 如果所有对应节点相等，则链表是回文链表。
	return true
}

// OddEvenList 将一个单链表重新排序，使得其节点按照奇偶序号交替排列。
// 该函数接受一个单链表的头节点指针作为参数，并返回重新排序后的链表头节点指针。
// 参数:
//
//	head - 指向单链表头节点的指针。
//
// 返回值:
//
//	指向重新排序后链表头节点的指针。
func OddEvenList(head *Node) *Node {
	// 如果链表为空或只有一个节点，则无需重新排序，直接返回头节点。
	if head == nil || head.Next == nil {
		return head
	}

	// 初始化奇数和偶数节点序列的头节点。
	odd, even := head, head.Next
	oddHead, evenHead := odd, even

	// 遍历链表，将节点按奇偶序号分离到不同的链表中。
	for even != nil && even.Next != nil {
		// 连接奇数序号的节点。
		odd.Next = even.Next
		odd = odd.Next
		// 连接偶数序号的节点。
		even.Next = odd.Next
		even = even.Next
	}

	// 将奇数序号链表的末尾连接到偶数序号链表的头部，完成重新排序。
	odd.Next = evenHead
	return oddHead
}

// DeleteDuplicates 删除排序链表中的所有重复元素，使每个元素只出现一次。
// 参数 head 是链表的头节点。
// 返回值是删除重复元素后的链表的头节点。
func DeleteDuplicates(head *Node) *Node {
	// cur 用于遍历链表，初始指向头节点。
	cur := head
	// 遍历链表直到 cur 为 nil，即链表末尾。
	for cur != nil {
		// next 指向当前节点 cur 的下一个节点。
		next := cur.Next
		// 如果 next 不为 nil 且其值等于当前节点 cur 的值，说明有重复。
		for next != nil && next.Val == cur.Val {
			// 继续遍历直到找到第一个不重复的节点。
			next = next.Next
		}
		// 将当前节点的 next 指针指向第一个不重复的节点，删除中间的重复节点。
		cur.Next = next
		// 将 cur 移动到下一个节点，继续遍历。
		cur = next
	}
	// 返回处理后的链表头节点。
	return head
}

// DeleteAllDuplicates 删除链表中所有重复的节点，只保留唯一的节点。
// 参数:
//
//	head *Node: 链表的头节点。
//
// 返回值:
//
//	*Node: 删除重复节点后的链表的头节点。
func DeleteAllDuplicates(head *Node) *Node {
	// 创建一个虚拟节点，方便处理头节点可能是重复节点的情况。
	dummy := &Node{Next: head}
	// prev用于记录最后一个不重复节点的位置。
	prev := dummy
	// cur用于遍历链表。
	cur := head

	// 遍历链表，直到cur为nil，表示链表遍历完毕。
	for cur != nil {
		// 如果cur的下一个节点不为空且值与cur相同，继续向后遍历直到找到第一个不重复的节点。
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next
		}
		// 如果prev的下一个节点是cur，说明cur没有重复，将prev移动到cur位置。
		if prev.Next == cur {
			prev = cur
		} else {
			// 否则，说明cur有重复，将prev的下一个节点指向cur的下一个节点，跳过所有重复的节点。
			prev.Next = cur.Next
		}
		// 将cur移动到下一个节点，继续遍历。
		cur = cur.Next
	}
	// 返回处理后的链表的头节点。
	return dummy.Next
}
