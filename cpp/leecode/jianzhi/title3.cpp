// 题目2 链表题目
#include<vector>
#include<iostream>
#include<stack>
using namespace std;

struct Node{
    int val;
    Node* next;
};

// 向链表尾部添加节点
void AddToTail(Node** head, int num){
    Node cur;
    cur.val=num;
    if (head == NULL || *head ==NULL){
        *head = &cur;
        return;
    }
    Node* pNode = *head;
    while(pNode->next != NULL){
        pNode=pNode->next;
    }
    pNode->next=&cur;
    return;
}

// 移除节点
void RemoveNode(Node** node,int num){
    if (node == NULL || *node ==NULL){
        return;
    }
    Node* pNode=*node;
    Node* pPre;
   
    if ( (*node)->val == num){
        *node = (*node)->next;
        return;
    }else{
        pNode=(*node)->next;
        pPre=*node;
    }
    while(pNode!=NULL){
        if (pNode->val == num){
            pPre->next=pNode->next;
            return;
        }
        pNode = pNode->next;
        pPre = pPre->next;
    }
    return;
}

// 题目五、从尾到头打印链表
void PrintList(Node* head){
    if(head == NULL){
        return;
    }

    stack<int>tmp;
    Node*cur = head;
    while( cur->next !=NULL ){
        tmp.push(cur->val);
        cur=cur->next;
    }
    
    while(!tmp.empty()){
        cout<<tmp.top()<<endl;
        tmp.pop();
    }
}