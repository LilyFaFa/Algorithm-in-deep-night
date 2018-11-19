#include<vector>
#include<iostream>
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
    TreeNode* find(TreeNode* n, TreeNode* p, TreeNode* q){
        // 边界判断
        if (n == NULL){
            return NULL;
        }
        if (n == p || n == q){
            return n;
        }
        // divide
        TreeNode* left = find(n->left, p, q);
        TreeNode* right = find(n->right, p, q);

        // conquer
        if (left!=NULL && right!=NULL){
            return n;
        }else if (left !=NULL){
            return left;
        }else if (right != NULL){
            return right;
        }

    }
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        TreeNode* result = find(root,p, q);
        return result;

    }
};