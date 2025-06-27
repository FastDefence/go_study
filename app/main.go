package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// JSON で現在時刻を返す API
	http.HandleFunc("/api/time", func(w http.ResponseWriter, r *http.Request) {
		locJST, _ := time.LoadLocation("Asia/Tokyo")
		locUTC, _ := time.LoadLocation("UTC")
		locPST, _ := time.LoadLocation("America/Los_Angeles")

		now := time.Now()
		response := fmt.Sprintf(`{
			"JST": "%s",
			"UTC": "%s",
			"PST": "%s"
		}`, now.In(locJST).Format(time.RFC3339),
			now.In(locUTC).Format(time.RFC3339),
			now.In(locPST).Format(time.RFC3339))

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, response)
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
  <div id="clock">
	<div>JST: <span id="jst">loading...</span></div>
	<div>UTC: <span id="utc">loading...</span></div>
	<div>PST: <span id="pst">loading...</span></div>
  </div>

  <script>
	const jstText = document.getElementById('jst');
	const utcText = document.getElementById('utc');
	const pstText = document.getElementById('pst');

	function toReadable(iso, timeZone) {
		const d = new Date(iso);
		return new Intl.DateTimeFormat('ja-JP', {
		timeZone,
		year: 'numeric',
		month: '2-digit',
		day: '2-digit',
		hour: '2-digit',
		minute: '2-digit',
		second: '2-digit',
		hour12: false
		}).format(d).replace(/\//g, '-');      // 2025/06/28 → 2025-06-28
	}

	async function updateClock() {
		try {
		const res  = await fetch('/api/time');
		const json = await res.json();

		jstText.textContent = toReadable(json.JST, 'Asia/Tokyo');
		utcText.textContent = toReadable(json.UTC, 'UTC');
		pstText.textContent = toReadable(json.PST, 'America/Los_Angeles');
		} catch {
		document.getElementById('clock').textContent = 'Error';
		}
	}

	updateClock();
	setInterval(updateClock, 1000);
  </script>

</body>
</html>`)
	})

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
