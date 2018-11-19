#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

struct TreeNode{
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x):val(x),left(NULL),right(NULL){};
};

class Solution {
public:
    int depth(TreeNode* t){
        if (t == NULL){
            return 0;
        }
        int l = depth(t->left);
        int r = depth(t->right);
        return max(l,r)+1;
    }
    
    TreeNode* subtreeWithAllDeepest(TreeNode* root) {
        if (root == 0){
            return NULL;
        }
        TreeNode* result = root;
        int leftDepth = depth(result->left);
        int rightDepth = depth(result->right);
        while(leftDepth != rightDepth){
            if (leftDepth > rightDepth){
                result = result->left;
            }else{
                result = result->right;
            }
            leftDepth = depth(result->left);
            rightDepth =depth(result->right);
        }
        return result;
    }
};
