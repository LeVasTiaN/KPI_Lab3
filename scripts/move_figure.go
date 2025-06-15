package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	initScript := `reset
white
figure 0.2 0.2
update`

	resp, err := http.Post("http://localhost:17000/", "text/plain", bytes.NewBufferString(initScript))
	if err != nil {
		log.Fatal("❌ Failed to send initial request:", err)
	}
	resp.Body.Close()
	log.Println("✅ Initial T-90 figure placed!")

	for i := 0; i <= 10; i++ {
		x := 0.2 + float64(i)*0.06
		y := 0.2 + float64(i)*0.06

		script := fmt.Sprintf(`move %.2f %.2f
update`, x, y)

		resp, err := http.Post("http://localhost:17000/", "text/plain", bytes.NewBufferString(script))
		if err != nil {
			log.Printf("❌ Failed to send move request: %v", err)
			continue
		}
		resp.Body.Close()
		log.Printf("✅ Moved T-90 figure to (%.2f, %.2f)", x, y)
		time.Sleep(1 * time.Second)
	}

	log.Println("🎉 Animation complete!")
}
