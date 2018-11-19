#include<vector>
#include<iostream>
using namespace std;
class Solution {
public:
    int findNumberOfLIS(vector<int>& nums) {
        if (nums.size()==0){
            return 0;
        }
        vector<int>result(nums.size(),1);
        vector<int>tmp(nums.size(),1);
        for (int i=1; i< nums.size(); i++){
            for (int j=0; j<i; j++){
                if (nums[j]<nums[i]){
                    if (result[j]+1>result[i]){
                        result[i]=result[j]+1;
                        tmp[i]=tmp[j];
                    }else if (result[j]+1 == result[i]){
                        tmp[i]=tmp[j]+tmp[i];
                    }
                    
                }
            }
        }
        int max =0;
        int fre =0;
        for (int i=0; i<result.size(); i++){
            cout<<result[i]<<tmp[i]<<endl;
            if (result[i]>max){
                max=result[i];
                fre=tmp[i];
            }else if (result[i]==max){
                fre=fre+tmp[i];
            }
        }
        return fre;
    }
};