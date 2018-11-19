#include<iostream>
#include<vector>
using namespace std;

struct node{
    int val;
    node* next;
};

void printK(node* head, int k){
    node* first=head;
    node* second=head;
    while(k--){
        if (first == NULL){
            return;
        }
        first = first->next;
    }
    while(first != NULL){
        first=first->next;
        second=second->next;
    }
    cout<<second->val;
    return;
}
int main(){
    node n1,n2,n3,n4,n5,n6,n7;
    n1.val=1;
    n2.val=2;
    n3.val=3;
    n4.val=4;
    n5.val=5;
    n6.val=6;
    n7.val=7;
    n1.next = &n2;
    n2.next = &n3;
    n3.next = &n4;
    n4.next = &n5;
    n5.next = &n6;
    n6.next = &n7;
    n7.next = NULL;
    printK(&n1,3);
}