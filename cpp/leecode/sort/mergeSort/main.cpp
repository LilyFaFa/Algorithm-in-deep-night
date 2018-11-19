#include<vector>
#include<iostream>
using namespace std;

void merge(vector<int>& nums, int start, int middle, int end){
    vector<int>tmp;
    int start2=middle+1;
    int start1=start;

    while(start<=middle && start2<=end){
        if(nums[start] <= nums[start2]){
            tmp.push_back(nums[start]);
            start++;
        }else{
            tmp.push_back(nums[start2]);
            start2++;
        }
    }

    if (start <= middle){
        for (int i=start; i<=middle+1; i++){
            tmp.push_back(nums[i]);
        }
    }

    if (start2 <= end){
        for (int i=start2; i<=end; i++){
            tmp.push_back(nums[i]);
        }
    }

    for(int i=start1; i<=end; i++){
        nums[i] = tmp[i-start1];
    }
}

void mergeSort(vector<int>& nums, int start, int end){
    if (start >= end){
        return;
    }
    int middle = (start + end)/2;
    
    mergeSort(nums, start, middle);
    mergeSort(nums, middle+1, end);
    merge(nums, start, middle, end);
    
}

int main(){
    int a[8] = {3,9,7,8,2,1,4,5};
    vector<int> tests(a, a+8);
    mergeSort(tests,0,tests.size()-1);
    for (int i=0; i<tests.size(); i++){
        cout<<tests[i]<<"\n";
    }
}

