#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;

bool compare(int a,int b){
    return a>b;
}
class Solution {
public:
    void wiggleSort(vector<int>& nums) {
        sort(nums.begin(),nums.end(),compare);
        vector<int>v1(nums.begin(),nums.begin()+(nums.size()/2));
        vector<int>v2(nums.begin()+(nums.size()/2),nums.end());
        for(int i=0; i<v1.size();i++){
            nums[i*2]=v2[i];
            nums[i*2+1]=v1[i];
        }
        if(v1.size()<v2.size()){
            nums[nums.size()-1]=v2[v2.size()-1];
        }
    }
};

int main(){
    vector<int> nums;
    nums.push_back(1);
    nums.push_back(2);
    nums.push_back(3);
    nums.push_back(4);
    vector<int>v1(nums.begin(),nums.begin()+ceil(nums.size()/2));
    vector<int>v2(nums.begin()+ceil(nums.size()/2),nums.end());
    cout<<v1.size()<<endl;
    cout<<v2.size()<<endl;
}