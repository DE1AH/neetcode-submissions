class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map <int, int> mp {};
        for(int i = 0; i < nums.size(); i++){
            mp[nums[i]] ++;
        }
        int temp = k;
        vector<int> max;
        while( temp > 0 ){
            int max_num = -10000;
            int max_freq = -10000;
            for (auto& pair : mp){
                if (pair.second > max_freq){
                    max_freq = pair.second;
                    max_num = pair.first;
                }
            }
            max.push_back(max_num);
            mp.erase(max_num);  
            temp --;
        }
        return max;
    }
};
