#include<vector>
#include<iostream>
using namespace std;
// 面试题39:二叉树的深度
struct Node{
    int val;
    Node* left;
    Node* right;
};

int getDeepOfTree(Node* root){
    if (root == NULL){
        return 0;
    }
    int rightDeep=getDeepOfTree(root->left);
    int leftDeep=getDeepOfTree(root->right);
    return max(rightDeep,leftDeep)+1;
}

// 是不是平衡二叉树
bool IsBalancer(Node* root){
    if(root == NULL){
        return true;
    }
    int leftDeep=getDeepOfTree(root->left);
    int rightDeep=getDeepOfTree(root->right);
    if(-1<=(leftDeep-rightDeep) && -1>=(leftDeep-rightDeep) ){
        return false;
    }
    return (IsBalancer(root->left) && IsBalancer(root->right));
}