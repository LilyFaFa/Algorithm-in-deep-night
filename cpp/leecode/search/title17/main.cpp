#include<vector>
#include<iostream>
#include<string>
using namespace std;

class Solution {
public:
    vector<string> letterCombinations(string digits) {
        vector<string> result;
        if (digits.empty()){
            return result;
        }
        vector< vector <char> > d(10);
        d[2].push_back('a');
        d[2].push_back('b');
        d[2].push_back('c');

        d[3].push_back('d');
        d[3].push_back('e');
        d[3].push_back('f');

        d[4].push_back('g');
        d[4].push_back('h');
        d[4].push_back('i');

        d[5].push_back('j');
        d[5].push_back('k');
        d[5].push_back('l');

        d[6].push_back('m');
        d[6].push_back('n');
        d[6].push_back('o');

        d[7].push_back('p');
        d[7].push_back('q');
        d[7].push_back('r');
        d[7].push_back('s');

        d[8].push_back('t');
        d[8].push_back('u');
        d[8].push_back('v');

        d[9].push_back('w');
        d[9].push_back('x');
        d[9].push_back('y');
        d[9].push_back('z');
        string cur;
        dfs(result,d,cur,digits,0);
        return result;
    }
private:
    void dfs(vector<string>& results, vector< vector <char> > d,string cur, string digits, int index){
        if ( index == digits.length() ){
            results.push_back(cur);
            return;
        }

        for (const char c : d[digits[index] - '0']){
            cur.push_back(c);
            dfs(results,d,cur,digits,index+1);
            cur.pop_back();
        }
    }
};
