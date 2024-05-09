package controller

import (
	valo "github.com/HammiAhlan/Donation"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetDonasiData(c *fiber.Ctx) error{
	ps :=valo.GetAllDonasi()
	return c.JSON(ps)
}
