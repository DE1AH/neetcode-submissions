/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */

class Solution {
private:
    int DFS(TreeNode* root, int& res){
        if(!root){ return 0; }
        int R = DFS(root->right, res);
        int L = DFS(root->left, res);
        res = max(res, R+L);
        return max(R,L) + 1;
    }

public:
    int diameterOfBinaryTree(TreeNode* root) {
        int res = 0;
        DFS(root, res);
        return res;
    }
};
