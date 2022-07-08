package entities

import (
	"fmt"
	"strconv"
	"strings"
)

type Account struct {
	Id      string `json:"id"`
	Balance int    `json:"balance"`
}

// Currency represents every currency that respect rule:
// c.Drobnaya * 100 = 1 * c.Celaya
type Currency struct {
	Celaya   int
	Drobnaya int
}

func (c Currency) MarshalJSON() ([]byte, error) {
	var strCurrency strings.Builder

	strCurrency.WriteString(strconv.Itoa(c.Celaya))
	strCurrency.WriteByte('.')
	strCurrency.WriteString(strconv.Itoa(c.Drobnaya))

	return []byte(strCurrency.String()), nil
}

func (c *Currency) UnmarshalJSON(b []byte) error {
	str := string(b)
	nums := strings.Split(str, ".")

	Celaya, err := strconv.Atoi(nums[0])
	if err != nil {
		return err
	}
	Drobnaya, err := strconv.Atoi(nums[1])
	if err != nil {
		return err
	}

	c.Celaya = Celaya
	c.Drobnaya = Drobnaya

	return nil
}

func (c Currency) String() string {
	return fmt.Sprintf("%d.%d", c.Celaya, c.Drobnaya)
}

func (c Currency) Add(value Currency) Currency {
	var result Currency

	result.Celaya = c.Celaya + value.Celaya
	dr := c.Drobnaya + value.Drobnaya
	if dr >= 100 {
		result.Celaya += dr / 100
		result.Drobnaya = dr % 100
	} else {
		result.Drobnaya = dr
	}
}
