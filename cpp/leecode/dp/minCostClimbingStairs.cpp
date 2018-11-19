#include<iostream>
#include<vector>
using namespace std;
#define MAX_INT 10000000

class Solution {
public:
    int minCostClimbingStairs(vector<int>& cost) {
        cost.push_back(0);
        vector<int>dp(cost.size()+1, MAX_INT);
        dp[0]=cost[0];
        dp[1]=cost[1];
        for (int i=2; i<cost.size(); i++){
            if(dp[i-1]< dp[i-2]){
                dp[i]=dp[i-1]+cost[i];
            }else{
                dp[i]=dp[i-2]+cost[i];
            }
        }
        return dp[cost.size()-1];
    }
};