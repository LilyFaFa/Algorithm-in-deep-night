#include <iostream>
#include <vector>
using namespace std;
void findPath(int* result, vector<vector<int> >paths, vector<int>visited, int curPoint,int target){
    if (curPoint == target){
        (*result)++;
        return;
    }
    for (int i=0; i<paths.size(); i++){
        if ( paths[curPoint][i] == 1 && visited[i]==0 ){
            visited[i]=1;
            findPath( result, paths, visited, i, target);
            visited[i]=0;
        }
    }
    return;
}
int main() {
	int t;
    cin >> t;
    while(t--){
        int v,e;
        cin >>v>>e;
        vector<vector<int> > edges;
        for (int i=0; i<v; i++){
            vector<int> edge(v,0);
            edges.push_back(edge);
        }

        for (int j=0; j<e; j++){
            int a,b;
            cin>>a>>b;
            edges[a][b]=1;
        }
        int result=0;
        int start,end;
        cin>>start>>end;
        vector<int> visited(v,0);
        visited[start]=1;
        findPath(&result, edges,visited, start, end);
        cout<<result<<endl;
    }
	return 0;
}