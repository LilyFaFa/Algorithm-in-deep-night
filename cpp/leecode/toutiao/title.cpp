// 红蓝两种球，总共N个， N>2, 排列组合，连续3个颜色一样是非法的，求合法的排列数量

// 使用的是递归算法,搜索所有的可能结果
#include<iostream>
#include<vector>
using namespace std;
void GetResult(int N, int cur, int& result,int colorNum){
    if (colorNum == 3 ){
        return;
    }
    if (cur == N){
        result++;
        return;
    }
    GetResult(N, cur+1, result, colorNum+1);
    GetResult(N, cur+1, result, 1);
}

int main(){
    int result=0;
    int N=0;
    GetResult(N,0,result,1);
    result=result*2;
    return 0;
}