#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;
class Solution {
public:
    void rotate(vector<vector<int> >& matrix) {
        int circle = matrix.size();
	    for (int i = 0; i < circle/2; i++ ){
		    for (int j = i; j < circle-i-1; j++ ){
			    int tmp = matrix[i][j];
			    matrix[i][j] = matrix[circle-1-j][i];
			    matrix[circle-1-j][i] = matrix[circle-1-i][circle-1-j];
			    matrix[circle-1-i][circle-1-j] = matrix[j][circle-1-i];
			    matrix[j][circle-1-i] = tmp;
		    }
	    }
    }
};