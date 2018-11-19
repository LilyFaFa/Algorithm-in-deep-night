#include<iostream>
#include<vector>
#include<stack>
using namespace std;
struct ListNode{
    int val;
    ListNode* next;
    ListNode(int x):val(x),next(NULL){};
};

class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        if (head == NULL || head->next ==NULL){
            return head;
        }

        ListNode* p1 = new ListNode(-1);
        ListNode* result=p1;
        ListNode* p2 = head;
        ListNode* p3 = head->next;
        while(p3 != NULL){
            while(p3 != NULL && p2->val == p3->val){
                p3 = p3->next;
            }
            if(p2->next == p3){
                p1->next =p2;
                p1=p1->next;
                p2=p2->next;
                p3=p3->next;
            }else{
                p2=p3;
                p3=p3->next;
            }
        }
        p1->next = NULL;
        return result->next;
    }
};