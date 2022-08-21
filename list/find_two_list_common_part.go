package list

// FindTwoListCommonPart 给定两个有序链表的头指针head1和head2，打印两个链表的公共部分
func FindTwoListCommonPart(head1 *SingleLinkedNode, head2 *SingleLinkedNode) []int {
	var res []int
	for head1 != nil && head2 != nil {
		if head1.Value == head2.Value {
			res = append(res, head1.Value)
			head1 = head1.Next
			head2 = head2.Next
		} else if head1.Value < head2.Value {
			head1 = head1.Next
		} else {
			head2 = head2.Next
		}
	}
	return res
}
