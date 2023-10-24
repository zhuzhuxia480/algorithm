//
// Created by zzx on 2023/10/24.
//

#ifndef CPP_2236_ROOT_EQUALS_SUM_OF_CHILDREN_H
#define CPP_2236_ROOT_EQUALS_SUM_OF_CHILDREN_H

/**
 * Definition for a binary tree node.
 */
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};
class Solution {
public:
    bool checkTree(TreeNode* root) {
        return root->val == root->left->val + root->right->val;
    }
};
#endif //CPP_2236_ROOT_EQUALS_SUM_OF_CHILDREN_H
