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
    bool VerifySquenceOfBST(vector<int> sequence) {
        if (sequence.size()==0){
            return false;
        }
        bool result =true;
        checkBST(sequence,0,sequence.size()-1,&result);
        return result;
    }
private:
    void checkBST(vector<int>sequence,int first ,int last  ,bool* bst){
        if (*bst ==false){
            return;
        }
        if (first >= last){
            return;
        }

        int cur;
        for ( cur =first; cur<last; cur++){
            if ( sequence[cur] > sequence[last]){
               break;
            }
        }

        for (int i =cur; i<last; i++){
            if (sequence[i] < sequence[last]){
               *bst=false;
               break;
            }
        }
        if (*bst){
            checkBST(sequence,first,cur-1,bst);
            checkBST(sequence,cur+1,last,bst);
        }
        return;
    }
};