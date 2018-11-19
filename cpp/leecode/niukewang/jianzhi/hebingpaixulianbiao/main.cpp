#include<iostream>
using namespace std;

struct ListNode {
	int val;
	struct ListNode *next;
	ListNode(int x) :
			val(x), next(NULL) {
	}
};

class Solution {
public:
    ListNode* Merge(ListNode* pHead1, ListNode* pHead2)
    {
        ListNode* result = new ListNode(0);
        ListNode* r = result;

        while ( pHead1 != NULL && pHead2 != NULL ){
            if ( pHead1->val < pHead2->val ) {
                result->next = pHead1;
                pHead1 = pHead1->next;
            }else{
                result->next = pHead2;
                pHead2 = pHead2->next;
            }
            result = result->next;
        }

        if ( pHead1 != NULL ){
            result->next = pHead1;
        }
        if ( pHead2 != NULL ){
            result->next = pHead2;
        }
        return r->next;
    }
    
};
