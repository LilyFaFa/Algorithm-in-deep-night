#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int rob(vector<int>& nums) {
        if ( nums.size() == 0 ){
            return 0;
        }

        vector<int> dp(nums.size(),0);
        dp[0]=nums[0];
        if ( nums[1] > nums[0] ){
            dp[1]=nums[1];
        }else{
            dp[1]=nums[0];
        }

        for(int i=2; i<nums.size(); i++){
            dp[i] = max( dp[i-1],dp[i-2]+nums[i]);
        }
        return dp[nums.size()-1];
    }
};