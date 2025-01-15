package response

import (
	"errors"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/timmbarton/errors"
)

type Response struct {
	Result string      `json:"result"`
	Error  *errs.Err   `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func initEmptySlice(v interface{}) interface{} {
	val := reflect.ValueOf(v)

	// Проверяем, является ли v слайсом и nil
	if val.Kind() == reflect.Slice && val.IsNil() {
		// Создаем пустой слайс с capacity 0 для типа элемента
		return reflect.MakeSlice(val.Type(), 0, 0).Interface()
	}

	return v
}

func (r Response) WithData(data interface{}) Response {
	r.Data = initEmptySlice(data) // Если data - это слайс, то инициализируем его, чтоб возвращать не null, а []
	return r
}

const (
	resultOk    = "ok"
	resultError = "error"
)

var (
	RespOk = Response{Result: resultOk}
)

func WithError(c *fiber.Ctx, err *errs.Err) error {
	if err == nil {
		err = new(errs.Err)
		errors.As(errs.New(errs.ErrCodeInternal, 0, "unknown message"), &err)
	}

	return c.Status(int(err.Code)).JSON(Response{Result: resultError})
}

func OkWithData(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(RespOk.WithData(data))
}

func Ok(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(RespOk)
}
