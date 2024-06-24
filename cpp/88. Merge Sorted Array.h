//
// Created by zzx on 2024/6/24.
//

#ifndef CPP_88_MERGE_SORTED_ARRAY_H
#define CPP_88_MERGE_SORTED_ARRAY_H

#include <vector>

using namespace std;

class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {

        if (nums2.size() <= 0) {
            return;
        }
        vector<int> ret;
        int i = 0, j = 0;
        while (i < m && j < n) {
            if (nums1[i] < nums2[j]) {
                ret.push_back(nums1[i]);
                i++;
            } else {
                ret.push_back(nums2[j]);
                j++;
            }
        }

        while (j < n) {
            ret.push_back(nums2[j]);
            j++;
        }

        while (i < m) {
            ret.push_back(nums1[i]);
            i++;
        }
        nums1.assign(ret.begin(), ret.end());
    }
};

#endif //CPP_88_MERGE_SORTED_ARRAY_H
