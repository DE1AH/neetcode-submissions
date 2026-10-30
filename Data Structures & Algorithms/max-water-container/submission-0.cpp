class Solution {
public:
    int maxArea(vector<int>& heights) {
        int V = 0;
        int V_max = -1000;
        int i = 0;
        int j = heights.size() - 1;
        while (i != j){
            if(heights[i] > heights[j]){
                V_max = min(heights[i], heights[j]) * (j - i);
                j --;
            }
            else{
                V_max = min(heights[i], heights[j]) * (j - i);
                i ++;
            }
            V = max(V, V_max);
        }
        return V;
    }
};
