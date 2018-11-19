#include<vector>
#include<algorithm>
#include<string>
#include<iostream>
using namespace std;

bool compare (int a, int b){
    string str1 = to_string(a);
    string str2 = to_string(b);
    return (str1+str2 > str2+str1);
}
class Solution {
public:
    string largestNumber(vector<int>& nums) {
        sort(nums.begin(),nums.end(),compare);
        string result;
        for (int i =0; i<nums.size(); i++){
            result+=to_string(nums[i]);
        }
        while(result.size()!=0 && result[0]=='0'){
            result.erase(result.begin());
        }
        if (result.size() == 0){
            return "0";  
        }
        return result;    
    }
};

