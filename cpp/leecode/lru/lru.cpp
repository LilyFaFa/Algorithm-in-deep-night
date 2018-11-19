#include<iostream>
#include<vector>
#include<map>
using namespace std;
// 基于双向链表和map来实现 lru 算法

struct CacheNode{
    int key;
    int value;
    CacheNode* pre;
    CacheNode* next;
};

class LRUCache{
private:
    int size;
    CacheNode* head;
    CacheNode* tail;
    map<int,CacheNode*>mp;
public:
    LRUCache (int capacity){
        size =capacity;
        head=NULL;
        tail=NULL;
    }

    // 获取某个值
    int get(int key){
        map<int,CacheNode*>::iterator ite = mp.find(key);
        if(ite != mp.end()){
            CacheNode* node = ite->second;
            remove(node);
            setHead(node);
            return node->value;
        }
        return -1;
    }
    
    // 向缓存中添加某个值
    void set(int key,int value){
        map<int,CacheNode*>::iterator ite = mp.find(key);
        if (ite != mp.end()){
            CacheNode* node =ite->second;
            remove(node);
            setHead(node);
        }else{
            CacheNode* node;
            node->key=key;
            node->value=value;
            if (mp.size() >= size){
                map<int, CacheNode *>::iterator iter = mp.find(tail -> key);
      	        remove(tail);
	            mp.erase(iter);
            }else{
                setHead(node);
                mp[key]=node;
            }
            
        }
    }

    // 从双向链表中删除某个节点
    void remove(CacheNode* node){
        if (node->pre !=NULL){
            node->pre->next=node->next;
        }else{
            head = node->next;
        }
        if (node->next != NULL){
            node->next->pre=node->pre;
        }else{
            tail=node->pre;
        }
    }

    // 添加到双向链表中某个节点
    void setHead(CacheNode* node){
        node->next=head;
        node->pre=NULL;
        if (head != NULL){
            head->pre = node;
        }
        head = node;
        if (tail == NULL){
            tail = node;
        }
    }
};
