//一个链表，奇数位升序，偶数位降序，如何整体排成升序？
#include<vector>
#include<iostream>
using namespace std;
struct Node{
    int cal;
    Node* next;
};

Node* GetMerge(Node* node){
    Node* head1 = node;
    Node* node1 = node;
    Node* head2 = node->next;
    Node* node2 = node->next;
    Node* cur = node->next->next;
    while(cur !=NULL){
        node1->next=cur;
        node1=node1->next;
        cur=cur->next;
        if (cur!=NULL){
            head2=cur;
            head2->next =node2;
            node2=head2;
        }else{
            break;
        }
        cur=cur->next;
    }

}