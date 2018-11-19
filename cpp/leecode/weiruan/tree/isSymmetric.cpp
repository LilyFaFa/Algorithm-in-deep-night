#include<vector>
#include<iostream>
#include<stack>
#include<map>
#include<queue>
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
    bool checkNode(TreeNode* l , TreeNode* r){
        if ( (l == NULL) != (r == NULL) ){
            return false;
        }else if (l == NULL && r== NULL){
            return true;
        }
        return  l->val == r->val && checkNode(l->left,r->right) && checkNode(l->right,r->left);
    }
    bool isSymmetric(TreeNode* root) {
        if (root == NULL){
            return true;
        }
        return checkNode(root->left,root->right);
    }
};
*/
class Solution {
public:
    bool checkNode(TreeNode* n1, TreeNode* n2){
        if (n1 == NULL && n2 ==NULL){
            return true;
        }else if ((n1 == NULL) != (n2 == NULL)){
            return false;
        }else if(n1->val != n2->val){
            return false;
        }
        return checkNode(n1->left,n2->right) && checkNode(n1->right,n2->left);

    }
    bool isSymmetric(TreeNode* root) {
        if (root == NULL){
            return true;
        }
        return checkNode(root->left,root->right);
    }
};