#include<iostream>
#include<vector>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
 
void Search(vector<int> visited, int curP, vector<TreeNode*> cur, vector<vector<TreeNode*> >& results){
    if (cur.size()==visited.size()){
        results.push_back(cur);
    }

    if ( ( curP>= visited.size() || visited[curP+1]==1) && (curP <=0 || visited[curP-1]==1)){
        return ;
    }

    if ( curP == 0 || visited[curP-1] == 1 ){
        cur.push_back(-1);
    }
    if ( curP == visited.size() || visited[curP+1] == 1 ){
        cur.push_back(-1);
    }

    for (int i = curP-1; i>=0; i--){
        visited[i]=1;
        cur.push_back(i);
        Search(visited,i,cur,results);
        cur.pop_back();
        visited[i]=0;
    }
    for (int j =curP+1; j<visited.size(); j++){
        visited[j]=1;
        cur.push_back(j);
        Search(visited,j,cur,results);
        cur.pop_back();
        visited[j]=0;
    }
    return;
}


class Solution {
public:
    vector<TreeNode*> generateTrees(int n) {
        int result = 0 ;
        vector<vector<int> > results;
        for (int i=0 ;i < n ; i++ ){
            vector<int> visited(n,0);
            vector<int> cur;
            cur.push_back(i);
            Search(visited,i,cur,results);
        }
        for (int i=0; i<results.size(); i++){
            for
        }
        cout<< result <<endl;
        return;
    }
};

