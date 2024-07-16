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
	var state mp3review.Mp3ScanState=mp3review.NewScanState(targetDir)



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
		var response mp3review.Mp3ReviewStatus=state.GetStatus()

		return c.JSON(response)
	})

	// open the current item, or nothing if ended
	app.Get("/open-item",func(c fiber.Ctx) error {
		if state.NoMoreItems() {
			log.Warn().Msg("tried to open item when already ended")
			return c.SendStatus(fiber.StatusConflict)
		}

		state.OpenItem()

		return c.SendStatus(fiber.StatusOK)
	})

	// change current item to the next item. returns new state
	app.Get("/next-item",func(c fiber.Ctx) error {
		if state.NoMoreItems() {
			log.Warn().Msg("tried to go to next item, but would be invalid item")
			return c.SendStatus(fiber.StatusConflict)
		}

		var newStatus mp3review.Mp3ReviewStatus=state.AdvanceItem()

		return c.JSON(newStatus)
	})

	// make decision on current item, and move to next item
	app.Post("/decide-item",func(c fiber.Ctx) error {
		if state.NoMoreItems() {
			log.Warn().Msg("tried to perform decide item when out of items")
			return c.SendStatus(fiber.StatusConflict)
		}

		var decisionReq ItemDecisionRequest
		var e error=c.Bind().JSON(&decisionReq)

		if e!=nil {
			panic(e)
		}

		var newstatus mp3review.Mp3ReviewStatus=state.DecideItem(decisionReq.Decision)

		return c.JSON(newstatus)
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


// request from frontend to decide on current item
type ItemDecisionRequest struct {
	Decision mp3review.Mp3Decision `json:"decision"`
}