#include<iostream>
#include<vector>
using namespace std;

struct TreeNode {
	int val;
	struct TreeNode *left;
	struct TreeNode *right;
	TreeNode(int x) :
			val(x), left(NULL), right(NULL) {
	}
};

class Solution {
public:
    void Mirror(TreeNode *pRoot) {
        Dfs(pRoot);
        return;
    }
private:
    void Dfs(TreeNode *pRoot) {
        if (pRoot == NULL){
            return;
        }
        TreeNode * tmp;
        tmp = pRoot->right;
        pRoot->right = pRoot->left;
        pRoot->left = tmp;
        Dfs(pRoot->right);
        Dfs(pRoot->left);
    }
};