package service

import (
	"banking-auth/domain"
	"banking-auth/dto"
	errs "banking-auth/error"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	CreateUser(dto.RegisterRequest) (*dto.RegisterResponse, *errs.AppError)
	LoginUser(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
	Verify(map[string]string) (bool, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.AuthRepository
	rolePermissions domain.RolePermissions
}

func (us DefaultAuthService) CreateUser(req dto.RegisterRequest) (*dto.RegisterResponse, *errs.AppError) {
	req.Validate()

	//transform the NewUserRequest to User domain to save in db
	domainUser := domain.User{
		Id:         "",
		Username:   req.Username,
		Password:   req.Password,
		CustomerId: req.CustomerId,
		Role: "USER",
	}
	//call the repository method to save the user to the database
	user, err := us.repo.CreateUser(domainUser)
	if err != nil {
		return nil, err
	}

	token, tokenErr := user.GenerateToken()
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &dto.RegisterResponse{
		Token: token,
	}, nil
}

func (us DefaultAuthService) LoginUser(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppError) {
	req.Validate()

	//call the repository method to find the user by username
	user, err := us.repo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	//compare the password
	if user.Password != req.Password {
		return nil, errs.NewValidationError("invalid password")
	}

	token, tokenErr := user.GenerateToken()
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}

/*
This method needs to verify both the validity of the token and the permissions
of the user to access the route and perform the action
*/
func (us DefaultAuthService) Verify(urlParams map[string]string) (bool, *errs.AppError) {
	//convert the jwt string to a JWT struct
	if jwtStruct, err := jwtTokenFromString(urlParams["token"]); err != nil {
		return false, errs.NewBadRequestError("invalid token")
	} else {
		if jwtStruct.Valid {
			mapClaims := jwtStruct.Claims.(jwt.MapClaims)
			//converting the JWT struct to a Claims struct
			if claims, err := domain.BuildClaimsFromJwtMapClaims(mapClaims); err != nil {
				return false, errs.NewBadRequestError("invalid token")
			} else {
				//check if the user has the right permissions to access the route
				if claims.IsUserRole() {
					if !claims.IsRequestVerifiedWithTokenClaims(urlParams) {
						return false, nil
					}
				}

				isAuthorizedRole := us.rolePermissions.IsAuthorizedFor(claims.Role, urlParams["routeName"])
				return isAuthorizedRole, nil
			}
		} else {
			return false, nil
		}
	}
}

func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func NewUserService(repo domain.AuthRepository) DefaultAuthService {
	return DefaultAuthService{
		repo: repo,
		rolePermissions: domain.GetRolePermissions(),
	}
}
