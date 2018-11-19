#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;

class Solution {
public:
    vector<int> largestDivisibleSubset(vector<int>& nums) {
        sort(nums.begin(),nums.end());

        // 记录最长的子序列长度
        vector<int>dp(nums.size(),1);
        // 记录当前节点最长的上一个位置
        vector<int>last(nums.size(),-1);
        int index=0;
        for (int i=1; i<nums.size(); i++){
            for (int j=0; j<i; j++){
                if ( nums[i]%nums[j] == 0 && dp[i] <= dp[j]){
                    dp[i] = dp[j]+1;
                    last[i] = j;
                    if ( dp[i] > dp[index] ){
                        index = i;
                    } 
                }
            }
        }
        vector<int>result;
        while ( index >= 0){
            result.push_back(index);
            index=last[index];
        }
    }
};