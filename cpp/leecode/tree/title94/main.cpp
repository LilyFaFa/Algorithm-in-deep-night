#include<iostream>
#include<vector>

using namespace std;


struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
 
void visit(TreeNode* root,vector<int>&result){
    if (root == NULL){
        return;
    }
    if (root->left != NULL){
        visit(root->left,result);
    }
    result.push_back(root->val);
    if (root->right != NULL){
        visit(root->right,result);
    }
    return;
}
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> result;
        visit(root,result);
        return result;
    }
};
