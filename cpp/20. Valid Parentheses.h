//
// Created by zzx on 2023/11/19.
//

#ifndef CPP_20_VALID_PARENTHESES_H
#define CPP_20_VALID_PARENTHESES_H
#include <bits/stdc++.h>


class Solution {
public:
    bool isValid(string s) {
        stack<char> tmpStack;
        tmpStack.push(s[0]);
        int len = s.size();
        for (int i = 1; i < len; ++i) {
            if (s[i] == '(' ||
                s[i] == '[' ||
                s[i] == '{') {
                tmpStack.push(s[i]);
            } else {
                if (tmpStack.size() == 0) {
                    return false;
                }
                if (s[i] == ')' && tmpStack.top() == '(') {
                    tmpStack.pop();
                    continue;
                }
                if (s[i] == ']' && tmpStack.top() == '[') {
                    tmpStack.pop();
                    continue;
                }
                if (s[i] == '}' && tmpStack.top() == '{') {
                    tmpStack.pop();
                    continue;
                }
                return false;
            }
        }
        if (tmpStack.size()) {
            return false;
        }
        return true;
    }
};
#endif //CPP_20_VALID_PARENTHESES_H
