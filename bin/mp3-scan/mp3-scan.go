package main

import (
	"fmt"
	"mp3s-reviewer/lib/mp3review"
	"mp3s-reviewer/lib/utils"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

func main() {
	utils.ConfigureDefaultZeroLogger()

	// --- config
	var targetDir string="C:/Users/ktkm2/Desktop/song jobs/2024-06-20"
	// var maybeMode bool=false
	// --- end config


	// --- auto vars
	var here string=utils.GetHereDirExe()


	// --- state initialise
	var targetFiles []string=getTargetFiles(targetDir)
	var currentFileIndex int=0


	log.Info().Msgf("tracking %d items",len(targetFiles))

	// --- fiber init
	var app *fiber.App = fiber.New(fiber.Config{
		CaseSensitive: true,
		ErrorHandler: func(c fiber.Ctx, err error) error {
			fmt.Println("fiber error")
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		},

	})




	// --- routes
	// get current status information
	app.Get("/get-status",func(c fiber.Ctx) error {
		var response Mp3ReviewStatus=createCurrentState(targetFiles,currentFileIndex)

		return c.JSON(response)
	})

	// open the current item, or nothing if ended
	app.Get("/open-item",func(c fiber.Ctx) error {
		if currentFileIndex>=len(targetFiles) {
			log.Warn().Msg("tried to open item when already ended")
			return c.SendStatus(fiber.StatusConflict)
		}

		var currentItem string=targetFiles[currentFileIndex]

		log.Info().Msgf("opening item: %s",currentItem)
		utils.OpenTargetWithDefaultProgram(currentItem)

		return c.SendStatus(fiber.StatusOK)
	})

	// change current item to the next item. returns new state
	app.Get("/next-item",func(c fiber.Ctx) error {
		if currentFileIndex+1>=len(targetFiles) {
			log.Warn().Msg("tried to go to next item, but would be invalid item")
			currentFileIndex=len(targetFiles)
			return c.SendStatus(fiber.StatusConflict)
		}

		currentFileIndex+=1

		var result Mp3ReviewStatus=createCurrentState(targetFiles,currentFileIndex)

		return c.JSON(result)
	})

	app.Post("/decide-item",func(c fiber.Ctx) error {
		if currentFileIndex>=len(targetFiles) {
			log.Warn().Msg("tried to perform decide item when out of items")
			return c.SendStatus(fiber.StatusConflict)
		}

		var decisionReq ItemDecisionRequest
		var e error=c.Bind().JSON(&decisionReq)

		if e!=nil {
			panic(e)
		}

		e=mp3review.DoItemDecision(
			targetFiles[currentFileIndex],
			decisionReq.Decision,
		)

		if e!=nil {
			panic(e)
		}

		// <advance state>
		// <construct and return current state>
	})



	// --- static
	app.Static("/",filepath.Join(here,"mp3-scan-v2-web/build"))


	// --- run
    var e error=utils.OpenTargetWithDefaultProgram("http://localhost:4200")

    if e!=nil {
        log.Err(e).Msg("failed to open webpage with default program")
    }

	app.Listen(":4200",fiber.ListenConfig{
		DisableStartupMessage: true,
	})
}




// status to give to frontend
type Mp3ReviewStatus struct {
	CurrentItem string `json:"currentItem"`
	CurrentItemFolder string `json:"currentItemFolder"`
	TotalItems int `json:"totalItems"`
	CurrentItemIndex int `json:"currentItemIndex"`
	NoMoreItems bool `json:"noMoreItems"`
}

// request from frontend to decide on current item
type ItemDecisionRequest struct {
	Decision mp3review.Mp3Decision `json:"decision"`
}

// uses find mp3s to find mp3s. shuffles and returns result
func getTargetFiles(targetDir string) []string {
	var foundFiles []string=mp3review.FindMp3s(targetDir)
	utils.ShuffleArray(foundFiles)

	return foundFiles
}

// create review status struct
func createCurrentState(files []string,currentIndex int) Mp3ReviewStatus {
	return Mp3ReviewStatus{
		CurrentItem: filepath.Base(files[currentIndex]),
		CurrentItemFolder: filepath.Base(filepath.Dir(files[currentIndex])),
		TotalItems: len(files),
		CurrentItemIndex: currentIndex,
	}
}