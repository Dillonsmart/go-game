package main

import (
	"encoding/json"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"go-win/utils"
)

var windowDimensions = [2]int32{1200, 750}
var ticks int32 = 0
var wave int32 = 1
var currentWaveEnemies []EnemyCounts
var waves []WaveDetails

type EnemyData struct {
	Name   string
	Colour string
	Health int32
}

type WaveFile struct {
	Waves []WaveDetails `json:"waves"`
}

type WaveDetails struct {
	Enemies []EnemyCounts `json:"enemies"`
}

type EnemyCounts struct {
	Type  int32 `json:"type"`
	Count int32 `json:"count"`
}

type ActiveEnemies struct {
	Enemies []EnemyData `json:"enemies"`
}

func main() {
	rl.InitWindow(windowDimensions[0], windowDimensions[1], "GoWin")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	waves = parseWaveData()

	// This is the infinite loop, which will run until the window is closed at 60 FPS
	for !rl.WindowShouldClose() {

		if len(currentWaveEnemies) == 0 {
			// no enemies have been spawned yet
			// spawn the enemies for the current wave
		}

		//ticks++
		//rl.BeginDrawing()
		//rl.ClearBackground(color.RGBA{R: 255, G: 255, B: 255, A: 255})
		//
		//rl.EndDrawing()
	}
}

func parseWaveData() []WaveDetails {
	data, err := utils.OpenFile("storage/data/waves.json")

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	var waveFile WaveFile
	// just ignore any error for now
	_ = json.Unmarshal(data, &waveFile)

	fmt.Printf("Waves: %+v\n", waveFile.Waves)

	return waveFile.Waves
}
