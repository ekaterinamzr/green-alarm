package console

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
)

type userController struct {
	uc User
	l  logger.Logger
}

func newUserController(uc User, l logger.Logger) *userController {
	return &userController{uc, l}
}

func (r *userController) getAll() error {
	output, err := r.uc.GetAll(context.Background())
	if err != nil {
		r.l.Error(err, "console - user - getAll")
		return fmt.Errorf("could not get users: %w", err)
	}
	fmt.Print(output)
	return nil
}

func (r *userController) getById() error {
	var input dto.GetUserByIdRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - user - getById")
		return err
	}

	input.Id = id

	output, err := r.uc.GetById(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - user - getById")
		return err
	}

	fmt.Print(output)
	return nil
}

func (r *userController) update() error {
	var input dto.UpdateUserRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - user - getById")
		return err
	}

	input.Id = id

	input.Role, _ = inputInt("Role")
	input.FirstName, _ = inputString("FirstName")
	input.LastName, _ = inputString("LastName")
	input.Username, _ = inputString("Username")
	input.Email, _ = inputString("Email")
	input.Password, _ = inputString("Password")

	err = r.uc.Update(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - user - update")
		return err
	}

	fmt.Print("Success\n")
	return nil
}

func (r *userController) delete() error {
	var input dto.DeleteUserRequest

	id, err := inputInt("Id")
	if err != nil {
		r.l.Error(err, "console - user - getById")
		return err
	}

	input.Id = id

	err = r.uc.Delete(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - user - delete")
		return err
	}

	fmt.Print("Success\n")
	return nil
}
