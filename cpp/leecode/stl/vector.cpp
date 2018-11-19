#include<vector>
#include<iostream>
#include<deque>
#include<list>
#include<stack>
#include<queue>
#include<set>
using namespace std;

// vector、deque和list这三者我们可以优先考虑vector。
// vector容器适用于大量读写，而插入、删除比较少的操作。
// list容器适用于少量读写，大量插入，删除的情况。
// deque折中了vector和deque， 如果你需要随机存取又关心数据的插入和删除，那么可以选择deque。

int main(){
    int a[8] = {1,6,9,7,3,2,0,6};
    vector<int>nums(a,a+8);
    // 1. 放入数据
    nums.push_back(12);
    // 2. 迭代器头和迭代器尾
    // vector<int>::iterator start = nums.begin();
    // vector<int>::iterator end = nums.end();
    /*
        *iter： 返回迭代器iter所指元素的引用
        iter->mem: 解引用iter并获取该元素的名为mem的成员，等价于(*item).mem
        ++iter: 另iter指向容器的下一个元素
        --iter: 另iter指向元素的前一个元素
        iter1 == iter2：判断两个迭代器是否相等
        iter1 != iter2: 判断两个迭代器是否不相等
    */
    auto start = nums.begin();
    auto end = nums.end();
    for (vector<int>::iterator iter=start; iter!=end; ++iter){
        cout<< *iter <<"\n";
    }
    auto mid = nums.begin() + nums.size()/ 2; 
    cout << "该容器的中间元素为:" << *mid << endl;

    /*--------------------- vector容器一些操作  --------------------*/
    nums.push_back(1);                  // push_back: 在末尾插入数据
    nums.pop_back();                    // pop_back: 去除末尾的元素
    nums.back();                        // back：取最后一个元素
    nums.front();                       // front：取第一个元素

    nums.size();                        // size: 已经保存的元素的数目
    nums.empty();                       // empty：判断容器是否为空
    nums.erase(nums.begin() + 1);       // erase：删除指定位置的元素
    nums.insert(nums.begin() + 1, 8);   // 在某个位置插入一个元素,效率低，不适合大批操作
    vector<int> vect2;
    vect2.assign(nums.begin(), nums.end()); // 赋值操作

    /*--------------------- deque容器一些操作  --------------------*/
    deque<int> d1;
    d1.push_back(1);         // 尾后压入元素
    d1.push_front(4);        // 队头压入元素
    d1.pop_back();           // 尾后弹出一个元素
    d1.pop_front();          // 队头弹出一个元素

    d1.front();              // 取队头元素
    d1.back();               // 取队尾元素

    /*--------------------- list容器一些操作  --------------------*/
    list<int> l;
    l.push_back(1);                      // 尾后压入元素
    l.push_front(4);                     // 队头压入元素
    l.pop_back();                        // 尾后弹出一个元素
    l.pop_front();                       // 队头弹出一个元素

    l.front();                           // 取队头元素
    l.back();                            // 取队尾元素

    l.insert(l.begin(), 88);             // 某个位置插入元素(性能好)
    l.remove(2);                         // 删除某个元素(和所给值相同的都删除)
    l.reverse();                         // 倒置所有元素
    l.erase(--l.end());                  // 删除某个位置的元素(性能好)

    /*--------------------- stack容器一些操作  --------------------*/
    stack<int> s;            // 定义栈
    s.push(2);		        //将item压入栈顶

    s.top();			    //返回栈顶的元素，但不会删除

    s.pop();			    //删除栈顶的元素，但不会返回
    s.size();			    //返回栈中元素的个数
    s.empty();			    //检查栈是否为空，如果为空返回true，否则返回false 
    
    /*--------------------- queue容器一些操作  --------------------*/
    queue<int> q;            // 定义队列
    q.push(1);           //将item压入队列尾部

    q.front();             //返回队首元素，但不删除
    q.back();               //返回队尾元素，但不删除

    q.pop();                //删除队首元素，但不返回
    q.size();               //返回队列中元素的个数
    q.empty();              //检查队列是否为空，如果为空返回true，否则返回false
    /*--------------------- queue容器一些操作  --------------------*/

    /*--------------------- set容器一些操作  --------------------*/
    set<int> ss;            // 定义队列
    ss.insert(1);           // 插入元素

    ss.find(1);

    ss.empty(); 　　　       //判断set容器是否为空
    ss.size(); 　　　　      //返回当前set容器中的元素个数
    ss.erase(1);            //删除资源
    
}