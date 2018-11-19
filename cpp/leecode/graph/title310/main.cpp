#include<vector>
#include<string>
#include<map>
#include<iostream>
#include<utility>

using namespace std;

class Solution {
public:
    vector<int> findMinHeightTrees(int n, vector<pair<int, int> >& edges) {
        
    }
private:
    void dfs(int& depth, vector< vector<int> >& points, int root){
        return;
    };

    bool cmp(const pair<int,char> m, const pair<int ,char> n){
        if ( m.first > n.first){
            return true;
        }
        return false;
    };

    void test(pair<int, int>& pa, int a, int b){
        pa.push_back(make_pair(a,b));
        //pa.push_back(make_pair(a,b));
        //sort(pa.begin(),pa.end(),cmp);
    }
    

};