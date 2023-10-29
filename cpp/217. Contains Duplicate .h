//
// Created by zzx on 2023/10/29.
//

#ifndef CPP_217_CONTAINS_DUPLICATE_H
#define CPP_217_CONTAINS_DUPLICATE_H
class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        int len = nums.size();
        map<int, bool> flag;
        for (int i = 0; i < len; ++i) {
            if (flag[nums[i]]) {
                return true;
            }
            flag[nums[i]] = true;
        }
        return false;
    }
};
#endif //CPP_217_CONTAINS_DUPLICATE_H
