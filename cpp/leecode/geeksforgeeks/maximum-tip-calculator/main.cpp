#include <iostream>
#include <vector>
#include <algorithm>
#include <math.h>
using namespace std;

bool compare(pair<int,int> x, pair<int,int> y)
{
    return ((x.first-x.second) > (y.first-y.second));
}

int main() {
	int t;
	cin>>t;
	int n,x,y;
	while(t--){
	    cin>>n>>x>>y;
	    vector<pair<int,int> >a(n);
	    for(int i=0; i<n; i++) {
            cin>>a[i].first;
        }
	    for(int i=0; i<n; i++){
            cin>>a[i].second;
        } 
	    sort(a.begin(),a.end(),compare);
	    int ans=0;
	    for(int i=0; i < n; i++){
            if ( x>0 && a[i].first>=a[i].second || y<(n-i)){
                ans+=a[i].first;
                x--;
            }else {
                ans+=a[i].second;
            }
	  }
	  cout<<ans<<endl;    
	}
	return 0;
}
