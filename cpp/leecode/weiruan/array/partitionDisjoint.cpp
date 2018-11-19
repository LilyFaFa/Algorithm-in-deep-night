#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;

class Solution {
public:
    int partitionDisjoint(vector<int>& A) {
        vector<int>tmpMax;
        int curMax=A[0];
        tmpMax.push_back(curMax);
        for (int i=1; i<A.size(); i++){
            if (A[i]>curMax){
                curMax=A[i];
            }
            tmpMax.push_back(curMax);
        }
        
        vector<int>tmpMin;
        int curMin=A[A.size()-1];
        tmpMin.push_back(curMin);
        for (int i=A.size()-1; i>=0; i--){
            if (A[i]<curMin){
                curMin=A[i];
            }
            tmpMin.push_back(curMin);
        }

        int result=0;
        for (; result<A.size()-1; result++){
            if (tmpMax[result] <= tmpMin[result+1]){
                break;
            }
        }
        return result+1;
    }
};