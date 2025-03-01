package validation

// func RegisterValidate(request model.RegisterRequest) {
// 	err := validation.ValidateStruct(&request,
// 		validation.Field(&request.Name, validation.Required.When(request.Name == "").Error("Name is Required")),
// 		validation.Field(&request.Username, validation.Required.When(request.Username == "").Error("Username is Required")),
// 		validation.Field(&request.Password, validation.Required.When(request.Password == "").Error("Password is Required")),
// 	)

// 	if err != nil {
// 		panic(exception.ValidationError{
// 			Message: err.Error(),
// 		})
// 	}
// }

// func LoginValidate(request model.LoginRequest) {
// 	err := validation.ValidateStruct(&request,
// 		validation.Field(&request.Username, validation.Required.When(request.Username == "").Error("Username is Required")),
// 		validation.Field(&request.Password, validation.Required.When(request.Password == "").Error("Password is Required")),
// 	)

// 	if err != nil {
// 		panic(exception.ValidationError{
// 			Message: err.Error(),
// 		})
// 	}
// }
