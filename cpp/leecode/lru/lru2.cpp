#include<iostream>
#include<map>
using namespace std;
struct CacheNode{
    int key;
    int value;
    CacheNode* pre;
    CacheNode* next;
};

class LRUCache{
private:
    map<int , CacheNode*>mp;
    int size;
    CacheNode* head;
    CacheNode* tail;
public:
    LRUCache(int capacity){
        size=capacity;
        head=NULL;
        tail=NULL;
    }

    int get(int key){
        map<int,CacheNode*>::iterator ite = mp.find(key);
        if (ite == mp.end()){
            return -1;
        }
        CacheNode* node = ite->second;
        remove(node);
        setHead(node);
        return node->value;
    }

    void set(int key, int value){
        CacheNode* node= new CacheNode;
        node->key=key;
        node->value=value;
        map<int,CacheNode*>::iterator ite =mp.find(key);
        if (ite == mp.end()){
            setHead(node);
        }
    }
    
    void remove(CacheNode* node){

    }

    void setHead(CacheNode* node){

    }
};