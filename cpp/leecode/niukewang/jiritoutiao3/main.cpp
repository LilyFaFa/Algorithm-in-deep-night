// 重新分房间
#include<iostream>
#include<vector>
#define MAX 100000
using namespace std;

vector<int> getInit(vector<int>cur, int home, int last){
    int min=MAX;
    int initHome=0;
    for (int i=0; i<cur.size(); i++){
        if (cur[(last+i)%cur.size()] <= min){
            min=cur[(last+i)%cur.size()];
            initHome = (last+i)%cur.size();
        }
    }
    
    for( int i=0; i<cur.size(); i++ ){
        if (i != initHome){
            cur[i]=cur[i]-min;
            cur[initHome]=cur[initHome]+min;
        }
    }
    
    if (initHome==(last-1)){
        return cur;
    }

    int i = initHome+1;
    while(i%cur.size()!=last){
        cur[i%cur.size()]--;
        cur[initHome]++;
        i++;
    }
    return cur;
}
int main(){
    int n,i;
    cin>>n>>i;
    vector<int>init(n,0);
    for (int j=0; j<n; j++){
        cin>>init[j];
    }
    vector<int>result = getInit(init,n,i);
    for (int j=0; j<n; j++){
        cout<<result[j]<<" ";
    }
    cout<<endl;
    return 0;
}