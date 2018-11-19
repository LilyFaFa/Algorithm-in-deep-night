#include<string>
#include<vector>
#include<map>
/*
using namespace std;
class Solution {
public:
    bool isIsomorphic(string s, string t) {
        map<char, int> shash; 
        map<char, int> thash;
        for (auto i=0; i<t.length(); i++){
            if ( shash.count(s[i]) != thash.count(t[i])){
                return false;
            }else if ( shash.count( s[i] ) == 1 && thash.count( t[i] ) == 1){
                if (shash[s[i]] != t[i] || thash[t[i]] != s[i]){
                    return false;
                }
            }else{
                shash[s[i]]=t[i];
                thash[t[i]]=s[i];
            }
        }
        return true;
    }
};
*/
using namespace std;
class Solution {
public:
    bool isIsomorphic(string s, string t) {
        if (s.size() != t.size()){
            return false;
        }

        map<char,char> hash; 
        for (int i=0; i<s.size(); i++){
            if (hash.find(s[i]) != hash.end()){
                if (hash[s[i]] != t[i]){
                    return false;
                }
            }else{
                hash[s[i]]=t[i];
            }
        }
    }
};
