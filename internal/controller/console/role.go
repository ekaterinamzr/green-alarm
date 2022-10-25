package console

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
)

type roleController struct {
	uc UserRole
	l  logger.Logger
}

func newRoleController(uc UserRole, l logger.Logger) *roleController{
	return &roleController{uc, l}
}

func (r *roleController) create(author int) error {
	var input dto.CreateRoleRequest
	var err error

	fmt.Print("Creating role\n")
	input.Name, _ = inputString("Name")

	output, err := r.uc.Create(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - role - create")
		return fmt.Errorf("could not create role: %w", err)
	}

	fmt.Printf("Created role with id = %d", output.Id)

	return nil
}

func (r *roleController) getAll() error {
	output, err := r.uc.GetAll(context.Background())
	if err != nil {
		r.l.Error(err, "console - role - getAll")
		return fmt.Errorf("could not get roles: %w", err)
	}
	fmt.Print(output)
	return nil
}

func (r *roleController) getById() error {
	var input dto.GetRoleByIdRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - role - getById")
		return err
	}

	input.Id = id

	output, err := r.uc.GetById(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - role - getById")
		return err
	}

	fmt.Print(output)
	return nil
}

func (r *roleController) update() error {
	var input dto.UpdateRoleRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - role - getById")
		return err
	}

	input.Id = id

	input.Name, _ = inputString("Name")

	err = r.uc.Update(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - role - update")
		return err
	}

	fmt.Print("Success\n")
	return nil
}

func (r *roleController) delete() error {
	var input dto.DeleteRoleRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - role - getById")
		return err
	}

	input.Id = id

	err = r.uc.Delete(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - role - delete")
		return err
	}

	fmt.Print("Success\n")
	return nil
}
