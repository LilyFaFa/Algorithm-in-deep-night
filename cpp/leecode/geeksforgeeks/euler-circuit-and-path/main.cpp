#include<iostream>
#include<vector>
#include <list>
using namespace std;
class Graph
{
    int V;
    list<int> *adj;
public:
    Graph(int V)   {this->V = V; adj = new list<int>[V]; }
    ~Graph() { delete [] adj; }
    void addEdge(int v, int w);
    int isEulerian();
    bool isConnected();
    void DFSUtil(int v, bool visited[]);
};
 
void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
    adj[w].push_back(v);
}
 
void Graph::DFSUtil(int v, bool visited[])
{
    visited[v] = true;
    list<int>::iterator i;
    for (i = adj[v].begin(); i != adj[v].end(); ++i)
        if (!visited[*i])
            DFSUtil(*i, visited);
}
void test(Graph &g)
{
    int res = g.isEulerian();
    if (res == 0)
        cout << "0";    //none
    else if (res == 1)
        cout << "1";    //Euler path
    else
        cout << "2";    //Euler cycle
}
int main()
{
	int t;
	cin>>t;
	while(t--){
		int v,e;
		cin>> v>>e;
		Graph g1(v);
		for(int i=0;i<e;i++){
			int a,b;
			cin>>a>>b;
		    g1.addEdge(a, b);
		}
		test(g1);
	}
    return 0;
}


/*Please note that it's Function problem i.e.
you need to write your solution in the form of Function(s) only.
Driver Code to call/invoke your function is mentioned above.*/

/*Complete the function below*/
int Graph::isEulerian()
{
//add code here.
}