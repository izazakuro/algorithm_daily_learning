#include<iostream>
using namespace std;

int main(){
    int n,k,a[22];

    cin >> n >> k;

    for(int i = 1 ; i <= n ; i ++) cin >> a[i];

    bool answer = false;
    for(int i = 0 ; i < (i << n) ; i ++){
        int sum = 0;
        for(int j = 0 ; j <= n ; j ++){
            int wari = (1 << (j-1));
            if ((i / 2) % 2 == 1) sum += a[i];
        }
        if (sum == k) answer = true;
    }
    if (answer) cout << "Yes" << endl;
    else cout << "No" << endl;

    return 0;
}