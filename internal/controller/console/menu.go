package console

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
)

type Auth interface {
	SignUp(context.Context, dto.SignUpRequest) (*dto.SignUpResponse, error)
	SignIn(context.Context, dto.SignInRequest) (*dto.SignInResponse, error)

	ParseToken(context.Context, string) (id int, role int, err error)
}

type Incident interface {
	Create(context.Context, dto.CreateIncidentRequest) (*dto.CreateIncidentResponse, error)
	GetAll(context.Context) (*dto.GetIncidentsResponse, error)
	GetById(context.Context, dto.GetIncidentByIdRequest) (*dto.GetIncidentByIdResponse, error)
	Update(context.Context, dto.UpdateIncidentRequest) error
	Delete(context.Context, dto.DeleteIncidentRequest) error
	GetByType(context.Context, dto.GetIncidentsByTypeRequest) (*dto.GetIncidentsResponse, error)
}

type IncidentType interface {
	Create(context.Context, dto.CreateTypeRequest) (*dto.CreateTypeResponse, error)
	GetAll(context.Context) (*dto.GetAllTypesResponse, error)
	GetById(context.Context, dto.GetTypeByIdRequest) (*dto.GetTypeByIdResponse, error)
	Update(context.Context, dto.UpdateTypeRequest) error
	Delete(context.Context, dto.DeleteTypeRequest) error
}

type IncidentStatus interface {
	Create(context.Context, dto.CreateStatusRequest) (*dto.CreateStatusResponse, error)
	GetAll(context.Context) (*dto.GetAllStatusesResponse, error)
	GetById(context.Context, dto.GetStatusByIdRequest) (*dto.GetStatusByIdResponse, error)
	Update(context.Context, dto.UpdateStatusRequest) error
	Delete(context.Context, dto.DeleteStatusRequest) error
}

type User interface {
	GetAll(context.Context) (*dto.GetAllUsersResponse, error)
	GetById(context.Context, dto.GetUserByIdRequest) (*dto.GetUserByIdResponse, error)
	Update(context.Context, dto.UpdateUserRequest) error
	Delete(context.Context, dto.DeleteUserRequest) error
	ChangeRole(context.Context, dto.ChangeRoleRequest) error
}

type UserRole interface {
	Create(context.Context, dto.CreateRoleRequest) (*dto.CreateRoleResponse, error)
	GetAll(context.Context) (*dto.GetAllRolesResponse, error)
	GetById(context.Context, dto.GetRoleByIdRequest) (*dto.GetRoleByIdResponse, error)
	Update(context.Context, dto.UpdateRoleRequest) error
	Delete(context.Context, dto.DeleteRoleRequest) error
}

func Menu(l logger.Logger, a Auth, i Incident, t IncidentType, s IncidentStatus, u User, r UserRole) {
	ac := newAuthController(a, l)
	ic := newIncidentController(i, l)
	tc := newTypeController(t, l)
	sc := newStatusController(s, l)
	uc := newUserController(u, l)
	rc := newRoleController(r, l)

	curUser, curRole := 0, 0

	mGuest := `
	
	Menu:
	1) Sign up
	2) Sign in

	Incidents
	11) All
	12) By type
	13) By id
	
	0) Quit
	
	`

	mUser := `
	
	Menu:
	3) Sign out

	Incidents
	11) All
	12) By type
	13) By id
	14) Create
	
	0) Quit
	
	`

	mModer := `
	
	Menu:
	3) Sign out

	Incidents
	11) All
	12) By type
	13) By id
	13) Create
	14) Update
	15) Delete
	
	0) Quit
	
	`

	mAdmin := `
	
	Menu:
	3) Sign out

	Incidents
	11) All
	12) By type
	13) By id
	14) Create
	15) Update
	16) Delete

	Users
	21) All
	22) By id
	24) Update
	25) Delete

	Roles
	31) All
	32) By id
	33) Create
	34) Update
	35) Delete

	Types
	41) All
	42) By id
	43) Create
	44) Update
	45) Delete

	Statuses
	51) All
	52) By id
	53) Create
	54) Update
	55) Delete
	
	0) Quit
	
	`

	m := []string{mGuest, mAdmin, mModer, mUser}
	wait := ""
	task := 1
	for task != 0 {
		fmt.Print("\nPress enter to continue")
		fmt.Scanf("%s", &wait)
		print(m[curRole])

		task, _ = inputInt("Task")

		switch task {
		case 1:
			ac.signUp()
		case 2:
			curUser, curRole, _ = ac.signIn()
		case 3:
			curUser, curRole = 0, 0
		case 11:
			ic.getAll()
		case 12:
			ic.getByType()
		case 13:
			ic.getById()
		case 14:
			ic.create(curUser)
		case 15:
			ic.update()
		case 16:
			ic.delete()
		case 21:
			uc.getAll()
		case 22:
			uc.getById()
		case 23:
			uc.update()
		case 24:
			uc.delete()
		case 31:
			rc.getAll()
		case 32:
			rc.getById()
		case 33:
			rc.create(curUser)
		case 34:
			rc.update()
		case 35:
			rc.delete()
		case 41:
			tc.getAll()
		case 42:
			tc.getById()
		case 43:
			tc.create(curUser)
		case 44:
			tc.update()
		case 45:
			tc.delete()
		case 51:
			sc.getAll()
		case 52:
			sc.getById()
		case 53:
			sc.create(curUser)
		case 54:
			sc.update()
		case 55:
			sc.delete()
		}
	}
}
