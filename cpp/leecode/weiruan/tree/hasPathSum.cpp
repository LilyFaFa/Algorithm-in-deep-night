#include<iostream>
#include<vector>
using namespace std;
struct TreeNode{
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x):val(x),left(NULL),right(NULL){};
};
/*
class Solution {
public:
    void search(TreeNode* t, int curSum, int sum, bool& finded){
        if (finded){
            return;
        }
        curSum +=t->val;
        if (t->left == NULL && t->right == NULL){ 
            if (curSum == sum){
                finded = true;
            }
            return ;
        }
        if (t->left != NULL){
            search(t->left, curSum, sum,finded);
        }
        if (t->right != NULL){
            search(t->right,curSum, sum,finded);
        }
    }
    bool hasPathSum(TreeNode* root, int sum) {
        if (root == NULL){
            return false;
        }
        bool result = false;
        int cur = 0;
        search(root,cur,sum,result);
        return result;
    }
};
*/

class Solution {
public:
    bool search(TreeNode* t, int curSum, int sum){
        // 1 边界
        if (t == NULL){
            return false;
        }
        if (t->left == NULL && t->right ==NULL){
            if (curSum+t->val == sum){
                return true;
            }
        }
        // 分治
        curSum+=t->val;
        return (search(t->left,curSum,sum) || search(t->right,curSum,sum));
    }
    bool hasPathSum(TreeNode* root, int sum) {
        if (root == NULL){
            return false;
        }
        return search(root,0,sum);
    }
};