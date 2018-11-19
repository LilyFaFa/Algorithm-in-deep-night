#include<vector>
#include<iostream>
#include<algorithm>
#include<map>
#include<string>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        vector<ListNode*>tmp;
        for (int i=0 ;i<lists.size(); i++){
            tmp.push_back(lists[i]);
        }
        
    }
};