package handlers

import  "github.com/gofiber/fiber/v3"



func Welcome(c *fiber.Ctx)error{
   return c.Render("Welcome", nil, "layouts/main")
}