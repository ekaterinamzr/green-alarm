package console

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
)

type typeController struct {
	uc IncidentType
	l  logger.Logger
}

func newTypeController(uc IncidentType, l logger.Logger) *typeController{
	return &typeController{uc, l}
}

func (r *typeController) create(author int) error {
	var input dto.CreateTypeRequest
	var err error

	fmt.Print("Creating type\n")
	input.Name, _ = inputString("Name")

	output, err := r.uc.Create(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - type - create")
		return fmt.Errorf("could not create type: %w", err)
	}

	fmt.Printf("Created type with id = %d", output.Id)

	return nil
}

func (r *typeController) getAll() error {
	output, err := r.uc.GetAll(context.Background())
	if err != nil {
		r.l.Error(err, "console - type - getAll")
		return fmt.Errorf("could not get types: %w", err)
	}
	fmt.Print(output)
	return nil
}

func (r *typeController) getById() error {
	var input dto.GetTypeByIdRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - type - getById")
		return err
	}

	input.Id = id

	output, err := r.uc.GetById(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - type - getById")
		return err
	}

	fmt.Print(output)
	return nil
}

func (r *typeController) update() error {
	var input dto.UpdateTypeRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - type - getById")
		return err
	}

	input.Id = id

	input.Name, _ = inputString("Name")

	err = r.uc.Update(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - type - update")
		return err
	}

	fmt.Print("Success\n")
	return nil
}

func (r *typeController) delete() error {
	var input dto.DeleteTypeRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - type - getById")
		return err
	}

	input.Id = id

	err = r.uc.Delete(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - type - delete")
		return err
	}

	fmt.Print("Success\n")
	return nil
}
