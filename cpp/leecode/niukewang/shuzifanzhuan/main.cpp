#include<iostream>
#include<string>
#include<map>
using namespace std;

int rev(int i){
    int cur = 0;
    int j = 0;
    while (i/10 != 0){
        j = i%10;
        i = i/10;
        cur = cur*10 + j;
    }
    if (i!=0){
        cur=cur*10+i;
    }
    return cur;
}

int main(){
    int a,b,c;
    while(scanf("%d%d",&a,&b)!=EOF){
        a = rev(a);
        b = rev(b);
        c = rev(a+b);
        cout <<c <<endl;
    }
    return 0;
}

