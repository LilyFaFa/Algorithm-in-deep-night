//海岛面积计算题，给一个矩阵，0表示海水，相连的1表示海岛，上下左右表示相连。
// 00011
// 11011
// 10001
//（1）求最大海岛面积。
//（2）求最大海岛面积和对应海岛的所有坐标。
//（3）求所有海岛的所有坐标，按海岛分。

// 每次遍历一个点标记为某个数值
#include<iostream>
#include<vector>
using namespace std;
void GetMax(vector< vector<int> >& nums, pair<int,int>curPoint, int cur , int& result){
    if( cur > result ){
        result=cur;
    }
    if (curPoint.first>=nums.size() || curPoint.second>=nums[0].size()){
        return;
    }
    if ( nums[curPoint.first][curPoint.second] == 1){
        nums[curPoint.first][curPoint.second] == 0;
        pair<int,int>nextPoint = curPoint;
        cur++;
        nextPoint.first++;
        GetMax(nums,nextPoint,cur,result);
        nextPoint.first--;
        nextPoint.second++;
        GetMax(nums,nextPoint,cur,result);
        nextPoint.second--;
        nextPoint.second--;
        GetMax(nums,nextPoint,cur,result);
    }
}
int main(){
    vector< vector<int> >nums;
    int result;
    for (int i=0; i<nums.size();i ++){
        for (int j=0; j<nums.size(); j++){
            if (nums[i][j]==1){
                int curResult =0;
                GetMax(nums,pair<int,int>(i,j),1,curResult);
                if (curResult>result){
                    result=curResult;
                }
            }
        }
    }
    return 0;
}
