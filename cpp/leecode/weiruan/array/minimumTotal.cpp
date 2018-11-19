#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;

class Solution {
public:
    int minimumTotal(vector<vector<int> >& triangle) {
        vector<vector<int> >bp(triangle.size(),vector<int>(triangle.size(),0));
        bp[0][0] = triangle[0][0];
        for(int i=1 ;i<triangle.size(); i++){
            for(int j=0; j<i+1; j++){
                if ( j==0 ){
                    bp[i][j]=bp[i-1][j]+triangle[i][j];
                }else if (j ==i){
                    bp[i][j]=bp[i-1][j-1]+triangle[i][j];
                }else{
                    bp[i][j]=min(bp[i-1][j-1],bp[i-1][j])+triangle[i][j];
                }
            }
        }

        int result=bp[triangle.size()-1][0];
        for (int i=1; i<triangle.size(); i++){
            if (result>bp[triangle.size()-1][i]){
                result=bp[triangle.size()-1][i];
            }
        }
        return result;
    }
};