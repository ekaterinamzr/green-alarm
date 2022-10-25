package console

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
)

type authController struct {
	uc Auth
	l  logger.Logger
}

func newAuthController(uc Auth, l logger.Logger) *authController{
	return &authController{uc, l}
}

func (r *authController) signUp() error {
	var input dto.SignUpRequest

	// input.FirstName = inputS

	fmt.Print("Registration\n")
	input.FirstName, _ = inputString("FirstName")
	input.LastName, _ = inputString("LastName")
	input.Username, _ = inputString("Username")
	input.Email, _ = inputString("Email")
	input.Password, _ = inputString("Password")

	output, err := r.uc.SignUp(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - user - create")
		return err
	}

	fmt.Printf("Created user with id = %d", output.Id)

	return nil
}

func (r *authController) signIn() (int, int, error) {
	var input dto.SignInRequest

	input.Username, _ = inputString("Username")
	input.Password, _ = inputString("Password")

	output, err := r.uc.SignIn(context.Background(), input)
	if err != nil {
		r.l.Error(err, "console - auth - signUp")
		return 0, 0, err
	}

	r.l.Info("Success")
	return output.Id, output.Role, nil
}
