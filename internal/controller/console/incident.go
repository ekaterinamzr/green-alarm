package console

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
)

type incidentController struct {
	uc Incident
	l  logger.Logger
}

func newIncidentController(uc Incident, l logger.Logger) *incidentController {
	return &incidentController{uc, l}
}

func (r *incidentController) create(author int) error {
	var input dto.CreateIncidentRequest

	fmt.Print("Creating incident\n")
	input.Name, _ = inputString("Name")
	input.Date, _ = inputDate("Date")
	input.Country, _ = inputString("Country")
	input.Latitude, _ = inputFloat("Lattitude")
	input.Longitude, _ = inputFloat("Longitude")
	input.Comment, _ = inputString("Comment")
	input.Type, _ = inputInt("Type")

	input.Author = author

	output, err := r.uc.Create(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - incident - create")
		return fmt.Errorf("could not create incident: %w", err)
	}

	fmt.Printf("Created incident with id = %d", output.Id)

	return nil
}

func (r *incidentController) getAll() error {
	output, err := r.uc.GetAll(context.Background())
	if err != nil {
		r.l.Error(err, "console - incident - getAll")
		return fmt.Errorf("could not get incidents: %w", err)
	}
	fmt.Print(output)

	return nil
}

func (r *incidentController) getByType() error {
	incidentType, _ := inputInt("Type")

	output, err := r.uc.GetByType(context.Background(), dto.GetIncidentsByTypeRequest{IncidentType: incidentType})
	if err != nil {
		r.l.Error(err, "console - incident - get by type")
		return err
	}
	fmt.Print(output)
	return nil
}

func (r *incidentController) getById() error {
	var input dto.GetIncidentByIdRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - incident - getById")
		return err
	}

	input.Id = id

	output, err := r.uc.GetById(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - incident - getById")
		return err
	}

	fmt.Print(output)
	return nil
}

func (r *incidentController) update() error {
	var input dto.UpdateIncidentRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - incident - getById")
		return err
	}

	input.Id = id

	input.Name, _ = inputString("Name")
	input.Date, _ = inputDate("Date")
	input.Country, _ = inputString("Country")
	input.Latitude, _ = inputFloat("Lattitude")
	input.Longitude, _ = inputFloat("Longitude")
	input.Comment, _ = inputString("Comment")
	input.Type, _ = inputInt("Type")
	input.Status, _ = inputInt("Status")

	err = r.uc.Update(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - incident - update")
		return err
	}

	fmt.Print("Success\n")
	return nil
}

func (r *incidentController) delete() error {
	var input dto.DeleteIncidentRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - incident - getById")
		return err
	}

	input.Id = id

	err = r.uc.Delete(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - incident - delete")
		return err
	}

	fmt.Print("Success\n")
	return nil
}
