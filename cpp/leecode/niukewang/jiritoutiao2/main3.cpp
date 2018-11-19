#include<iostream>
#include<string>
#include<vector>
#include<map>
#define MAX 100000
using namespace std;

int check(vector<int> positions, int maxChangeTime){
    // 1. 二维vector创建
    vector< vector<int> >dp;
    for(int i=0; i<positions.size(); i++){
        vector<int> ps(positions.size(), MAX);
        dp.push_back(ps);
    }
    for (int i=0; i<positions.size(); i++){
        dp[i][i]=0;
        if (i<positions.size()-1){
            dp[i][i+1]= positions[i+1] - positions[i]-1;
        }
    }

    // 2. 从连续个数为1开始动态规划
    int dis = 2;
    while(dis < positions.size()){
        for ( int i=0; i<positions.size()-dis; i++ ){
            dp[i][i+dis] = dp[i+1][i+dis-1] + positions[i+dis] - positions[i] - dis;
        }
        dis++;
    }
    
    int max=-1;
    for (int i=0; i<dp.size(); i++){
        for (int j=i; j<dp.size(); j++){
            if (dp[i][j]<= maxChangeTime && (j-i+1)>max){
                max=j-i+1;
            }
        }
    }
    return max;
}

int main(){
    string str;
    int maxChangeTime;
    cin >> str;
    cin >> maxChangeTime;
    map<char, vector<int> > positions;
    for (int i=0; i < str.size(); i++){
        positions[str[i]].push_back(i);
    }
    int result = 0;
    for (map<char, vector<int> >::iterator iter = positions.begin(); iter != positions.end(); iter++){
        vector<int>cur = iter->second;
        int r = check(cur,maxChangeTime);
        cout<<r<<iter->first<<"\n";
        if (r >result){
            result=r;
        }
    }
    cout<<result<<endl;

}