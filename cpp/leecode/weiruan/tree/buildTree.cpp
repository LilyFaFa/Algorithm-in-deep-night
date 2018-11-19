/*
#include<iostream>
#include<vector>
#include<algorithm>//注意要包含该头文件
using namespace std;
int main(){
    int nums[] = {3,1,4,1,5,9};
    int num_to_find = 9;
    int start = 0;
    int end = 5;
    int* result = find(nums+start, nums+end, num_to_find);
    if( result == nums+end ){
        cout<< "Did not find any number matching "<< num_to_find << endl;
    }else{
        cout<< "Found a matching number:" << *result << endl;
    }
    return 0;
}
*/


#include<vector>
#include<iostream>
#include<stack>
#include<map>
#include<queue>
#include <algorithm>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    TreeNode* getRoot(vector<int>&preoder, vector<int>& inorder,int firPre, int endPre, int firIn, int endIn){
        TreeNode* root;
        int mid = preoder[firPre];
        vector<int>::iterator f = find(inorder.begin() + firIn, inorder.begin() + endIn + 1, mid);
        int  l = distance(inorder.begin()+firIn, f);

        root = new TreeNode(mid);
        root->left = getRoot(preoder, inorder, firPre+1, firPre+l, firIn, firIn+l-1);
        root->right = getRoot(preoder, inorder, firPre+1+l, endPre, firIn+l+1, endIn);
        return root;
    }
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        TreeNode*result;
        result=getRoot(preorder,inorder, 0, preorder.size()-1, 0, inorder.size()-1);
        return result;
    }
};
