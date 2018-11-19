// 字符串表示形式有两种
//1 char[10]:长度为10来表示字符串
//2 string 表示字符串时，结尾为'\0'
// 所以，char str[10]; strcpy(str,"0123456789");会有字符串越界的情况发生。

// https://blog.csdn.net/qian2213762498/article/details/79553372
#include<iostream>
#include<vector>
#include<string>
#include<algorithm>
using namespace std;

// 4、面试题4 替换空格
void ReplaceEmpty(string& str){
    cout<<str<<endl;
    // 替换某个字符
    replace(str.begin(), str.end(), ' ', '*');
    cout<<str<<endl;
}

int main(){
    char str1[] = "hello world";
    char str2[] = "hello world";
    if ( str1 == str2 ){
        cout<<"ok"<<"\n";
        cout<<str1<<"\n";
        cout<<str2<<"\n";
    }else{
        cout<<"!ok"<<endl;
        cout<< str1 <<endl;
        cout<< str2 <<endl;
    }
    string str ="he llo";
    ///转小写
    transform(str.begin(),str.end(),str.begin(),::tolower);
    cout<<"转化为小写后为："<<str<<endl;
    transform(str.begin(),str.end(),str.begin(),::toupper);
    cout<<"转化为大写后："<<str<<endl;
    ReplaceEmpty(str);
    return 0;
}

