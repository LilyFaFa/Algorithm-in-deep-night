// 手串
#include<iostream>
#include<vector>
#include<unordered_map>
#include<set>

using namespace std;

int Check ( unordered_map<int, vector<int> > shouchuan, int dis ){
    int result =0 ;
    for( unordered_map<int, vector<int> >::iterator iter = shouchuan.begin(); iter!=shouchuan.end(); iter++){
        vector<int>iter2 =iter->second;
        for (int i=0; i<iter2.size(); i++){
            int j=(i+1)%iter2.size();
            if (abs(iter2[i]-iter2[j])<dis){
                result++;
                cout<<iter->first<<"\n";
                break;
            }
        }
    }
    return result;
}
int main(){
    int n,m,c;
    unordered_map<int, vector<int> > shouchuan;
    while(cin>>n>>m>>c){
        for (int i=0; i<n; i++){
            int num;
            cin>>num;
            for(int j=0; j<num; j++){
                int c;
                cin>>c;
                shouchuan[c].push_back(i);
            }
        }
        int result=Check(shouchuan,m);
        cout<<result<<endl;
    }
    
    return 0;
}