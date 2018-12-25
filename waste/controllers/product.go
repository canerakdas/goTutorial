package controllers

import (
	"net/http"
	"waste/models"
)

type Products struct {
	Product []models.Product
}

// Controller, need to model struct
func (p *Products) Get(w http.ResponseWriter, r *http.Request){
}

func (p *Products) GetWithId(w http.ResponseWriter, r *http.Request){
}

func (p *Products) Post(w http.ResponseWriter, r *http.Request){
}

func (p *Products) PostWithId(w http.ResponseWriter, r *http.Request){
}

func (p *Products) Put(w http.ResponseWriter, r *http.Request){
}

func (p *Products) Patch(w http.ResponseWriter, r *http.Request){
}

func (p *Products) Delete(w http.ResponseWriter, r *http.Request){
}

var Product Products