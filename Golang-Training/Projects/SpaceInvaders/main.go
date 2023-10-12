package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bullet struct {
	posX     int32
	posY     int32
	velocity int32
	radius   float32
	Draw     bool
	Color    rl.Color
}
type Enemy struct {
}

func main() {
	// Window size
	screenWidth := int32(600)
	screenHeight := int32(900)

	// Variables
	var x_coords int32 = 2
	var y_coords int32 = 850

	// Bullets
	bullets := []Bullet{}
	var shouldShoot bool = true

	// Initialize window
	rl.InitWindow(screenWidth, screenHeight, "Space Invaders")
	ShipImage := rl.LoadImage("assets/ship.png")
	Ship := rl.LoadTextureFromImage(ShipImage)

	// End and unload window
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		//Draw ship and env
		rl.DrawTexture(Ship, x_coords, y_coords, rl.White)
		rl.DrawText("Score: ", 0, 0, 20, rl.White)

		//Bullet Mechanism
		for index1, current_bullet := range bullets {
			if current_bullet.Draw {
				bullets[index1].posY = bullets[index1].posY - current_bullet.velocity
				rl.DrawCircle(current_bullet.posX-16, current_bullet.posY, current_bullet.radius, current_bullet.Color)
				if current_bullet.posY < 0 || current_bullet.posY > screenHeight {
					bullets[index1].Draw = false

				}
			}
		}

		// Movement for ship
		if rl.IsKeyDown(rl.KeyRight) {
			if x_coords+1 >= screenWidth-20 {
			} else {
				x_coords += 5
			}
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			if x_coords-1 <= -15 {
			} else {
				x_coords -= 5
			}
		}
		if rl.IsKeyDown(rl.KeySpace) && shouldShoot {
			current_bullet := Bullet{int32(x_coords + 25), int32(y_coords + 25), 5, float32(5), true, rl.White}
			bullets = append(bullets, current_bullet)
			shouldShoot = false
		}
		rl.EndDrawing()
	}
	//End program
	rl.CloseWindow()
}
