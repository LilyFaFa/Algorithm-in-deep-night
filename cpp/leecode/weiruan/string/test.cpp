#include<iostream>
#include<vector>
#include<string>
using namespace std;

struct TreeNode{
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x):val(x),left(NULL),right(NULL){}
};

int main(){
    int i = 123;
    string str = to_string(i);
    cout<<str<<endl;
    int a = stoi(str);
    cout<<a<<endl; 
    TreeNode* t= new TreeNode(1);
    cout << t->val;
    return 0;
}