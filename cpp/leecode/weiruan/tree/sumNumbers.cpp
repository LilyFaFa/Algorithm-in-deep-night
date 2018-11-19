#include<vector>
#include<iostream>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    void search(TreeNode* t, vector<int>&result, int cur){
        if (t->left == NULL && t->right == NULL){
            result.push_back( cur*10 + t->val );
        }
        if (t->left != NULL){
            search(t->left,result,cur*10+t->val);
        }
        if (t->right != NULL){
            search(t->right,result,cur*10+t->val);
        }

    }
    int sumNumbers(TreeNode* root) {
        int result=0;
        vector<int>tmp;
        search(root,tmp,0);
        for(int i=0; i<tmp.size(); i++){
            result+=tmp[i];
        }
        return result;
    }
};