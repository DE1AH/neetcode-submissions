/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */

class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        if(!head){ return {}; }
        vector<ListNode*> nodes;
        while(head){
            nodes.push_back(head);
            head = head -> next;
        }
        reverse(nodes.begin(), nodes.end());
        ListNode* tmp = nodes[0];
        ListNode* result = tmp;
        int i = 1;
        while(i < nodes.size()){
            tmp -> next = nodes[i];
            tmp = tmp -> next;
            i ++;
        }
        tmp -> next = nullptr;
        return result;
    }
};
