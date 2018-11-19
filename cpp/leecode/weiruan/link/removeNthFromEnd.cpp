#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* fake =new ListNode(1);
        fake->next = head;
        
        ListNode* next = fake;
        ListNode* l = fake;
        while( n > 0){
            next=next->next;
            n--;
        }
        while(next->next != NULL){
            l=l->next;
            next=next->next;
        }

        l->next =l->next->next;
        ListNode* result = fake->next;
        return result;
    }
};