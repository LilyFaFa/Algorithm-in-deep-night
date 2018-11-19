#include<iostream>
#include<vector>
#include<queue>
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

class Solution {
public:
    bool isSymmetrical(TreeNode* pRoot)
    {
       if (pRoot == NULL){
           return true;
       }
       queue<TreeNode*> curqueue;
       stack<TreeNode*> curstack;
       curqueue.push(pRoot);
       curstack.push(pRoot);
       bool result=true;
       cengce(curqueue,curstack, &result);
       return result;
    }
private:
    void cengce(queue<TreeNode*> curqueue, stack<TreeNode*> curstack, bool* sym){
        if (*sym ==false){
            return;
        }
        int len = curqueue.size();
        if (len == 0){
            return;
        }
        for(int i=0; i<len; i++){
            TreeNode* node1 = curqueue.front();
            TreeNode* node2 = curstack.top();
            curqueue.pop();
            curstack.pop();
            if (node1->left == NULL){
                if (node2->right != NULL){
                    *sym=false;
                    break;
                } 
            }else{
                if (node2->right == NULL){
                    *sym=false;
                    break;
                }else{
                    if (node1->left->val !=node2->right->val){
                        *sym=false;
                        break;
                    }
                }
                curqueue.push(node1->left);
            }

            if (node1->right == NULL){
                if (node2->left != NULL){
                    *sym=false;
                    break;
                } 
            }else{
                if (node2->left == NULL){
                    *sym=false;
                    break;
                }else{
                    if (node1->right->val != node2->left->val){
                        *sym=false;
                        break;
                    }
                }
                curqueue.push(node1->right);
            }
            
        }
        if (*sym ==true){
            int len = curqueue.size();
            for (int i=0;i<len;i++){
                curstack.push(curqueue.front());
                curqueue.push(curqueue.front());
                curqueue.pop();
            }
            cengce(curqueue,curstack,sym);
        }
    }

};