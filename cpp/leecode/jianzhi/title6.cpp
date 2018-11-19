// 面试题28 字符串的排列
// 输入一个字符串，打印字符串字母的全排列
#include<vector>
#include<iostream>
#include<string>
#include<set>
#include<map>
using namespace std;
void SearchCon(vector<vector <char> >&result, string str, vector<bool>used, vector<char>cur){
    
    if(cur.size() == str.size()){
        result.push_back(cur);
        return;
    }

    for (int i=0 ;i<used.size(); i++){
        if (!used[i]){
            cur.push_back(str[i]);
            used[i]=true;
            SearchCon(result,str,used,cur);
            cur.pop_back();
        }
    }
}

// 面试题30 最小的k个数字
// 利用set
vector<int> GetLeastNumbers(vector<int>nums, int k){
    vector<int>result;
    if (k ==0 || nums.size()==0){
        return result;
    } 
    set<int>tmp;
    for (int i=0; i<k; i++){
        tmp.insert(nums[i]);
    }
    for (int i= k; i<nums.size(); i++){
        int curMax = *(tmp.rbegin());
        if (curMax > nums[i] && tmp.find(nums[i])==tmp.end()){
            tmp.erase(curMax);
            tmp.insert(nums[i]);
        }
    }
    
    for (set<int>::iterator ite = tmp.begin(); ite!=tmp.end(); ite++){
        result.push_back(*ite);
    }
    return result;
}

// 面试题35 第一个只出现一次对的字符
char findChar(vector<char>inputs){
    map<char,int>check;
    for (vector<char>::iterator ite = inputs.begin(); ite!=inputs.end(); ite++){
        if (check.find(*ite) != check.end()){
            check[*ite]++;
        }else{
            check[*ite]=1;
        }
    }
    for(int i=0; i<inputs.size(); i++){
        if (check[inputs[i]]==1){
            return inputs[i];
        }
    }
}
int main(){
    int a[10]={3,4,2,5,2,6,2,8,0,1};
    vector<int>num(a,a+10);
    vector<int>result=GetLeastNumbers(num,2);
    for (vector<int>::iterator ite= result.begin(); ite !=result.end(); ite++){
        cout<<*ite<<endl;
    }
    return 0;
}