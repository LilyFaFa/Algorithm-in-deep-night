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

// 转换为图，然后广度优先搜索
// 深度计算，
class Solution {
public:
    void collect(vector<int>& result, TreeNode* root, int dis){
        if (dis < 0 || root == NULL){
            return;
        }
        queue<TreeNode*>tmp;
        tmp.push(root);
        int dep =-1;
        while( !tmp.empty() ){
            dep++;
            int s =tmp.size();
            if (dep == dis){
                for (int i=0; i<s; i++){
                    result.push_back(tmp.front()->val);
                    tmp.pop();
                }
                return;
            }
            for (int i=0; i<s; i++){
                if(tmp.front()->left !=NULL){
                    tmp.push(tmp.front()->left);
                }
                if (tmp.front()->right != NULL){
                    tmp.push(tmp.front()->right);
                }
                tmp.pop();
            }
        }

    }
    int dis(TreeNode* root, TreeNode* target,int distance,vector<int>& result){
        if (root == target){
            collect(result,target,distance);
            return 0;
        }
        if (root == NULL){
            return -1;
        }
        int l = dis(root->left,target,distance,result);
        int r = dis(root->right,target,distance,result);
        if (l >=0){
            if (l ==distance-1){
                result.push_back(root->val);
            }else{
                collect(result,root->right,distance-l-2);
            }
            return l+1;
        }

        if (r >=0){
            if (r ==distance-1){
                result.push_back(root->val);
            }else{
                collect(result,root->left,distance-r-2);
            }
            return r+1;
        }

        return -1;
    }
    vector<int> distanceK(TreeNode* root, TreeNode* target, int K) {
        map<TreeNode*,vector<TreeNode*> >graph;
        vector<int>result;
        dis(root,target,K,result);
        return result;
    }
};