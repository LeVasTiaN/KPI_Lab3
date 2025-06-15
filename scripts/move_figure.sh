#!/bin/bash

echo "🎭 Starting T-90 figure animation..."

curl -X POST http://localhost:17000/ -H "Content-Type: text/plain" -d "reset
white
figure 0.2 0.2
update"

echo "✅ Initial T-90 figure placed!"

for i in {0..10}; do
    x=$(echo "scale=2; 0.2 + $i * 0.06" | bc)
    y=$(echo "scale=2; 0.2 + $i * 0.06" | bc)

    curl -X POST http://localhost:17000/ -H "Content-Type: text/plain" -d "move $x $y
update"

    echo "✅ Moved T-90 figure to ($x, $y)"
    sleep 1
done

echo "🎉 Animation complete!"
