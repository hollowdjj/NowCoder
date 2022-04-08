package list

import "nowcoder/utility"

/*
合并 k 个升序的链表并将结果作为一个升序的链表返回其头节点。
数据范围：节点总数0≤n≤5000，每个节点的val满足 |val| <= 1000
要求：时间复杂度 O(nlogn)
*/

//MergeKLists 合并k个升序链表
func MergeKLists(lists []*utility.ListNode) *utility.ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	} else if n == 1 {
		return lists[0]
	}

	return mergeTwoLists(MergeKLists(lists[:n/2]), MergeKLists(lists[n/2:]))
}

//mergeTwoLists 归并两升序链表
func mergeTwoLists(head1, head2 *utility.ListNode) *utility.ListNode {
	dummy := &utility.ListNode{Val: -1}
	temp := dummy
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			temp.Next = head1
			head1 = head1.Next
		} else {
			temp.Next = head2
			head2 = head2.Next
		}
		temp = temp.Next
	}
	if head1 != nil {
		temp.Next = head1
	} else {
		temp.Next = head2
	}

	return dummy.Next
}

/*
另一个思路是使用优先队列。每次取出队头元素，加入到新链表尾部，然后将它的下一个元素
放入有限队列中。

ListNode* mergeKLists(vector<ListNode*>& lists) {
        auto cmp = [](ListNode* lhs,ListNode* rhs) {
            return lhs->val > rhs->val;
        };
        priority_queue<ListNode*,vector<ListNode*>,decltype(cmp)> q(cmp);
        auto dummy = new ListNode(-1);
        auto head = dummy;
        for(int i = 0;i<lists.size();++i) {
            if(lists[i]) {
                q.push(lists[i]);
            }

        }
        while(!q.empty()) {
            auto top = q.top();
            q.pop();
            head->next = top;
            if(top->next) {
                q.push(top->next);
            }
            head = head->next;
        }
        return dummy->next;
    }
*/
