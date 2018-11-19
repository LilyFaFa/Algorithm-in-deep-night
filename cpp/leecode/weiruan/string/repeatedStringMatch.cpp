#include<iostream>
#include<vector>
#include<string>
using namespace std;

int main(){
    string str1 = "abcdf";
    string str2 = "bc";
    if ( str1.find(str2) != -1 ){
        cout<<str1.find(str2)<<endl;
    }
    
}