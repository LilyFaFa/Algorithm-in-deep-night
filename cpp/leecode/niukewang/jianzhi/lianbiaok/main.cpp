#include<iostream>
#include<vector>


struct ListNode {
	int val;
	struct ListNode *next;
	ListNode(int x) :
			val(x), next(NULL) {
	}
};

class Solution {
public:
    ListNode* FindKthToTail(ListNode* pListHead, unsigned int k) {

        if ( pListHead == NULL ){
            return NULL;
        }
        ListNode* first = pListHead;
        ListNode* second = pListHead;
        for (int i=0; i<k-1; i++){
            if ( first->next == NULL){
                return NULL;
            }
            first = first->next;
        }

        while ( first->next != NULL){
            first = first->next;
            second = second->next;
        }
        return second;
    }
};

int main(){
    return 0;
}