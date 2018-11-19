#include<iostream>

struct ListNode {
	int val;
	struct ListNode *next;
	ListNode(int x) :
			val(x), next(NULL) {
	}
};

class Solution {
public:
    ListNode* ReverseList(ListNode* pHead) {
        if (pHead == NULL || pHead->next == NULL){
            return pHead;
        }
        ListNode* first = pHead;
        ListNode* second = pHead->next;
        ListNode* third ;
         first->next=NULL;

        while ( second->next != NULL ){
            third = second->next;
            second->next = first;
            first = second;
            second = third;
        }
        second->next = first;
        return second;
    }
    
};