class Solution {
public:
    bool isPalindrome(string s) {
        string rev = "";
        for(char ch : s){
            if (isalnum(ch)){
                rev += tolower(ch);
            }
        }
        return rev == string(rev.rbegin(), rev.rend());
    }
};
