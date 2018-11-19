#include<vector>
#include<iostream>
#include<string>
#include<algorithm>
using namespace std;

class Solution {
public:
    bool searchMatrix(vector<vector<int> >& matrix, int target) {
        if (matrix.size()==0){
            return false;
        }
        int curX=0;
        int curY=matrix[0].size();
        while(curX<matrix.size() && curY>-1){
            if (matrix[curX][curY]==target){
                return true;
            }else if(matrix[curX][curY]>target){
                curY--;
            }else{
                curX++;
            }
        }
        return false;
    }
};