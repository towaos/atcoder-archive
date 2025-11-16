#include <bits/stdc++.h>

using namespace std;

int main(){
  ios_base::sync_with_stdio(false);
  cin.tie(nullptr);

  int N;
  cin >> N;

  string S;
  string status = "logout";
  int count = 0;

  for(int i=0;i<N;i++){
    cin >> S;
    if(S == "login" || S == "logout"){
      status = S;
    }
    if(status == "logout"){
      if(S == "private"){
        count++;
      }
    }
  }
  cout << count << endl;
  return 0;
}