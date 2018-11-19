#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;

class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};

class Solution {
public:
    void getDepth(Node* root, int& depth,int cur){
        cur++;
        if (cur > depth){
            depth=cur;
        }
        for (int i=0; i<root->children.size(); i++){
            getDepth(root->children[i],depth,cur);
            cout<<cur;
        }
    }
    int maxDepth(Node* root) {
        if (root == NULL){
            return 0;
        }
        int depth =0;
        int cur=0;
        getDepth(root,depth,cur);
        return depth;
    }
};