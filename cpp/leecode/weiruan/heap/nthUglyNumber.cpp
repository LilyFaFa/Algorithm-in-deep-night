#include<vector>
#include<iostream>
using namespace std;

class Solution {
public:
    int nthUglyNumber(int n) {
        vector<int> tmp;
        tmp.push_back(1);
        int n2 = 0;
        int n3 = 0;
        int n5 = 0;
        for (int i=1; i<n; i++){
            int c2 = tmp[n2]*2;
            int c3 = tmp[n3]*3;
            int c5 = tmp[n5]*5;
            int curMin = min(c2,min(c3,c5));
            tmp.push_back(curMin);
            if (c2 == curMin){
                n2++;
            }
            if (c3 == curMin){
                n3++;
            }
            if (c5 == curMin){
                n5++;
            }
        }
        return tmp[n-1];
    }
};