// 两个字符串中最长公共子序列
// 子序列是可以不连续的
// 使用动态规划
#include<vector>
#include<iostream>
using namespace std;
int GetMax(vector<int>num1, vector<int>num2){
    int result=0;
    vector< vector<int> >tmp(num1.size(), vector<int>(num2.size(), 0));
    if (num1[0]== num2[2]){
        tmp[0][0]=1;
    }
    tmp[0][1]=tmp[0][0];
    tmp[1][0]=tmp[0][0];
    
    for (int i=1; i<num1.size(); i++){
        for(int j=1; j<num2.size(); j++){
            if (num1[i] == num2[j]){
                tmp[i][j]=tmp[i-1][j-1]+1;
            }else{
                if (tmp[i-1][j]>tmp[i][j-1]){
                    tmp[i][j]=tmp[i-1][j]+1;
                }else{
                    tmp[i][j]=tmp[i][j-1]+1;
                }
            }
        }
    }
    return(tmp[num1.size()][num2.size()]);
}
