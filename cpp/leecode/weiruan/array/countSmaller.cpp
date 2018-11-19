#include<vector>
#include<iostream>
using namespace std;

class Solution {
public:
    vector<int> countSmaller(vector<int>& nums) {
        vector<int>result(nums.size(),0);
        if (nums.size() == 0){
            return result;
        }
        result[nums.size()-1]=0;
        for (int i =nums.size()-2; i>=0; i--){
            int count=0;
            for (int j = i+1; j<nums.size(); j++){
                if (nums[j]<nums[i]){
                    count++;
                }
            }
            result[i]=count;
        }
        return result;
    }
};