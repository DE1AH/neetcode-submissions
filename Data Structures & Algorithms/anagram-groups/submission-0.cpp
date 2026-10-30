class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>> result;
        unordered_map<string, vector<string>> mp;
        for(string& str : strs){
            string tmp = str;
            sort(str.begin(), str.end());
            string sortedS = str;
            mp[sortedS].push_back(tmp);
        }
        for(auto& pair : mp){
            result.push_back(pair.second);
        }
        return result;
    }
};
