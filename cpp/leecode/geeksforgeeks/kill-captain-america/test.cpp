#include<vector>
#include<iostream>
#include<map>
#include<array>
#include <algorithm>

#include<unordered_map>

using namespace std;

using ll = long long;
#define all(A) A.begin(),A.end()
 
void dfs(unordered_map<ll,vector<ll> >&ma ,vector<bool> &visited,ll source)
{
    visited[source]=true;
    for(ll adj : ma[source])
    {
        if(!visited[adj])
            dfs(ma,visited ,adj);
    }
}
 
ll kosaraju(ll n ,unordered_map<ll,vector<ll> >&ma,unordered_map<ll,vector<ll> >&am)
{
    vector<bool> visited (n);
    ll mother =0;    
    for(ll i=0;i<n;i++)
    {
        if(!visited[i])
        {
            dfs(ma,visited,i);
            mother = i;
        }
    }
 
    fill(all(visited),false);
    dfs(ma,visited,mother);
    if(count(all(visited),true)!=n)
        return 0;
    fill(all(visited),false);
    dfs(am,visited,mother);
    return count(all(visited),true);
}
 
 
 
int main() {
    ll T;
    cin>>T;
    for(ll t =0;t<T;t++)
    {
        ll V,E;
        cin>>V>>E;
        unordered_map<ll,vector<ll> > ma,am;
        for(ll i=0;i<E;i++)
        {
            ll p,q;
            cin>>p>>q;
            ma[q-1].push_back(p-1);//transpose graph
            am[p-1].push_back(q-1);//actual
        }
        cout<<kosaraju(V,ma,am)<<endl;
    }
 
	//code
	return 0;
}