#include<iostream>
#include<vector>
#include<stack>
#include<set>
#include<map>
/*
Please note that it's Function problem i.e.
you need to write your solution in the form of Function(s) only.
Driver Code to call/invoke your function would be added by GfG's Online Judge.*/


//User function Template for C++
/*The function should return a string denoting the 
order of the words in the Alien Dictionary */
#define pb push_back
using namespace std;
void topo(map<int, vector<int> > &mp, stack<int>& st, int u, set<int> &vs){
    if(vs.find(u) != vs.end()) return;
    vs.insert(u);
    for(int i = 0; i < mp[u].size(); i++){
        int v = mp[u][i];
        topo(mp, st, v, vs);
    }
    st.push(u);
}
bool hasCircle(map<int, vector<int> > &mp, int u, set<int>& vs, set<int>& st){
    if(vs.find(u) == vs.end()){
        vs.insert(u);
        st.insert(u);
        if(mp.find(u) != mp.end()){
            for(int i = 0; i < mp[u].size(); i++){
                int v = mp[u][i];
                if(vs.find(v) == vs.end() && hasCircle(mp, v, vs, st)){
                    return true;
                }else if(st.find(v) != st.end()){
                    return true;
                }
            }
        }
    }
    st.erase(u);
    return false;
}
string printOrder(string dict[], int N, int k)
{
   //Your code here
   map<int, vector<int> > mp;
   for(int i = 0; i < N; i++){
       string s = dict[i];
       for(int j = 0; j < s.size(); j++){
           char c = s[j];
           if(mp.find(c) == mp.end()){
               mp.insert(pair<int, vector<int> >(c, vector<int>()));
           }
       }
   }
   for(int i = 0; i < N - 1; i++){
       string s1 = dict[i], s2 = dict[i + 1];
       for(int j = 0; j < min(s1.size(), s2.size()); j++){
           if(s1[j] != s2[j]){
               mp[s1[j]].pb(s2[j]);
               break;
           }
       }
   }
   set<int> vs; set<int> rec;
   for(auto it = mp.begin(); it != mp.end(); it++){
       if(hasCircle(mp, it->first, vs, rec)){
           return "";
       }
   }
   vs.clear();
   stack<int> st;
  for(auto it = mp.begin(); it != mp.end(); it++){
      int u = it->first;
      if(vs.find(u) == vs.end()){
          topo(mp, st, u, vs);
      }
  }
  string rs = "";
  while(!st.empty()){
      rs += st.top();
      st.pop();
  }
   return rs;
}

int main(){
    
}