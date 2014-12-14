package controllers

import (
	"fmt"

	"github.com/ggaaooppeenngg/OJ/app/models"
	"github.com/ggaaooppeenngg/OJ/app/routes"

	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

//URL: user/rating/:index
func (u *User) Rating(index int64) revel.Result {
	var users []models.User
	pagination := &Pagination{}
	pagination.isValidPage(u.Validation, models.User{}, index)
	if u.Validation.HasErrors() {
		u.FlashParams()
		u.Validation.Keep()
		return u.Redirect(routes.Notice.Crash())
	}
	err := engine.Limit(perPage, perPage*(pagination.current-1)).Desc("solved").Asc("id").Find(&users)
	if err != nil {
		fmt.Println(err)
	}
	err = pagination.Page(models.User{}, perPage, u.Request.Request.URL.Path, index)
	if err != nil {
		u.Flash.Error(err.Error())
		u.Redirect(routes.Notice.Crash())
	}
	return u.Render(users, pagination)
}

func (u *User) ProfileVisit(id int64) revel.Result {
	var user models.User
	has, err := engine.Id(id).Get(&user)
	if !has || err != nil {
		fmt.Println(has)
		u.Redirect(routes.Notice.Crash())
	}
	return u.Render(user)
}

func (u *User) Profile() revel.Result {
	var user models.User
	username := u.Session["username"]
	engine.Where("name = ?", username).Get(&user)
	return u.Render(user)
}

func (u *User) Solved() revel.Result {
	username := u.Session["username"]
	user := models.GetCurrentUser(username)
	if user != nil {
		usps, err := models.FindSovledProblems(user.Id)
		if err != nil {
			u.Flash.Error(err.Error())
			return u.Render()
		} else {
			return u.Render(usps)
		}
	} else {
		return u.Render()
	}

}
