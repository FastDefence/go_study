package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// JSON で現在時刻を返す API
	http.HandleFunc("/api/time", func(w http.ResponseWriter, r *http.Request) {
		jst := time.FixedZone("Asia/Tokyo", 9*60*60) // JST (UTC+9)
		now := time.Now().In(jst).Format(time.RFC3339)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"now":"%s"}`, now)
	})

	// ルート: 時刻表示ページを返す
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `
<!doctype html>
<html lang="ja">
<head>
<meta charset="utf-8">
<title>Current Time</title>
<style>
  body { font-family: sans-serif; text-align: center; margin-top: 5rem; }
  #clock { font-size: 3rem; }
</style>
</head>
<body>
  <h1>Current time</h1>
  <h3>JST </h3>
  <div id="clock">loading…</div>

  <script>
    async function updateClock() {
      try {
        const res  = await fetch('/api/time');
        const json = await res.json();
        document.getElementById('clock').textContent = json.now;
      } catch (e) {
        document.getElementById('clock').textContent = 'Error';
      }
    }
    updateClock();                   // 初回
    setInterval(updateClock, 1000);  // 毎秒
  </script>
</body>
</html>`)
	})

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
