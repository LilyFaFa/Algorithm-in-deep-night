#include<vector>
#include<iostream>
using namespace std;

// 冒泡排序，复杂度和数据分布状况没有关系
void BubbleSort(vector<int>& nums){
    if (nums.size()<2){
        return;
    }
    for(int i=0; i<nums.size(); i++){
        for (int j=nums.size()-1; j>i; j--){
            if (nums[j] < nums[j-1]){
                int tmp=nums[j];
                nums[j]=nums[j-1];
                nums[j-1]=tmp;
            }
        }
    }
    return;
}

// 选择排序，复杂度和数据分布状况没有关系
void SelectSort(vector<int>& nums){
    if (nums.size()<2){
        return;
    }
    for(int i=0; i<nums.size(); i++){
        for(int j=i+1; j<nums.size(); j++){
            if (nums[j]<nums[i]){
                int tmp=nums[i];
                nums[i]=nums[j];
                nums[j]=tmp;
            }
        }
    }
    return;
}

// 插入排序，复杂度和数据分布状况有关系
void InsertSort(vector<int>& nums){
    for (int i=1; i<nums.size(); i++){
        for (int j=i; j>0; j--){
            if( nums[j-1] > nums[j] ){
                int tmp=nums[j];
                nums[j]=nums[j-1];
                nums[j-1]=tmp;
            }
        }
    }
    return;
}

// 归并排序，好好想象函数栈
void MergeSort(vector<int>& nums, int start, int end){
    if (start >= end){
        return;
    }
    if (start-end == 1){
        if (nums[start] > nums[end]){
            int tmp=nums[start];
            nums[start]=nums[end];
            nums[end]=tmp;
        }
        return;
    }

    MergeSort(nums,start,(start+end)/2);
    MergeSort(nums,(start+end)/2+1,end);

    int start1=start;
    int start2=(start+end)/2+1;
    vector<int>tmp;
    while(start1 <= (start+end)/2 && start2 <=end){
        if (nums[start1]<nums[start2]){
            tmp.push_back(nums[start1]);
            start1++;
        }else{
            tmp.push_back(nums[start2]);
            start2++;
        }
    }
    for (int i=start1; i<=(start+end)/2; i++){
        tmp.push_back(nums[i]);
    }
    for (int i=start2; i<=end; i++){
        tmp.push_back(nums[i]);
    }
    for (int i=start; i<=end; i++){
        nums[i]=tmp[i-start];
    }
    return;
}

// 快速排序，好好想象函数栈
void QuickSort(vector<int>& nums, int start, int end){
    if (start >= end){
        return;
    }
    int begin=start+1;
    int fin = end;
    while( begin <= end ){
        while(nums[begin]<=nums[start]){
            begin++;
        }
        while(nums[end]>nums[start]){
            end--;
        }
        if(end>begin){
            int tmp=nums[begin];
            nums[begin]=nums[end];
            nums[end]=tmp;
            begin ++;
            end --;
        }  
    }
    int tmp=nums[end];
    nums[end]=nums[start];
    nums[start]=tmp;
    QuickSort(nums,start,end-1);
    QuickSort(nums,end+1,fin);
    return;
}
int main(){
    int a[6] = {7,2,4,6,1,5};
    vector<int>nums;
    nums.insert(nums.begin(), a, a+6);
    //BubbleSort(nums);
    //SelectSort(nums);
    //InsertSort(nums);
    QuickSort(nums,0,nums.size()-1);
    //MergeSort(nums,0,nums.size()-1);
    for (int i=0; i<nums.size(); i++){
        cout<<nums[i]<<" ";
    }
}