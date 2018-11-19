#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;

class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        if ( nums.size() ==0 ){
            return 0;
        }
        int begin=0;
        int next=0;
        int minCur = nums.size()+1;
        int curSum = nums[0];

        while ( next < nums.size() ){
            if( curSum >= s ){
                if( minCur > (next-begin+1) ){
                    minCur = next-begin+1;
                }
                curSum-=nums[begin];
                begin++;
            } else {
                next++;
                if (next<nums.size()){
                    curSum+=nums[next];
                }
            }
        }
        if (minCur == nums.size()+1){
            return 0;
        }
        return minCur;
    }
};