#include<iostream>
#include<vector>
#include<math.h>
using namespace std;

void GetYueShu(int num, vector<int>& yueshu){
    for (int i = 2; i<= sqrt(num); i++ ){
        if (num%i == 0){
            yueshu.push_back(i);
            yueshu.push_back(num/i);
        }
    }
    return;
}

int main(){
    int n,m;
    while( scanf("%d%d",&n,&m) !=EOF ){
        vector<int> base (m+1 ,m+1);
        base[n] = 0;
        for (int i=n; i<=m; i++){
            vector<int> curyue;
            GetYueShu(i,curyue);
            for (int j=0; j<curyue.size(); j++){
                if ( i+curyue[j] <= m){
                    if ( base[i+curyue[j]] > base[i]+1){
                        base[i+curyue[j]] = base[i]+1;
                    }  
                }
            }
        }
        if (base[m]==m+1){
            cout<<-1<<endl;
        }else{
            cout << base[m] << endl;
        }
        
    }
    return 0;
}