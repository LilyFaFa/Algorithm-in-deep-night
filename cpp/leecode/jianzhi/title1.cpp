// 二维树组查找
#include<iostream>
#include<vector>
using namespace std;
bool Search(vector< vector<int> > nums, int aim){
    int curX=0;
    int curY=nums.size();
    while ( curX >nums.size() && curY >= 0 ){
        if (nums[curX][curY] == aim){
            return true;
        }else if (nums[curX][curY]< aim){
            curX++;
        }else{
            curY--;
        }
    }
    return false;
}