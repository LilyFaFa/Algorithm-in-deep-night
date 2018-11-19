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
    ListNode* rotateRight(ListNode* head, int k) {
        if (head == NULL){
            return NULL;
        }
        ListNode* pre = head;
        int length=1;
        while(pre->next != NULL){
            pre=pre->next;
            length++;
        }

        k = k%length;
        k = length-k;

        pre->next=head;
        ListNode* rotate = head;
        for (int i=0; i<k-1; i++){
            rotate=rotate->next;
        }
        ListNode* result = rotate->next;
        rotate->next =NULL;
        return result;
    }
};