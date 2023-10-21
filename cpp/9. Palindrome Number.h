//
// Created by zzx on 2023/10/16.
//

#ifndef LEETCODE_PALINDROME_NUMBER_H
#define LEETCODE_PALINDROME_NUMBER_H

#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    bool isPalindrome(int x) {
        string str = to_string(x);
        int len = str.size();
        for (int i = 0; i < len; ++i) {
            if (str[i] != str[len - i - 1]) {
                return false;
            }
        }
        return true;
    }
};

#endif //LEETCODE_PALINDROME_NUMBER_H
