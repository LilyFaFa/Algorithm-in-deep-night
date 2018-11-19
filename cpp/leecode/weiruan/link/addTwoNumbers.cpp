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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        stack<int>s1;
        stack<int>s2;
        ListNode* head1=l1;
        ListNode* head2=l2;
        while(head1!=NULL){
            s1.push(head1->val);
            head1 = head1->next;
        }
        while(head2!=NULL){
            s2.push(head2->val);
            head2 = head2->next;
        }
        ListNode* result= new ListNode(-1);
        int n = 0;
        while(!s1.empty() && !s2.empty()){
            int i1 = s1.top();
            int i2 = s2.top();
            s1.pop();
            s2.pop();
            ListNode* cur;
            cur->val = (i1+i2+n)%10;
            n=(i1+i2+n)/10;
            cur->next=result->next;
            result->next=cur;
        }
        while (!s2.empty()){
            int i2 = s2.top();
            ListNode* cur;
            cur->val = (i2+n)%10;
            n=(i2+n)/10;
            cur->next=result->next;
            result->next=cur;
        }
        while (!s1.empty()){
            int i1 = s1.top();
            ListNode* cur;
            cur->val = (i1+n)%10;
            n=(i1+n)/10;
            cur->next=result->next;
            result->next=cur;
        }
        if (n==1){
            ListNode* cur;
            cur->val = 1;
            cur->next=result->next;
            result->next=cur;
        }
        return result->next;
    }
};