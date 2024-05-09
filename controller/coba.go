package controller

import (
	hahay "github.com/HammiAhlan/Web_Donasi"
	
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := Web_Donasi.GetAllPresensi()
	return c.JSON(ps)
}

