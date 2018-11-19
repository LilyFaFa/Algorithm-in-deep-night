#include<iostream>
#include<vector>
using namespace std;
void quickSort(vector<int>& nums, int start, int end){
    if (start >= end){
        return;
    }
    int tmp= nums[start];
    int i=start+1, j=end;
    while(i <= j){
        while ( nums[i]<=tmp && i<=end){
            i++;
        }
        while( nums[j]>tmp && j>start){
            j--;
        }

        if (i <= j){
            int t = nums[i];
            nums[i] = nums[j];
            nums[j] = t;
            i++;
            j--;
        }
        
    }
    if ( j > start ) {
        int t = nums[j];
        nums[j] = tmp;
        nums[start] = t;
    }

    quickSort(nums, start, j-1);
    quickSort(nums, j+1, end);
}
int main(){
    int a[8] = {3,9,7,8,2,1,4,5};
    vector<int> tests(a, a+8);
    quickSort(tests,0,tests.size()-1);
    for (int i=0; i<tests.size(); i++){
        cout<<tests[i]<<"\n";
    }
}