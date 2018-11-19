#include<iostream>
#include<vector>
#include<cmath>

using namespace std;

int getMax(vector< vector<int> > nums){
    int max=0;
    
    for (int i=1; i<nums.size(); i++){
        for (int j=1; j<nums[0].size(); j++){
            int min1 = min(nums[i-1][j], nums[i][j-1]);
            int min2 = min(nums[i-1][j-1],min1)+1;
            if (nums[i][j]<min2){
                nums[i][j]=min2;
            }
            if (nums[i][j]>max){
                max=nums[i][j];
            }
        }
    }
    
    return pow(max,2);
}

int main(){
    int a1[6]={1,1,1,0,1,1};
    int a2[6]={1,1,1,0,0,1};
    int a3[6]={1,1,1,0,0,1};
    vector<int>v1 ;
    vector<vector<int> > vs;
    v1.insert(v1.begin(), a1, a1+6);
    vs.push_back(v1);
    v1.insert(v1.begin(), a2, a2+6);
    vs.push_back(v1);
    v1.insert(v1.begin(), a3, a3+6);
    vs.push_back(v1);
    int result=getMax(vs);
    cout<<result<<endl;
    return 0;
}