#include"bits/stdc++.h"
using namespace std;
#define int long long
#define rep(i,n) for(int i=0;i<n;i++)
const long long mod = 1000000007;
const long long inf = 1ll << 61;
typedef pair<int, int> P;
typedef pair<double,P> PP;
struct edge { int to; int cost; };
int s[15][15];
int k[15];
int p[15];
signed main() {
	// take input
	int n, m; cin >> n >> m;
		// rep(i, m) {
		// take input switches
	for(int i=0;i<m;i++) {
		cin >> k[i];
		rep(j, k[i]) {
			cin >> s[i][j];
			s[i][j]--;
		}
	}
	// rep(i, m)cin >> p[i];
	for(int i=0;i<m;i++){
		cin >> p[i];
	}

	// 
	int cnt = 0;
	// rep(i, 1 << n) {
	for(int i=0;i<1<<n;i++) {
		int sc = 0;
		// rep(j, m) {
		for(int j=0;j<m;j++) {
			int tc = 0;
			// rep(l, k[j]) {
			for(int l=0;l<k[j];l++) {
				if (i&(1 << s[j][l]))
					tc++;
			}
			if ((tc & 1) == (p[j] & 1))
				sc++;
		}
		if (sc == m)
			cnt++;
	}
	cout << cnt << endl;
}
