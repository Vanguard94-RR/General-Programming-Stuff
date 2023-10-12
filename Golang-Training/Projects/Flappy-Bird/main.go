package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Apple struct {
	posx   int32
	posy   int32
	width  int32
	height int32
	Color  rl.Color
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)
	rl.InitWindow(screenWidth, screenHeight, "FlappyApples")
	rl.InitAudioDevice()
	eat_noise := rl.LoadSound("sound/eat.wav")
	rl.SetTargetFPS(60)
	bird_down := rl.LoadImage("assets/bird-up.png")
	bird_up := rl.LoadImage("assets/bird-up.png")
	texture := rl.LoadTextureFromImage(bird_up)
	rand.Seed(time.Now().UnixNano())
	var apple_loc int = rand.Intn(450-2+1) - 22

	Apples := []Apple{}
	current_apple := Apple{screenWidth, int32(apple_loc), 15, 15, rl.Red}
	Apples = append(Apples, current_apple)

	var X_coord int32 = screenWidth/2 - texture.Width/2
	var Y_coord int32 = screenHeight/2 - texture.Height/2 - 40
	var score int = 0

	var velocity = 50000000

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.DrawTexture(texture, X_coord, Y_coord, rl.White)

		rl.DrawText("Current Score: "+strconv.Itoa(score), 0, 0, 30, rl.LightGray)

		rl.ClearBackground(rl.RayWhite)

		if !rl.IsKeyDown(rl.KeySpace) {
			texture = rl.LoadTextureFromImage(bird_up)
			Y_coord += 5
		} else {
			texture = rl.LoadTextureFromImage(bird_down)
			Y_coord -= 5
		}
		for io, current_apple := range Apples {
			rl.DrawRectangle(current_apple.posx, current_apple.posy, current_apple.width, current_apple.height, current_apple.Color)
			Apples[io].posx = Apples[io].posx - 3
			if current_apple.posx < 0 {
				Apples[io].posx = 800
				Apples[io].posy = int32(rand.Intn(450-2+1) - 2)
				score--
			}
			if rl.CheckCollisionRecs(rl.NewRectangle(float32(X_coord), float32(Y_coord), float32(44), float32(34)), rl.NewRectangle(float32(current_apple.posx), float32(current_apple.posy), float32(current_apple.width), float32(current_apple.height))) {
				Apples[io].posx = 800
				Apples[io].posy = int32(rand.Intn(450-2+1) - 2)
				score++
				rl.PlaySound(eat_noise)
			}
		}
		if Y_coord > 450 {
			rl.UnloadTexture(texture)
			Apples = nil
			rl.DrawText("Your final score is : "+strconv.Itoa(score), 30, 40, 35, rl.Red)

		}
		time.Sleep(time.Duration(velocity))
		rl.EndDrawing()
	}
	rl.StopSound(eat_noise)
	rl.UnloadSound(eat_noise)
	rl.UnloadTexture(texture)
	rl.CloseWindow()
}
