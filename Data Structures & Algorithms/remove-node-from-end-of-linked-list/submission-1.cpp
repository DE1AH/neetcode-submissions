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
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        if(!head){ return {}; }
        if(head && n == 1 && !head -> next){ return {}; }
        vector<ListNode*> nodes;
        while(head){
            nodes.push_back(head);
            head = head -> next;
        }
        nodes.erase(nodes.end() - n);
        ListNode* result = nodes[0];
        ListNode* temp = result;
        int i = 1;
        while(i < nodes.size()){
            temp -> next = nodes[i];
            temp = temp -> next;
            i++;
        }
        temp -> next = nullptr;
        return result;
    }
};
