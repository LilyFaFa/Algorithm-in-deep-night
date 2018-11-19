#include<iostream>
#include<vector>
using namespace std;

struct node{
    int val;
    node* next;
};

vector<int> GetMerge(node* node1 ,node* node2){
    vector<int>result;
    if (node1 == NULL || node2 == NULL){
        return result;
    }
    
    node* first = node1;
    node* second = node2;

    while(first != NULL &&  second != NULL){
        if (first->val == second->val){
            result.push_back(first->val);
            first=first->next;
            second=second->next;
        }else if (first->val > second->val){
            second=second->next;
        }else{
            first=first->next;
        }
    }
    
    return result;
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
    
    vector<int> result = GetMerge(&n1, &n3);
    for (int i =0; i<result.size(); i++){
        cout<<result[i]<<"\n";
    }
    
    return 0;
}