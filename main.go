package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "<h1>Hello my</h1>")
	})
	mux.HandleFunc("/my2", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "<h1>Hello my2</h1>")
	})
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		html := `
<!DOCTYPE html>
<html>
<head>
    <style>
        .fancy-text {
            font-family: "Brush Script MT", "Comic Sans MS", cursive;
            font-size: 3em;
            text-align: center;
            margin-top: 20%;
            color: #2c3e50;
        }
    </style>
    <title>Home</title>
</head>
.<body>
    <div class="fancy-text">Welcome to Home Page</div>
</body>
</html>`
		_, _ = io.WriteString(w, html)
	})
	mux.HandleFunc("/snake", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		html := `
<!DOCTYPE html>
<html>
<head>
    <title>Snake Game</title>
    <style>
        body {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            background-color: #f0f0f0;
            font-family: Arial, sans-serif;
        }
        #game-container {
            text-align: center;
        }
        canvas {
            border: 2px solid #333;
            background-color: #fff;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
        }
        h1 { margin-bottom: 10px; color: #333; }
        .controls { margin-top: 10px; color: #666; font-size: 0.9em; }
    </style>
</head>
<body>
    <div id="game-container">
        <h1>Snake Game</h1>
        <canvas id="gameCanvas" width="400" height="400"></canvas>
        <div class="controls">Use Arrow Keys to Move</div>
        <div id="score">Score: 0</div>
    </div>

    <script>
        const canvas = document.getElementById('gameCanvas');
        const ctx = canvas.getContext('2d');
        const gridSize = 20;
        const tileCount = canvas.width / gridSize;
        
        let velocityX = 0;
        let velocityY = 0;
        let trail = [];
        let tail = 5;
        let playerX = 10;
        let playerY = 10;
        let appleX = 15;
        let appleY = 15;
        let score = 0;

        function game() {
            playerX += velocityX;
            playerY += velocityY;

            // 撞墙死亡机制
            if (playerX < 0 || playerX > tileCount - 1 || playerY < 0 || playerY > tileCount - 1) {
                tail = 5;
                score = 0;
                playerX = 10;
                playerY = 10;
                velocityX = 0;
                velocityY = 0;
                document.getElementById('score').innerText = 'Score: ' + score;
                return;
            }

            ctx.fillStyle = 'white';
            ctx.fillRect(0, 0, canvas.width, canvas.height);

            ctx.fillStyle = 'lime';
            for (let i = 0; i < trail.length; i++) {
                ctx.fillRect(trail[i].x * gridSize, trail[i].y * gridSize, gridSize - 2, gridSize - 2);
                if (trail[i].x === playerX && trail[i].y === playerY) {
                    tail = 5;
                    score = 0;
                    document.getElementById('score').innerText = 'Score: ' + score;
                }
            }

            trail.push({x: playerX, y: playerY});
            while (trail.length > tail) {
                trail.shift();
            }

            if (appleX === playerX && appleY === playerY) {
                tail++;
                score++;
                document.getElementById('score').innerText = 'Score: ' + score;
                appleX = Math.floor(Math.random() * tileCount);
                appleY = Math.floor(Math.random() * tileCount);
            }

            ctx.fillStyle = 'red';
            ctx.fillRect(appleX * gridSize, appleY * gridSize, gridSize - 2, gridSize - 2);
        }

        function keyPush(evt) {
            switch(evt.keyCode) {
                case 37: // Left
                    if (velocityX !== 1) { velocityX = -1; velocityY = 0; }
                    break;
                case 38: // Up
                    if (velocityY !== 1) { velocityX = 0; velocityY = -1; }
                    break;
                case 39: // Right
                    if (velocityX !== -1) { velocityX = 1; velocityY = 0; }
                    break;
                case 40: // Down
                    if (velocityY !== -1) { velocityX = 0; velocityY = 1; }
                    break;
            }
        }

        document.addEventListener('keydown', keyPush);
        // 速度调慢一倍 (15 -> 7.5)
        setInterval(game, 1000/7.5);
    </script>
</body>
</html>`
		_, _ = io.WriteString(w, html)
	})
	mux.HandleFunc("/my", func(w http.ResponseWriter, r *http.Request) {
		// 打开图片文件
		file, err := os.Open("my.jpg") // 替换为你的图片路径
		if err != nil {
			http.Error(w, "图片未找到", http.StatusNotFound)
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		// 设置 Content-Type（根据实际图片类型设置）
		w.Header().Set("Content-Type", "image/jpeg")

		// 将图片内容复制到响应
		_, err = io.Copy(w, file)
		if err != nil {
			// 如果已经写了部分响应，可能无法再返回错误码，
			// 但可以记录日志
			return
		}
	})
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}
}
