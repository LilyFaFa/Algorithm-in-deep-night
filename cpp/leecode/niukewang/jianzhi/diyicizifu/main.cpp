#include<iostream>
#include<vector>
#include<string>
using namespace std;

class Solution {
public:
    int FirstNotRepeatingChar(string str) {
        vector<int>find(126,0);
        for (int i=0; i<str.size(); i++){
            find[str[i]]++;
        }
        for (int i=0; i<str.size(); i++){
            if  ( find[str[i]] == 1 ){
                return i;
            }
        }
        return -1;
    }
};