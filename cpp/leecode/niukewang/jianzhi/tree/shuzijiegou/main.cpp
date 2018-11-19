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
    bool HasSubtree(TreeNode* pRoot1, TreeNode* pRoot2)
    {
        if (pRoot1 == NULL || pRoot2 == NULL){
            return false;
        }
        return checkTree(pRoot1,pRoot2,pRoot2);
    }
private:
    bool checkTree(TreeNode* pRoot1, TreeNode* pRoot2 , TreeNode* root2){
        if ( pRoot2 == NULL){
            return true;
        }
        if ( pRoot1 == NULL ){
            return false;
        }

        if ( pRoot1->val == pRoot2->val && pRoot1->val == root2->val){
            return checkTree(pRoot1->left,pRoot2->left,root2) && checkTree(pRoot1->right,pRoot2->right,root2) || checkTree(pRoot1->right,root2->right,root2) && checkTree(pRoot1->left,root2->left,root2) ||checkTree(pRoot1->left,root2,root2) || checkTree(pRoot1->right,root2,root2);
        }else if ( pRoot1->val == pRoot2->val){
            return checkTree(pRoot1->left,pRoot2->left,root2) && checkTree(pRoot1->right,pRoot2->right,root2)||checkTree(pRoot1->left,root2,root2) || checkTree(pRoot1->right,root2,root2);
        }else if (pRoot1->val == root2->val){
            return checkTree(pRoot1->left,root2->left,root2) && checkTree(pRoot1->right,root2->right,root2)||checkTree(pRoot1->left,root2,root2) || checkTree(pRoot1->right,root2,root2);
        }
        return checkTree(pRoot1->left,root2,root2) || checkTree(pRoot1->right,root2,root2);
    
    }
};