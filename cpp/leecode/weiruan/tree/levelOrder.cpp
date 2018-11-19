#include<vector>
#include<iostream>
#include<queue>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    vector<vector<int> > levelOrder(TreeNode* root) {
        vector<vector <int> > result;
        if (root == NULL){
            return result;
        }
        queue<TreeNode*>tmp;
        tmp.push(root);
        while( !tmp.empty()){
            int num =tmp.size();
            vector<int>cur;
            for (int i=0; i<num; i++){
                cur.push_back(tmp.front()->val);
                if (tmp.front()->left != NULL){
                    tmp.push(tmp.front()->left);
                }
                if (tmp.front()->right != NULL){
                    tmp.push(tmp.front()->right);
                }
                tmp.pop();
            }
            result.push_back(cur);
        }
        return result;
    } 
};