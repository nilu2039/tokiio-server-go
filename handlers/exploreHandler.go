package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/nilu2039/tokiio-server-go/utils/requests"
)

type ErrorMessageStruct struct {
	message string
}

func TopAiring(c *fiber.Ctx) error {
	const URL = `https://api.consumet.org/meta/anilist/trending?page=1`

	res, err := requests.Get(URL)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "something went wrong"})
	}

	return c.JSON(json.RawMessage(res))

}
