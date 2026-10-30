class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        unordered_map<int,int> mp{};
        for (int i = 0; i < numbers.size(); i++){
            mp[numbers[i]] = i;
        }
        for (int i = 0; i < numbers.size(); i ++){
            int temp = target - numbers[i];
            if (mp.count(temp) && mp[temp] != i){
                return {i + 1, mp[temp] + 1};
            }
        }
        return {};
    }
};
