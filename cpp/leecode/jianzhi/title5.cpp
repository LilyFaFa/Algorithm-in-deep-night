#include<vector>
#include<iostream>
#include<string>
using namespace std;

// 题目12 打印长度<=n 的所有的整数
void PrintToMaxOfDigits(int n){
    if (n ==0 ){
        return ;
    }
    char* nums = new char[n+1];
    nums[n]='\0';
    cout <<*nums<<endl;
    cout <<"hello";
}
int main(){
    PrintToMaxOfDigits(3);
    return 0;
}