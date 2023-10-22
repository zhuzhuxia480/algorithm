//
// Created by zzx on 2023/10/22.
//

#ifndef CPP_21_MERGE_TWO_SORTED_LISTS_H
#define CPP_21_MERGE_TWO_SORTED_LISTS_H
/**
 * Definition for singly-linked list.
 */

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ListNode head;
        if (list1 == nullptr) {
            return list2;
        }
        if (list2 == nullptr) {
            return list1;
        }

        ListNode *p = &head;
        while (list1 && list2) {
            if (list1->val < list2->val) {
                p->next = list1;
                list1 = list1->next;
            } else {
                p->next = list2;
                list2 = list2->next;
            }
            p = p->next;
        }
        while (list1) {
            p->next = list1;
            list1 = list1->next;
            p = p->next;
        }
        while (list2) {
            p->next = list2;
            list2 = list2->next;
            p = p->next;
        }
        return head.next;
    }
};
#endif //CPP_21_MERGE_TWO_SORTED_LISTS_H
















