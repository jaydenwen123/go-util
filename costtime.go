package util

import (
	"fmt"
	"time"
)

//Cost
type Cost struct {
	Start time.Time
	End  	time.Time
}

func NewCost(start time.Time) *Cost {
	return &Cost{Start: start}
}

func (c *Cost) SetEndTime(end time.Time)  {
		c.End=end
}


//设置终止时间
func (c *Cost) CalcCost() time.Duration {
		//return c.
	duration := c.End.Sub(c.Start)
	//return fmt.Sprintf("%+v",duration)
	return duration
}
//设置终止时间
func (c *Cost) ShowCalcCostAsString() string {
		//return c.
	return fmt.Sprintf("%+v", c.CalcCost())
}

//不用设置终止时间
func (c *Cost) CostWithNow() time.Duration {
	//return c.
	duration := time.Since(c.Start)
	//return fmt.Sprintf("%+v",duration)
	return duration
}

func (c *Cost) CostWithNowAsString() string {
		//return c.
	return fmt.Sprintf("%+v", c.CostWithNow())
}


