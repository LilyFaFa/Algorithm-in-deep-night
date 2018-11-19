#include<vector>
#include<iostream>
using namespace std;
// 大顶堆排序
// 对cur点进行一个下沉
void maxHeapDown(vector<int>&heap, int cur, int end){
    while( cur <=end){
        int left=cur*2+1;
        int right=cur*2+2;
        if ( left <=end && right <=end){
            if (heap[cur]>=heap[left] && heap[cur]>=heap[right]){
                break;
            }else if (heap[left]<heap[right]){
                int tmp=heap[right];
                heap[right]=heap[cur];
                heap[cur]=tmp;
                cur=right;
            }else{
                int tmp=heap[left];
                heap[left]=heap[cur];
                heap[cur]=tmp;
                cur=left;
            }
        }else if (left <=end && heap[cur]<heap[left]){
            int tmp=heap[left];
            heap[left]=heap[cur];
            heap[cur]=tmp;
            cur=left;
        }
    }
}

// 升序
void  heapSortAsc(vector<int>& nums){
    // 1 初始化创建堆
    for (int i=nums.size()/2-1; i>=0 ;i--){
        maxHeapDown(nums, i, nums.size()-1);
    }
    // 2 调整堆
    for (int i=nums.size()-1; i>=0; i--){
        int tmp = nums[0];
        nums[0] = nums[i];
        nums[i] = tmp;
        maxHeapDown(nums,0,i);
    }
}