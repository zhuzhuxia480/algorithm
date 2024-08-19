//
// Created by zzx on 2024/8/19.
//

#ifndef CPP_344_REVERSE_STRING_H
#define CPP_344_REVERSE_STRING_H
#include <vector>
using namespace std;

class Solution {
public:
    void reverseString(vector<char>& s) {
        int length = s.size();
        for (int i = 0; i < length / 2; i++) {
            char tmp = s[i];
            s[i] = s[length - 1 - i];
            s[length - 1 - i] = tmp;
        }
    }
};

#endif //CPP_344_REVERSE_STRING_H
