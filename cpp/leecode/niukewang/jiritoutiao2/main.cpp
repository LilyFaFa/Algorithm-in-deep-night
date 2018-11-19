// 用户评价

#include<iostream>
#include<vector>
#include<unordered_map>
#include<algorithm>
using namespace std;

int main(){
    int num;
    while(cin>>num){
        // vector是按序进去的，所以后面可以二分查找
        unordered_map<int, vector<int> > check;
        unordered_map<int, vector<int> >::iterator iter;
        for (int i=1; i<num+1; i++){
            int tmp;
            cin>>tmp;
            check[tmp].push_back(i);
        }
        int test,l,r,k;
        cin>>test;
        for (int i=0; i<test; i++){
            cin>>l>>r>>k;
            int result = 0;
            iter = check.find(k);
            if (iter == check.end()){
                cout<<result<<endl;
                continue;
            }
            vector<int> v = iter->second;
            vector<int>::iterator ln = lower_bound(v.begin(), v.end(),l);
            vector<int>::iterator rn = upper_bound(v.begin(), v.end(),r);
            cout<<*ln<<"\n";
            cout<<*rn<<"\n";
            cout << rn - ln << endl;
        }
    }
    return 0;
}
    