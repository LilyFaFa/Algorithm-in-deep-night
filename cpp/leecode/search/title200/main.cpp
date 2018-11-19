#include<iostream>
#include<vector>
#include<string>
using namespace std;

class Solution {
public:
    int numIslands(vector<vector<char> >& grid) {
        int result = 0 ;
        for ( int i=0; i < grid.size(); i++){
            for ( int j=0; j < grid[0].size(); j++){
                if ( grid[i][j] == '1' ){
                    result++;
                    dfs(grid,i,j);
                }
            }
        }
        return result;
    }
private:
    void dfs(vector<vector<char> >& grid,int m,int n){
        if (m >= grid.size() || n >= grid[0].size() || m<0 || n<0 ){
            return;
        }
        if (grid[m][n] == '1'){
            grid[m][n] = '0';
            dfs(grid,m+1,n);
            dfs(grid,m,n+1);
            dfs(grid,m-1,n);
            dfs(grid,m,n-1);
        }
    }
};