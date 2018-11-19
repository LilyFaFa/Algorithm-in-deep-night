#include<iostream>
#include<vector>

using namespace std;

void findCombination( vector<int> elements, bool* find, int sum, int cur, int num, int target){
    if (sum == target && num == 4){
        *find=true;
        return;
    }
    if (num > 4 || sum >target || cur>=elements.size()){
        return;
    }
    if (*find){
        return;
    }
    for (int i=cur; i<elements.size(); i++){
        findCombination(elements, find, sum+elements[i], i+1, num+1, target);
    }
}

int main() {
	int t;
	cin >> t;
	int n, target;
	while(t--){
	    cin>>n;
	    vector<int> a(n);
        bool result =false;
	    for(int i=0; i<n; i++) {
            cin>>a[i];
        }
        cin>>target; 
        findCombination(a,&result,0,0,0,target);
        if (result){
            cout<<1<<endl;  
        }else{
            cout<<0<<endl;
        }
	      
	}
	return 0;
}

