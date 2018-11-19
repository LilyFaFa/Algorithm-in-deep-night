#include<vector>
#include<iostream>
#include<algorithm>
#include<map>
#include<string>
using namespace std;

string getXunHuan(int a, int b){
    int flag = 0;
    string result = "";
    map<int,bool>tmp;
    flag = (a*10)%b;
    
    while( tmp.find(flag) == tmp.end() ){
        tmp[flag] = true;
        result += to_string(flag);
        flag = (flag*10)%b;
        cout<<"cur"<<endl;
    }
    return result;
}
int main(){
    string result = getXunHuan(1,7);
    cout<<result<<endl;
    return 0;
}