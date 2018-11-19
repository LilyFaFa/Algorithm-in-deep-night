#include<iostream>
#include<vector>
#include<queue>
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
    vector<int> PrintFromTopToBottom(TreeNode* root) {
        vector<int> r;
        if (root == NULL){
            return r;
        }
        vector<int>& results=r;
        queue<TreeNode*> cur;
        cur.push(root);
        printTree(results, cur);
        return r;
    }
private:
   void printTree(vector<int>& results, queue<TreeNode*>nodes){
        int len= nodes.size();
        if (len == 0 ){
            return;
        }
        for (int i=0; i<len; i++){
            results.push_back(nodes.front()->val);
            if ( nodes.front()->left !=NULL ){
                nodes.push(nodes.front()->left);
            }
            if ( nodes.front()->right !=NULL ){
                nodes.push(nodes.front()->right);
            }
            nodes.pop();
        }
        printTree(results, nodes);
    }
};