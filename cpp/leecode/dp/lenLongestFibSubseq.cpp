#include<iostream>
#include<vector>
#include<map>
using namespace std;


/*
class Solution {
public:
    int lenLongestFibSubseq(vector<int>& A) {
        vector< vector<int> >dp;
        // 1、初始化dp，设置为2，即斐波那契数列的前两个数字
        for (int i=0; i<A.size(); i++){
            vector<int>tmp(A.size(),2);
            dp.push_back(tmp);
        }

        // 2、动态规划计算
        int result=0;
        for (int k=2; k<A.size(); k++){
            for(int j=0; j<k; j++){
                for (int m=0; m<j; m++){
                    if (A[m] + A[j] == A[k] && dp[m][j] >= dp[j][k]){
                        //cout<<k<<"end";
                        dp[j][k] = dp[m][j] + 1;
                        if (result < dp[j][k]){
                            result = dp[j][k];
                        }
                    }
                }
            }
        }
        return result;
    }
};

*/

class Solution {
public:
    int lenLongestFibSubseq(vector<int>& A) {
        map<int, int>nums;
        int n = A.size();
        // 1、初始化dp，设置为2，即斐波那契数列的前两个数字
        vector< vector<int> >dp(n,vector<int>(n, 2));
        for (int i=0; i<n; i++){
            nums[A[i]] = i;
        }

        // 2、动态规划计算
        int result=0;
        for (int k=2; k<n; k++){
            for(int j=0; j<k; j++){
                if ( nums.find(A[k]-A[j]) != nums.end() && nums[A[k] - A[j]] < j && dp[A[k] - A[j]][j] >= dp[j][k]){
                        dp[j][k] = dp[A[k] - A[j]][j] + 1;
                        if (result < dp[j][k]){
                            result = dp[j][k];
                        }
                    }
                
            }
        }
        return result;
    }
};