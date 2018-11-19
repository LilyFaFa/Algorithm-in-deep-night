#include<iostream>
#include<vector>
using namespace std;

int buy(int num){
    int result=0;
    result = (num/24) * 3;
    int next= num%24 ;

    vector< int > v6(24,24);
    vector< int > v8(24,24);
    v6[0]=0;
    v8[0]=0;
    v6[6]=1;
    v6[12]=2;
    v6[18]=3;
    v8[8]= 1;
    v8[16]=2;
    for (int i=0;i<v8.size();i++){
        if (i<8){
            v8[i]=v6[i];
            continue;
        }

        if (v8[i-8] < v6[i]){
            v8[i]=v8[i-8]+1;
        }else{
            v8[i]=v6[i];
        }
    }
    if (v8[next]==24){
        return -1;
    }
    result = result + v8[next];
    return result;
}

int main(){
    int a;
    while(scanf("%d",&a)!=EOF){
        a = buy(a);
        cout <<a <<endl;
    }
    return 0;
}