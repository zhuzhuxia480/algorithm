//
// Created by zzx on 2023/10/18.
//

#ifndef LEETCODE_104_MAXIMUM_DEPTH_OF_BINARY_TREE_H
#define LEETCODE_104_MAXIMUM_DEPTH_OF_BINARY_TREE_H

/**
 * Definition for a binary tree node.
 */
#include <algorithm>
using namespace std;
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;

    TreeNode() : val(0), left(nullptr), right(nullptr) {}

    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}

    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

TreeNode* makeTree(){
    TreeNode *root = new TreeNode(0);
    root->right = new TreeNode(2);
    return root;
}

class Solution {
public:
    int depth = 0;

    int getMaxDepth(TreeNode *root, int depth) {
        if (root == nullptr) {
            return depth;
        }
        int maxLeft = getMaxDepth(root->left, depth + 1);
        int maxRight = getMaxDepth(root->right, depth + 1);
        return max(maxLeft, maxRight);
    }

    int maxDepth(TreeNode *root) {
        return getMaxDepth(root, 0);
    }
};

#endif //LEETCODE_104_MAXIMUM_DEPTH_OF_BINARY_TREE_H
