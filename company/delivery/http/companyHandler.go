package http

import (
	"backend/company"
	"backend/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CompanyHandler struct {
	companyUsecase company.Usecase
}

func NewCompanyHandler(router *mux.Router, cu company.Usecase) {
	handler := &CompanyHandler{
		companyUsecase: cu,
	}

	// mapping to the url
	router.HandleFunc("/companies", handler.GetAllCompanies).Methods("GET")
	router.HandleFunc("/companies/{id}", handler.GetCompanyById).Methods("GET")
	router.HandleFunc("/companies", handler.CreateCompany).Methods("POST")
	router.HandleFunc("/companies/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/companies/{id}", handler.DeleteCompany).Methods("DELETE")

}

// -------------------------------------------------------------------------

// Get all companies
func (h CompanyHandler) GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	companies, err := h.companyUsecase.GetAllCompanies()
	if err != nil {
		log.Panic(err)
		json.NewEncoder(w).Encode(err.Error())
	}
	if companies == nil {
		json.NewEncoder(w).Encode("No User Found!")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(companies)
}

// Get company by id
func (h CompanyHandler) GetCompanyById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// taking route values from the current request
	vars := mux.Vars(r)
	// taking uuid from the route path variable
	companyId, err := uuid.Parse(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Error parsing Id")
		return
	}
	user, err1 := h.companyUsecase.GetCompanyById(companyId)
	if err1 != nil {
		json.NewEncoder(w).Encode(err1.Error())
		return
	}
	json.NewEncoder(w).Encode(user)
}

// create company
func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	company := models.Company{}
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	inputErr := h.companyUsecase.SaveCompany(&company)
	if inputErr != nil {
		json.NewEncoder(w).Encode(inputErr.Error())
		return
	}
	json.NewEncoder(w).Encode(company)
}

// update company
func (h CompanyHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	vars := mux.Vars(r)
	companyId, err := uuid.Parse(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Error parsing Id")
		return
	}
	company := models.Company{}
	if err = json.NewDecoder(r.Body).Decode(&company); err != nil {
		json.NewEncoder(w).Encode("Error parsing the input data")
		return
	}
	updateErr := h.companyUsecase.UpdateCompany(companyId, &company)
	if updateErr != nil {
		json.NewEncoder(w).Encode(updateErr.Error())
		return
	}
	json.NewEncoder(w).Encode(company)
}

// delete company
func (h CompanyHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	vars := mux.Vars(r)
	companyId, err := uuid.Parse(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Error parsing Id")
		return
	}
	if err := h.companyUsecase.DeleteCompany(companyId); err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Delete Successful")
}
