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
/*
class Solution {
public:
    bool IsBalanced_Solution(TreeNode* pRoot) {
        return checkBalance(pRoot);
    }
private:
    int depth(TreeNode* node){
        if (node == NULL){
            return 0;
        }
        if ( depth(node->left) > depth(node->right)){
            return 1+depth(node->left);
        }else{
            return 1+depth(node->right);
        }
    }

    bool checkBalance(TreeNode* node){
        if (node == NULL){
            return true;
        }
        if ( depth(node->left)>depth(node->right)+1 || depth(node->left)<depth(node->right)-1){
            return false;
        }
        if ( !checkBalance(node->left)){
            return false;
        }
        if ( !checkBalance(node->right)){
            return false;
        }
        return true;
    }
};
*/

class Solution {
public:
    bool IsBalanced_Solution(TreeNode* pRoot) {
        bool balancer=true;
        int i= depth(pRoot,&balancer);
        return balancer;

    }
private:
    int depth(TreeNode* node ,bool* balancer){
        if (node == NULL){
            return 0;
        }
        if (*balancer == false){
            return 0;
        }

        int leftHight = depth(node->left, balancer);
        int rightHight = depth(node->right, balancer);
        
        if (leftHight > rightHight){
            if ( leftHight - rightHight >1){
                *balancer =false;
                return 0;
            }else{
                return 1 + depth(node->left,balancer);
            }
        } else{
           if ( rightHight - leftHight  >1){
                *balancer =false;
                return 0;
            }else{
                return 1+depth(node->right,balancer);
            } 
        }
    }
};