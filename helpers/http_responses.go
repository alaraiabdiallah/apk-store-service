package helpers

import "github.com/labstack/echo"

func jsonMessage(status string, message interface{}) interface{}  {
	return echo.Map{
		"status": status,
		"message": message,
	}
}

func SuccessJsonMessage(message string)  interface{} {
	return jsonMessage("true", message)
}

func FailedJsonMessage(message interface{})  interface{} {
	return jsonMessage("false", message)
}