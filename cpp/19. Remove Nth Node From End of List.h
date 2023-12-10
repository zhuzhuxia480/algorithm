//
// Created by zzx on 2023/12/10.
//

#ifndef CPP_19_REMOVE_NTH_NODE_FROM_END_OF_LIST_H
#define CPP_19_REMOVE_NTH_NODE_FROM_END_OF_LIST_H

//Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;

    ListNode() : val(0), next(nullptr) {}

    ListNode(int x) : val(x), next(nullptr) {}

    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
    ListNode *removeNthFromEnd(ListNode *head, int n) {
        ListNode *p = head;
        int length = 1;
        while (p->next != nullptr) {
            p = p->next;
            length++;
        }
        int index = length - n + 1;
        ListNode pre;
        pre.next = head;
        ListNode *p1 = &pre;
        for (int i = 0; i < index - 1; ++i) {
            p1 = p1->next;
        }
        p1->next = p1->next->next;
        return pre.next;
    }
};

#endif //CPP_19_REMOVE_NTH_NODE_FROM_END_OF_LIST_H
