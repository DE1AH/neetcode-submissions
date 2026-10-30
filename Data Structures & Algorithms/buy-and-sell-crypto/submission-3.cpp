class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int max_profit = 0;
        int i = 0;
        int j = 1;
        while (j < prices.size()){
            if (prices[i] < prices[j]){
                max_profit = max(max_profit, prices[j] - prices[i]);
            }
            else{
                i = j;
            }
            j ++;
        }
        return max_profit;
    }
};
