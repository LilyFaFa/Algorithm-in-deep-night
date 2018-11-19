#include<algorithm>
using namespace std;

struct ListNode{
    int val;
    ListNode* next;
    ListNode(int x):val(x),next(NULL){};
};

class Solution {
public:
    // 链表插入排序
    ListNode* insertionSortList(ListNode* head) {
        if (head == NULL){
            return NULL;
        }

        ListNode* fake=new ListNode(-1);
        fake->next=head;
        ListNode* cur = head->next;
        ListNode* curFre = head;
        while(cur != NULL){
            ListNode* f=head;
            ListNode* l=fake;
            while(f != cur && f->val < cur->val ){
                f=f->next;
                l=l->next;
            }
            if (f == cur){
                cur = cur->next;
                curFre = curFre->next;
            }else{    
                ListNode* cur2 = cur->next;
                curFre->next = cur->next;
                l->next=cur;
                cur->next=f;
                if(l==fake){
                    head =fake->next;
                }
                cur=cur2;
            }
        }
        return head;
    }
};