#include<iostream>
#include<vector>
#include<string>
using namespace std;

// sort比较函数
struct Student{
    int age;
    string name;
};

bool Less(Student s1, Student s2){
    // 年龄从小到大排序，第一个小返回true；
   return s1.age < s2.age;
}

int main(){
    Student s1 = {12, "lily"};
    Student s2 = {10, "tom"};
    vector<Student>ss;
    ss.push_back(s1);
    ss.push_back(s2);
    sort(ss.begin(),ss.end(),Less);
    cout<<ss[0].name<<endl;
}