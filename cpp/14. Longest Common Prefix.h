//
// Created by zzx on 2023/11/19.
//

#ifndef CPP_14_LONGEST_COMMON_PREFIX_H
#define CPP_14_LONGEST_COMMON_PREFIX_H

#include <bits/stdc++.h>
class Solution {
public:

    bool checkIndexSame(int index, vector <string> &strs){
        int size = strs.size();
        if (strs[0].size() <= index) {
            return false;
        }
        char flag = strs[0][index];
        for (int i = 1; i < size; ++i) {
            if (strs[i].size() <= index) {
                return false;
            }
            if (flag != strs[i][index]) {
                return false;
            }
        }
        return true;
    }
    string longestCommonPrefix(vector<string>& strs) {

        int preFixLen = 0;
        int len = strs[0].size();
        for (int i = 0; i < len; ++i) {
            if (!checkIndexSame(preFixLen, strs)) {
                break;
            }
            preFixLen++;
        }

        return strs[0].substring(0, preFixLen);
    }
};

#endif //CPP_14_LONGEST_COMMON_PREFIX_H
