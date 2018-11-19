#include<iostream>
#include<vector>
#include<map>
using namespace std;
class Solution {
public:
    int rob(vector<int>& nums) {
        // 1、维护两个vector，一个对第一个房间进行偷盗，另一个从第二个房间进行偷盗；
        if (nums.size() ==0 ){
            return 0;
        }
        vector<int> indp(nums.size(),0);
        vector<int> outdp(nums.size(),0);

        // 2、不包含第一个元素，最后一个元素可以包含也可以不包含
        outdp[0] = 0;
        outdp[1] = nums[1];
        for(int i=2; i<nums.size(); i++){
            outdp[i] = max( outdp[i-1], outdp[i-2]+nums[i]);
        }
        
        // 3、包含第一个元素，不包含最后一个元素，将最后一个元素设置为负数
        indp[0] = nums[0];
        indp[1] = nums[0];
        nums[nums.size()-1]=-1;
        for(int i=2; i<nums.size(); i++){
            indp[i] = max( indp[i-1], indp[i-2]+nums[i]);
        }
        
        if (indp[nums.size()-1]>outdp[nums.size()-1]){
            return indp[nums.size()-1];
        }else{
            return outdp[nums.size()-1];
        }

    }
};