class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        if(nums.size() == 0){ return 0; }
        if(nums.size() == 1){ return 1; }
        set<int> temp;
        for (int i = 0; i < nums.size(); i ++){
            temp.insert(nums[i]);
        }
        int cnt = 0;
        int tmp = 0;
        while(!temp.empty()){
            auto it = temp.begin();
            if(temp.count(*it + 1)){
                cnt ++;
            }
            else{
                cnt = 0;
            }
            tmp = max(cnt, tmp);
            temp.erase(it);
        }
        return tmp + 1;
    }
};
