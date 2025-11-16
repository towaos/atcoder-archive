#include <bits/stdc++.h>

using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(nullptr);

  string S;
  cin >> S;
  string res = "";

  for(int i=0;i<S.size();i++){
    if(isupper(S[i])){
      res += S[i];
    }
  }

  cout << res << endl;
  return 0;

}