//
// Created by zzx on 2024/7/7.
//

#ifndef CPP_169_MAJORITY_ELEMENT_H
#define CPP_169_MAJORITY_ELEMENT_H

#include <map>

using namespace std;

class Solution {
public:
    int majorityElement(vector<int> &nums) {
        map<int, int> m;
        for (auto &e: nums) {
            m[e]++;
        }
        int n = nums.size();
        for (auto &e: m) {
            if (e.second > n / 2) {
                return e.first;
            }
        }
        return 0;
    }

};

#endif //CPP_169_MAJORITY_ELEMENT_H
