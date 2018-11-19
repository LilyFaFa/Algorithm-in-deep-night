#include <iostream>
#include <map>
#include <string>

using namespace std;

int main(){
	map<int,string> m;
	m[1]="hello";
	m[2]="world";
	cin>>m[3];
	cout<<m[1]<<"\n";
	cout<<m[3]<<"\n";
	// find函数，查找
	if ( m.find(2) != m.end()){
		cout<<m[2];
	}
	return 0;
}