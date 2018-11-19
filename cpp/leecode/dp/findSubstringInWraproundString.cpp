#include<vector>
#include<iostream>
using namespace std;

class Solution {
public:
    int findSubstringInWraproundString(string p) {
        int result=0;
        // 将子序列合并起来处理
        vector<int>dp(p.size(),1);
        for ( int i=1; i<p.size(); i++ ){
            if ( p[i] == 'a'){
                if ( p[i-1] == 'z'){
                    dp[i] = dp[i-1]+1;
                    dp[i-1] = 0;
                }
            }else{
                if (dp[i]-dp[i-1]==1){
                    dp[i]=dp[i-1]+1;
                    dp[i-1] =0;
                }
            }
        }
        for (int i=0; i<dp.size(); i++){
            if (dp[i]!=0){
                result+=(dp[i]*(dp[i]+1))/2;
            }
        }
        return result;
    }
};