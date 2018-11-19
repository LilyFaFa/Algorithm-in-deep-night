#include<vector>
#include<iostream>
using namespace std;

class Solution {
public:
    bool canPartitionKSubsets(vector<int>& nums, int k) {
        int result = 0;
        for(int i=0; i<nums.size(); i++){
            result+=nums[i];
        }
        if ( result%k !=0 ){
            return false;
        }
        
    }
};