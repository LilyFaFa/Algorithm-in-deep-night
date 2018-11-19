#include<iostream>
#include<vector>

using namespace std;

class Solution {
public:
    void reOrderArray(vector<int> &array) {
        if (array.size()==0 || array.size()==1){
            return ;
        }
        int index=0;
        for (int i=0; i<array.size(); i++){
            if ( array[i]%2 ==0 ){
                index++;
            }else{
                int tmp = array[i];
                for (int j=0; j<index; j++){
                    array[i-j+1] = array[i-j];
                }
                array[i-index]=tmp;
            }
        }
    }
};