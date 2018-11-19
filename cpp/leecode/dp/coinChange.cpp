#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;

#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;

#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;

class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
        sort(coins.begin(),coins.end());
        int max=amount/coins[0]+1;
        vector< vector<int> >tmp(coins.size()+1, vector<int>(amount+1,max));
        for (int i=0; i<coins.size()+1; i++){
            tmp[i][0]=0;
        }
        
        for (int i=1; i<coins.size()+1; i++){
            for(int j=1; j<amount+1; j++){
                if(j>=coins[i-1]){
                    tmp[i][j]=min(tmp[i][j-coins[i-1]]+1,tmp[i-1][j]);
                }else {
                    tmp[i][j]=tmp[i-1][j];
                }
            }
        }
        if (tmp[coins.size()][amount]==amount/coins[0]+1){
            return -1;
        }else{
            return tmp[coins.size()][amount];
        }

    }
};