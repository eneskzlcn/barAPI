package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)
const PONG = "pong"
type PingCount struct{
	Times int `json:"times"`
}
type PingPongs struct{
	Pongs []string `json:"pongs"`
}

func (p *PingPongs) Construct(times int){
	p.Clear()
	for i := 0; i < times; i++ {
		p.Pongs = append(p.Pongs, PONG)
	}
}
func (p *PingPongs) Clear(){
	p.Pongs =[]string{}
}
func StartServer(port int){
	app:= fiber.New()
	app.Use(cors.New())
	app.Post("/ping", func(ctx *fiber.Ctx) error {
		body := PingCount{}
		if err := ctx.BodyParser(&body); err != nil {
			return err
		}
		var pingPongs = PingPongs{}
		pingPongs.Construct(body.Times)
		return ctx.JSON(pingPongs)
	})
	err:= app.Listen(fmt.Sprintf(":%d",port))
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	StartServer(3000)
}
