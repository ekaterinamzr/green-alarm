package console

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
)

type statusController struct {
	uc IncidentStatus
	l  logger.Logger
}

func newStatusController(uc IncidentStatus, l logger.Logger) *statusController{
	return &statusController{uc, l}
}

func (r *statusController) create(author int) error {
	var input dto.CreateStatusRequest
	var err error

	fmt.Print("Creating status\n")
	input.Name, _ = inputString("Name")

	output, err := r.uc.Create(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - status - create")
		return fmt.Errorf("could not create status: %w", err)
	}

	fmt.Printf("Created status with id = %d", output.Id)

	return nil
}

func (r *statusController) getAll() error {
	output, err := r.uc.GetAll(context.Background())
	if err != nil {
		r.l.Error(err, "console - status - getAll")
		return fmt.Errorf("could not get statuss: %w", err)
	}
	fmt.Print(output)
	return nil
}

func (r *statusController) getById() error {
	var input dto.GetStatusByIdRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - status - getById")
		return err
	}

	input.Id = id

	output, err := r.uc.GetById(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - status - getById")
		return err
	}

	fmt.Print(output)
	return nil
}

func (r *statusController) update() error {
	var input dto.UpdateStatusRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - status - getById")
		return err
	}

	input.Id = id

	input.Name, _ = inputString("Name")

	err = r.uc.Update(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - status - update")
		return err
	}

	fmt.Print("Success\n")
	return nil
}

func (r *statusController) delete() error {
	var input dto.DeleteStatusRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - status - getById")
		return err
	}

	input.Id = id

	err = r.uc.Delete(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - status - delete")
		return err
	}

	fmt.Print("Success\n")
	return nil
}
