package main

import (
	"context"

	"github.com/cathalgarvey/fmtless"
	"github.com/checkinhq/checkin/web/apis/checkin/user/v1alpha"
	"github.com/checkinhq/checkin/web/pkg/config"
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

type Model struct {
	*js.Object

	authenticationService user.AuthenticationClient

	Form *Form `js:"form"`
	Rules map[string][]map[string]interface{} `js:"rules"`
}

type Form struct {
	*js.Object

	Email    string `js:"email"`
	Password string `js:"password"`
}

func (m *Model) Submit() {
	vm := vue.GetVM(m)

	vm.Get("$refs").Get("login").Call("validate", func(valid bool) bool {
		if !valid {
			vm.Call(
				"$alert",
				"Please enter valid login information!",
				map[string]string{
					"confirmButtonText": "OK",
				},
			)

			return false
		}

		request := &user.LoginRequest{
			Email:    m.Form.Email,
			Password: m.Form.Password,
		}

		go func() {
			response, err := m.authenticationService.Login(context.Background(), request)

			if err != nil {
				vm.Call(
					"$alert",
					fmt.Sprintf("login error: %s", err),
					map[string]string{
						"confirmButtonText": "OK",
					},
				)

				return
			}

			vm.Call(
				"$alert",
				fmt.Sprintf("login success: %s", response.GetToken()),
				map[string]string{
					"confirmButtonText": "OK",
				},
			)
		}()

		return valid
	})
}

func main() {
	c, err := config.Load()
	if err != nil {
		panic(err)
	}

	m := &Model{
		Object:                js.Global.Get("Object").New(),
		authenticationService: user.NewAuthenticationClient(c.GrpcHost),
	}

	// field assignment is required in this way to make data passing works
	m.Form = &Form{
		Object: js.Global.Get("Object").New(),
	}
	m.Form.Email = ""
	m.Form.Password = ""
	m.Rules = map[string][]map[string]interface{}{
		"email": {
			{
				"type":     "email",
				"required": true,
				"message":  "Please enter a valid email address!",
				"trigger":  "blur",
			},
		},
		"password": {
			{
				"required": true,
				"message":  "Please enter your password!",
				"trigger":  "blur",
			},
		},
	}

	// create the VueJS viewModel using a struct pointer
	vue.Use(js.Global.Get("ELEMENT"))
	vue.New("#app", m)
}
