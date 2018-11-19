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
    int TreeDepth(TreeNode* pRoot){
        return checkDepth(pRoot);
    }
private:
    int checkDepth(TreeNode* node){
        if (node == NULL){
            return 0;
        }
        if (checkDepth(node->right) > checkDepth(node->left)){
            return 1+checkDepth(node->right) ;
        }
        return 1+checkDepth(node->left);
    }
};