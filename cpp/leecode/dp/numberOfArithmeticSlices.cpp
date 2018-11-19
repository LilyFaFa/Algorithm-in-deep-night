#include<iostream>
#include<vector>
using namespace std;
class Solution {
public:
    int numberOfArithmeticSlices(vector<int>& A) {
        if (A.size()<3){
            return 0;
        }

        int result;
        int tmp=2;
        int cur = A[1]-A[0];
        for (int i=2; i<A.size(); i++){
            if ( A[i]-A[i-1] == cur ){
                tmp++;
            }else{
                if (tmp>=3){
                    result+=(1+tmp-2)*(tmp-2)/2;
                    cur=A[i]-A[i-1];
                    tmp=2;
                }
            }
        }
        if (tmp>=3){
            result+=(1+tmp-2)*(tmp-2)/2;
        }
        return result;
    }
};