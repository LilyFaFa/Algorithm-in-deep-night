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
/*
class Solution {
public:
    // 递归调用
    void visit(TreeNode* node, vector<int>& result){
        if (node ==NULL){
            return;
        }
        result.push_back(node->val);
        visit(node->left,result);
        visit(node->right,result);
    }
    
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> result;
        visit(root,result);
        return result;
    }
};
*/
/*
class Solution {
public:
    // 非递归实现
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> result;
        TreeNode* cur = root;
        stack<TreeNode*>tmp;
        while( cur != NULL ){
            // 访问当前节点
            result.push_back(cur->val);
            // 右节点压栈
            if (cur->right != NULL){
                tmp.push(cur->right);
            }
            // 访问左节点，左节点如果为空，则访问栈顶的右节点；
            if (cur->left != NULL){
                cur=cur->left;
            }else if (!tmp.empty()){
                cur=tmp.top();
                tmp.pop();
            }else{
                cur=NULL;
            }
        }
        return result;
    }
};
*/
//另一个种非递归实现
class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        stack<TreeNode*>tmp;
        vector<int>result;
        TreeNode* cur =root;

        while(cur != NULL || !tmp.empty()){
            while(cur != NULL){
                result.push_back(cur->val);
                tmp.push(cur);
                cur = cur->left;
            }
            if (!tmp.empty()){
                cur = tmp.top();
                cur = cur->right;
                tmp.pop();
            }
        }
        return result;
    }
};