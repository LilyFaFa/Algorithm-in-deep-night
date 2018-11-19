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

bool compare(Interval i1, Interval i2){
    return i1.start<i2.start;
}

// 从小到大排序
class Solution {
public:
    vector<Interval> merge(vector<Interval>& intervals) {
        vector<Interval>result;
        if (intervals.size()==0){
            return result;
        }

        // 定义一个比较函数
        sort(intervals.begin(),intervals.end(),compare);
        // 第一一个临时变量，进行合并
        Interval tmp;
        tmp.start = intervals[0].start;
        tmp.end = intervals[0].end;
        for (int i=1; i<intervals.size(); i++){
            if (intervals[i].start <= tmp.end){
                tmp.end= max(tmp.end, intervals[i].end);
            }else{
                result.push_back(tmp);
                tmp.start=intervals[i].start;
                tmp.end=intervals[i].end;
            }
        }
        result.push_back(tmp);
        return result;
    }
};