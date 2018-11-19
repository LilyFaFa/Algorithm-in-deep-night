#include<iostream>
#include<vector>
#include<string>
using namespace std;

class Solution {
public:
    bool exist(vector< vector <char> >& board, string word) {
        if (board.size() == 0) {
            return false;
        }

        h = board.size();
        w = board[0].size();

        for (int i=0;i<h;i++){
            for (int j=0;j<w;j++){
                if (search(board,word,0,i,j)){
                    return true;
                }
            }
        }
        return false;
    }
private:
    int h;
    int w;
    bool search(vector< vector <char> > &board, string target,int cur,int m, int n){
        if ( cur == target.length() ){
            return true;
        }
        if ( m >= board.size() || m <0 || n>=board[0].size() || n<0 ){
            return false;
        }
        if ( board[m][n]== target[cur] ) {
            char tmp;
            tmp=board[m][n];
            board[m][n]='0';
            if (search(board,target,cur+1,m+1,n)){
                return true;
            }
            if (search(board,target,cur+1,m,n+1)){
                return true;
            }
            if (search(board,target,cur+1,m-1,n)){
                return true;
            }
            if (search(board,target,cur+1,m,n-1)){
                return true;
            }
            board[m][n]=tmp;
        }
        return false;
    }
};


int main()
{
    vector<vector <char> > tmp;
    vector<char>t;
    tmp.push_back(t);
    tmp.push_back(t);
    t.push_back('a');
    t.push_back('b');
    cout<< t.size() << endl;
    cout<< tmp.size() << endl;
    cout<< tmp[0].size()<<endl;
    cout<< t[0]<< endl;
    tmp[0].push_back('k');
    cout<< tmp[0][0] << endl;
    return 0;
}