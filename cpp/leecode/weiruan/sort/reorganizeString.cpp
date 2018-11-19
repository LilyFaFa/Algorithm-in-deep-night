#include<vector>
#include<iostream>
#include<map>
#include<stack>
using namespace std;

class Solution {
public:
    string reorganizeString(string S) {
        vector<int> tmp(26,0);
        int max =0;
        for (int i=0; i<S.size(); i++){
            tmp[S[i]-'a']++;
            if (tmp[S[i]-'a']>max){
                max=tmp[S[i]-'a'];
            }
        }
        if (max*2-1 > S.size()){
            return "";
        }
        vector<char>result(S.size(),'a');
        int cur = 0;
        for (int i=0; i<tmp.size(); i++){
            for (int j=0; j<tmp[i]; j++){
                result[cur]='a'+i;
                cur+=2;
                if (cur>=S.size()){
                    cur=1;
                }
            }
        }
        string str = "";
        for (int i=0; i<result.size(); i++){
            if (i >0 && result[i]==result[i-1]){
                 str=result[i]+str;
            }else{
                 str+=result[i];
            }
            
        }
        return str;
    }  
};
