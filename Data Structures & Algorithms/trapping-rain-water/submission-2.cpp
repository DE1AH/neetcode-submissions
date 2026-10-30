class Solution {
public:
    int trap(vector<int>& h) {
        int V = 0;
        int l = 0;
        int r = h.size() - 1;
        int l_m = h[l];
        int r_m = h[r];
        while (l < r){
            if(l_m < r_m){
                l ++;
                l_m = max(l_m, h[l]);
                V += l_m - h[l];
            }
            else{
                r --;
                r_m = max(r_m, h[r]);
                V += r_m - h[r];
            }
        }
        return V;
    }
};
