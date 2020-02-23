package solution 

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    	
	//先定义一个新的节点，初始化链表
		temp := &ListNode{0,nil}
		dump := temp
	
		//
	for l1 !=nil && l2 !=nil {
		if l1.Val > l2.Val{
			dump.Next= l2
			l2 = l2.Next
		} else {
			dump.Next = l1
			l1 = l1.Next
		}
		//将遍历的指针向后移动一位
		  dump = dump.Next
	}
	if l1 == nil {
		dump.Next = l2
	}
	if l2 == nil {
		dump.Next = l1
	}
	
	 return temp.Next
		
	
	}