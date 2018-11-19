#include<iostream>
#include<stack>
using namespace std;

struct TreeNode {
	int val;
	struct TreeNode *left;
	struct TreeNode *right;
	TreeNode(int x) :
			val(x), left(NULL), right(NULL) {
	}
};

void search(TreeNode* root){
    if (root == NULL){
        return;
    }
    stack<TreeNode*> s;
    s.push(root);
    TreeNode* tmp = s.top()->left;

    while (tmp !=NULL){
        if (tmp->left !=NULL){
            tmp = tmp->left;
            s.push(tmp);
        }else{
            cout<<tmp->val;
            
            if (tmp->right != NULL){
                tmp = tmp->right;
            }else{
                tmp=s.top();
                s.pop();
            }
        }
        
    }

    while(!s.empty()){
        tmp=s.top();
        cout<<tmp->val;
        if ( tmp->right != NULL){
            
        }
    }
    cout<<s.top()->val;
}
