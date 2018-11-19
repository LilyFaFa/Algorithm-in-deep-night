#include<iostream>
#include<vector>
using namespace std;
struct TreeNode{
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x):val(x),left(NULL),right(NULL){};
};
class Solution {
public:
    void seachPath(TreeNode* t, vector<int>cur, vector<vector<int> >& results, int curSum, int sum){
        cur.push_back(t->val);
        curSum+=t->val;
        if (t->left == NULL && t->right == NULL){
            if (curSum == sum){
                results.push_back(cur);
                return;
            }
        }
        
        if (t->left != NULL){
            seachPath(t->left,cur,results,curSum,sum);
        }
        if (t->right != NULL){
            seachPath(t->right,cur,results,curSum,sum);
        }
    }
    vector<vector<int> > pathSum(TreeNode* root, int sum) {
        vector<vector<int> >result;
        if (root == NULL){
            return result;
        }
        vector<int>cur;
        int curSum=0;
        seachPath(root,cur,result,curSum,sum);
        return result;
    }
};