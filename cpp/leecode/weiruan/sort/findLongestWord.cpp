#include<iostream>
#include<vector>
#include<string>
using namespace std;

class Solution {
public:
    string findLongestWord(string s, vector<string>& d) {
        string result = "";
        for (int i=0; i<d.size(); i++){
            int sd=0;
            int ss=0;
            while(sd<d[i].size()&&ss<s.size()){
                if (d[i][sd]==s[ss]){
                    sd++;
                }
                ss++;
            }
            if (sd == d[i].size()){
                if(result.size()<d[i].size() || result.size()==d[i].size() && result>d[i]){
                    result=d[i];
                }
            }
        }
        return result;
    }
};