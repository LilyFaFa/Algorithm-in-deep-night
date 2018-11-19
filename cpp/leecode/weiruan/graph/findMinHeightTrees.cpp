#include<iostream>
#include<vector>
#include<set>
#include<map>
using namespace std;

class Solution {
public:
    vector<int> findMinHeightTrees(int n, vector<pair<int, int> >& edges) {
        set<int> point;
        map<int,int> nums;
        for (int i=0; i<edges.size(); i++){
            if (nums.find(edges[i].first) !=nums.end()){
                nums[edges[i].first]++;
            }else{
                nums[edges[i].first]=1;
                point.insert(edges[i].first);
            }

            if (nums.find(edges[i].second) !=nums.end()){
                nums[edges[i].second]++;
            }else{
                nums[edges[i].second]=1;
                point.insert(edges[i].second);
            }
        }
        map<int,int>tmp;
        while(point.size()>2){
            for (map<int,int>::iterator it =tmp.begin(); it!=tmp.end(); it++){
                nums[(*it).first]=nums[(*it).first]-(*it).second;
            }
            tmp.clear();
            for(int i =0; i<edges.size(); i++){
                if (nums[edges[i].first]==1 && nums[edges[i].second]>1 && tmp.find(nums[edges[i].first])==tmp.end()){
                    nums[edges[i].first]--;
                    point.erase(edges[i].first);
                    if (tmp.find(edges[i].second)==tmp.end()){
                        tmp[edges[i].second]=1;
                    }else{
                        tmp[edges[i].second]++;
                    }
                }else if (nums[edges[i].second]==1 && nums[edges[i].first]>1 && tmp.find(nums[edges[i].second])==tmp.end()){
                    nums[edges[i].second]--;
                    point.erase(edges[i].second);
                    if (tmp.find(edges[i].second)==tmp.end()){
                        tmp[edges[i].second]=1;
                    }else{
                        tmp[edges[i].second]++;
                    }
                }
            }
            
        }
        vector<int>result;
        for (set<int>::iterator ite =point.begin(); ite!=point.end(); ite++){
            result.push_back(*ite);
        }
        return result;
    }
};