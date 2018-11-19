#include<vector>
#include<iostream>
#include<string>
#include<algorithm>
#include<map>
using namespace std;

class Solution {
public:
    int subarraySum(vector<int>& nums, int k) {
        // 因为有正数也有负数，所以不能采用两个指针一前一后遍历
        // 1.方法一：使用hashmap存储前k项的和,本次使用方法一
        // 2.方法二：使用二维vector来保存所有的和；
        map<int,int>tmp;
        // 前k项和恰好为target
        tmp[0]=1;
        int curSum = 0;
        int result = 0;
        for (int i=0; i<nums.size(); i++){
            curSum+=nums[i];
            if ( tmp.find(curSum-k)!=tmp.end() ){
                result+=tmp[curSum-k];
            }
            if(tmp.find(curSum)!=tmp.end()){
                tmp[curSum]++;
            }else{
                tmp[curSum]=1;
            }
        }
        return result;
    }
};