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
    vector<vector<int> > FindPath(TreeNode* root,int expectNumber) {
        vector<vector<int> > r;
        vector<vector<int> >& results = r;
        vector<vector<int> > r2;
        vector<int> cur;
        find(root,cur,expectNumber,0,results);
        for (int i=results.size()-1; i>=0; i--){
            r2.push_back(results[i]);
        }
        return r2;
    }
private:
    void find(TreeNode* root, vector<int> cur, int expectNumber, int curSum, vector<vector<int> >& results){
        if (root == NULL ){
            return;
        }
        cur.push_back(root->val);
        curSum+=root->val;
        if ( curSum == expectNumber ){
            if (root->right == NULL && root->left == NULL){
                results.push_back(cur);
            }        
        }else if (curSum < expectNumber){
            find(root->right,cur,expectNumber,curSum,results);
            find(root->left,cur,expectNumber,curSum,results);
        }
        return;
    }
};