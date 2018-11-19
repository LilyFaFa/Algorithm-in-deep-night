#include<vector>
#include<iostream>
#include<stack>
using namespace std;

struct Node{
    int val;
    Node* left;
    Node* right;
};

// 1. 前序遍历
// 非递归，借助栈
vector<int> getPreSearch1(Node* node){
    vector<int>result;
    stack<Node*>tmp;
    Node*cur =node;
    while(cur !=NULL){
        result.push_back(cur->val);
        if ( cur->left != NULL){
            if (cur->right !=NULL){
                tmp.push(cur->right);
            }
            cur=cur->left;
        }else if(cur->right !=NULL){
            cur=cur->right;
        }else{
            cur=tmp.top();
            tmp.pop();
        }
    }
    return result;
}
// 递归
void getPreSearch2(vector<int>& result, Node* node){
    if (node == NULL){
        return;
    }
    result.push_back(node->val);
    getPreSearch2(result,node->left);
    getPreSearch2(result,node->right);
    
}
// 2. 中序遍历
// 
// 3. 后序遍历