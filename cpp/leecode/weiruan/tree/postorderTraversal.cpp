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


// 递归解决
/*
class Solution {
public:
    void visit(vector<int>& result, TreeNode* node){
        if ( node == NULL ){
            return;
        }
        visit(result,node->left);
        visit(result,node->right);
        result.push_back(node->val);

    }
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int>result;
        visit(result,root);
        return result;
    }
};
*/

// 非递归实现，有点难，没有实现。
class Solution {
public:
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int>result;
        TreeNode* cur =root;
        stack<TreeNode*>tmp;
        while( cur != NULL ){
            tmp.push(cur);
            if (cur->right != NULL){
                tmp.push(cur->right);
            }
            if (cur->left != NULL){
                tmp.push(cur->left);
                cur = cur->left;
            }else if (cur->right != NULL){
                cur = cur->right;
            }else{
                cur = NULL;
            }
        }
        return result;
    }
};