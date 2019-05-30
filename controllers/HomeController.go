package controllers

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	c.setTpl()
}
