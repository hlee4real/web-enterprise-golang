package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PageFilter struct {
	Page int
	Size int
}

func (p *PageFilter) Check(c *gin.Context) error {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s", err))
	}
	p.Page = page
	if p.Page <= 0 {
		return fmt.Errorf("error: page must be greater than 0")
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "5"))
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s", err))
	}
	p.Size = size
	if p.Size <= 0 {
		return fmt.Errorf("error: size must be greater than 0")
	}

	return nil
}
