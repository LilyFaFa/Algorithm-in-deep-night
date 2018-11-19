#include<vector>
#include<iostream>
#include<string>
#include<algorithm>
#include<map>
using namespace std;

struct Interval {
    int start;
    int end;
    Interval() : start(0), end(0) {}
    Interval(int s, int e) : start(s), end(e) {}
};

class Solution {
public:
    vector<Interval> insert(vector<Interval>& intervals, Interval newInterval) {
        vector<Interval>result;
        if (intervals[0].start>newInterval.end){
            result.push_back(newInterval);
            for (int i=0; i<intervals.size(); i++){
            result.push_back(intervals[i]);
            }
            return result;
        }
        int start = -1;
        int end = -1;
        for (int i=0; i<intervals.size(); i++){
            if (intervals[i].start<=newInterval.start && intervals[i].end>=newInterval.end){
                return intervals;
            }else if (newInterval.end>intervals[i].start){
                continue;
            }
        }
        

    }
};