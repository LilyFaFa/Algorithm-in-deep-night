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
    void dfs(TreeNode* cur, vector< pair<int,int> >& tmp, int level, int order){
        if (cur == NULL){
            return;
        }
        if (tmp.size() < level){
            pair<int, double> p(order,order);
            tmp.push_back(p);
        }else{
            tmp[level-1].second=order;
        }
        dfs(cur->left,tmp,level+1,order*2);
        dfs(cur->right,tmp,level+1, order*2+1);
        return;
    }
    int widthOfBinaryTree(TreeNode* root) {
        if ( root == NULL ){
            return 0;
        }
        vector< pair<int,int> >tmp;
        dfs(root, tmp, 1,1);
        int result =0;
        for (int i=0; i<tmp.size(); i++){
            if (tmp[i].second - tmp[i].first >= result){
                result= tmp[i].second - tmp[i].first +1;
            }
        }
        return result;
    }
};