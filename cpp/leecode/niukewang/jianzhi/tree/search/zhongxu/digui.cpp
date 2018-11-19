#include<vector>
#include<iostream>

using namespace std;
struct TreeNode {
	int val;
	struct TreeNode *left;
	struct TreeNode *right;
	TreeNode(int x) :
			val(x), left(NULL), right(NULL) {
	}
};

void search(TreeNode* node){
    if (node == NULL){
        return;
    }
    cout<< node->val <<"\n";
    search(node->left);
    search(node->right);
}

int main(){
    TreeNode root = TreeNode(3);
    TreeNode left = TreeNode(2);
    TreeNode right = TreeNode(1); 
    root.left=&left;
    root.right=&right;
    search(&root);
    return 0;
}