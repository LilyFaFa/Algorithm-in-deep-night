/*
   题目描述：
   一个int数组，里面数据无任何限制，要求求出所有这样的数a[i]，
   其左边的数都小于等于它，右边的数都大于等于它。能否只用一个额外数组和少量其它空间实现。
*/

#include<iostream>
#include<vector>
using namespace std;

vector<int> getResult(vector<int> nums){
    vector<int> result;
    if ( nums.size() == 0 ){
        return result;
    }

    //1. 从前向后遍历，记录到目前的最大值
    int maxCur = nums[0];
    vector<int> tmp;
    for(vector<int>::iterator iter=nums.begin();iter!=nums.end(); ++iter){
        if(*iter>maxCur){
            maxCur=*iter;
        }
        tmp.push_back(maxCur);
    }

    //2. 从后向前遍历，记录到目前的最小值
    int minCur ;
    minCur = nums[nums.size()-1]+1;
    for (int i=nums.size()-1; i>=0; i--){
        if(nums[i] < minCur && nums[i]==tmp[i]){
            result.push_back(nums[i]);
            minCur=nums[i];
        }
        
    }

    return result;
}
int main(){
    int a[8] = {3,5,4,6,7,8};
    vector<int> tests(a, a+6);
    vector<int>result=getResult(tests);
    for (vector<int>::iterator ite=result.begin();ite!=result.end(); ++ ite){
        cout<< *ite<<" ";
        cout<<endl;
    }
}

