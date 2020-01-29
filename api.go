package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Register(c echo.Context) error {
	m := &MessageResponse{}
	u := &User{}
	err := c.Bind(u)
	if err != nil {
		log.Print("la estructura recibida no es valida: %v", err)
		m.Type = "error"
		m.Message = "la estructura no es valida"
		return c.JSON(http.StatusBadRequest, m)
	}

	addUser(u)
	m.Type = "ok"
	m.Message = "registrado correctamente"
	return c.JSON(http.StatusCreated, m)

}

func Login(c echo.Context) error {
	m := &MessageResponse{}
	u := &User{}
	err := c.Bind(u)
	if err != nil {
		log.Print("la estructura recibida no es valida: %v", err)
		m.Type = "error"
		m.Message = "la estructura no es valida"
		return c.JSON(http.StatusBadRequest, m)
	}

	if !login(u) {
		m.Type = "error"
		m.Message = "usuario o contraseña invalido"
		return c.JSON(http.StatusUnauthorized, m)

	}

	m.Type = "ok"
	m.Message = "Bienvenido"
	m.Data = token
	return c.JSON(http.StatusOK, m)

}
