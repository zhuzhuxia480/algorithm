//
// Created by zzx on 2023/11/6.
//

#ifndef CPP_367_VALID_PERFECT_SQUARE_H
#define CPP_367_VALID_PERFECT_SQUARE_H
class Solution {
public:
    bool isPerfectSquare(int num) {
        for (int i = 1; i <= num; ++i) {
            long long tmp = (long long)i * i;
            if (tmp == num) {
                return true;
            }
            if (tmp > num) {
                return false;
            }
        }
        return false;
    }
};
#endif //CPP_367_VALID_PERFECT_SQUARE_H
