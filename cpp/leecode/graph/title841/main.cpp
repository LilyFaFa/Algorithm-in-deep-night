#include<vector>
#include<iostream>
#include<map>
#include<string>
#include<stack>
#include<set>
#include<queue>
using namespace std;
/*
class Solution {
public:
    bool canVisitAllRooms(vector< vector<int> >& rooms) {
        map<int, bool> v;
        //map<int, bool>& visited = v;
        dfs(v,rooms,0);
        //dfs(&v,rooms,0);
        if ( v.size() == rooms.size()){
            return true;
        }else{
            return false;
        }
    }
private:
    void dfs(map<int,bool>& visited,vector< vector<int> >& rooms,int room){
        if (visited.count(room) == 1){
            return;
        }
        visited[room]=true;
        if ( rooms[room].size() == 0){   
            return;
        }else{
            for (auto i=0; i<rooms[room].size(); i++){
                dfs(visited,rooms,rooms[room][i]);
            }
        }
    }
};
*/

// 利用递归
/*
void dfs(map<int , bool>& visited, vector<vector < int > >&rooms ,int room){
    if (visited.count(room)==1){
        return;
    }
    visited[room]=true;
    if (visited.size()==rooms.size()){
        return;
    }
    for (int i=0; i<rooms[room].size(); i++){
        dfs(visited,rooms,rooms[room][i]);
    }
}

bool visitAll(vector< vector<int> >& rooms){
    map<int, bool> v;
    dfs(v,rooms,0);
    if(v.size()==rooms.size()){
        return true;
    }
    return false;
}
*/

//利用栈
/*
bool visitAll(vector< vector<int> >& rooms){
    map<int, bool> v;
    stack<int> roomStack;
    roomStack.push(0);
    while(!roomStack.empty()){
        int cur=roomStack.top();
        v[cur] =true;
        roomStack.pop();
        for (int i=0; i<rooms[cur].size(); i++){
            if (v.find(rooms[cur][i]) ==v.end()){
                continue;
            }
            roomStack.push(rooms[cur][i]);
        }
    }
    if (v.size()==rooms.size()){
        return true;
    }else{
        return false;
    }
    
}
*/

//利用队列
bool visitAll(vector< vector<int> >& rooms){
    set<int>v;
    queue<int>q;
    q.push(0);
    while( !q.empty() ){
        int cur=q.front();
        v.insert(cur);
        q.pop();
        for(int i=0; i<rooms[cur].size(); i++){
            if (v.find(i)!=v.end()){
                q.push(rooms[cur][i]);
            }
        }
    }
}