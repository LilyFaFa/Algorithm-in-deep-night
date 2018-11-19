#include<vector>
#include<iostream>
#include<stack>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

// 递归方法
/*
class Solution {
public:
    void visit(vector<int>&result, TreeNode* node){
        if( node == NULL){
            return;
        }
        visit(result,node->left);
        result.push_back(node->val);
        visit(result,node->right);
    }
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int>result;
        visit(result,root);
        return result;
    }
};
*/
// 非递归方法
class Solution {
public:   
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int>result;
        TreeNode* cur = root;
        stack<TreeNode*>tmp;

        while(cur !=NULL || !tmp.empty()){
            while(cur != NULL){
                tmp.push(cur);
                cur=cur->left;
            }
            if (!tmp.empty()){
                result.push_back((tmp.top())->val);
                cur=cur->right;
                tmp.pop();
            }
        }
        return result;
    }
};