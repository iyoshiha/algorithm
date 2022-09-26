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
	int n, m; cin >> n >> m;
	rep(i, m) {
		cin >> k[i];
		rep(j, k[i]) {
			cin >> s[i][j];
			s[i][j]--;
		}
	}
	rep(i, m)cin >> p[i];
	int cnt = 0;
	rep(i, 1 << n) {
		int sc = 0;
		rep(j, m) {
			int tc = 0;
			rep(l, k[j]) {
				if (i&(1 << s[j][l]))tc++;
			}
			if ((tc & 1) == (p[j] & 1))sc++;
		}
		if (sc == m)cnt++;
	}
	cout << cnt << endl;
}
