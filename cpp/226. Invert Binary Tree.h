//
// Created by zzx on 2023/11/20.
//

#ifndef CPP_226_INVERT_BINARY_TREE_H
#define CPP_226_INVERT_BINARY_TREE_H
#include <bits/stdc++.h>

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
    TreeNode* invertTree(TreeNode* root) {
        if (root == nullptr) {
            return root;
        }
        TreeNode *tmp = root->left;
        root->left = root->right;
        root->right = tmp;
        if (root->left) {
            invertTree(root->left);
        }
        if (root->right) {
            invertTree(root->right);
        }
        return root;
    }
};
#endif //CPP_226_INVERT_BINARY_TREE_H
