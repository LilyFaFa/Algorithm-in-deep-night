### 题目
给定一组数据，可以使用乘号和加号，给定一个目标值，添加符号使表达式的和为目标值
Input: num = "123", target = 6
Output: ["1+2+3", "1*2*3"]

Input: num = "105", target = 5
Output: ["1*0+5","10-5"]
 
### 使用搜索的方式
如果使用动态规划，广度优先搜索这个题目会爆掉
要求输出所有的解，所以需要使用深度优先搜索