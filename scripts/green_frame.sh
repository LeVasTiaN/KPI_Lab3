#!/bin/bash

echo "🎨 Creating green frame with T-90 figures..."

curl -X POST http://localhost:17000/ -H "Content-Type: text/plain" -d "reset
white
bgrect 0.25 0.25 0.75 0.75
figure 0.5 0.5
green
figure 0.6 0.6
update"

if [ $? -eq 0 ]; then
    echo "✅ Green frame script executed successfully!"
else
    echo "❌ Failed to execute green frame script"
    exit 1
fi
