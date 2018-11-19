#include <iostream>
#include<vector>
#include <algorithm>
using namespace std;

void findPath(vector<vector<long long> >& edges, vector<long long>& visited, int start, int cur){
    if (visited[cur] == 0){
        visited[cur] = 1;
    }else{
        return;
    }
    for (int i=1; i<edges.size(); i++){
        if ( edges[i][cur]==1 && visited[i] == 0){
            findPath(edges, visited, start, i);
        }
    }
    return;
}
int main() {
	int n;
    cin>>n;
    while(n--){
        int v,e;
        cin>>v>>e;
        vector<vector<long long> >edges;
        for (int i=0; i<v+1; i++){
            vector<long long>edge(v+1,0);
            edges.push_back(edge);
        }
        for (int i=0;i<e;i++){
            int a,b;
            cin>>a>>b;
            edges[a][b]=1;
        }

        int result=0;
        vector<long long>visited (v+1, 0);
        for (int i=1; i<v+1; i++){
            // 将vector置为某一个值
            fill(visited.begin(), visited.end(),0);
            findPath(edges, visited, i, i);
            int num = count( visited.begin(), visited.end(), 1);
            if ( num >= v){
                result++;
            }
        }
        cout<<result<<endl;
    }
	return 0;
}