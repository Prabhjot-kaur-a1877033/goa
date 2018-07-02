// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// cars service
//
// Command:
// $ goa gen goa.design/goa/examples/streaming/design -o
// $(GOPATH)/src/goa.design/goa/examples/streaming

package carssvc

import (
	"context"

	carssvcviews "goa.design/goa/examples/streaming/gen/cars/views"
)

// The cars service lists car models by body style.
type Service interface {
	// Creates a valid JWT
	Login(context.Context, *LoginPayload) (res string, err error)
	// Lists car models by body style.
	List(context.Context, *ListPayload, ListServerStream) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "cars"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"login", "list"}

// ListServerStream is the interface a "list" endpoint server stream must
// satisfy.
type ListServerStream interface {
	// Send streams instances of "Car".
	Send(*Car) error
	// Close closes the stream.
	Close() error
	// SetView sets the view used to render the result before streaming.
	SetView(view string)
}

// ListClientStream is the interface a "list" endpoint client stream must
// satisfy.
type ListClientStream interface {
	// Recv reads instances of "Car" from the stream.
	Recv() (*Car, error)
}

// Credentials used to authenticate to retrieve JWT token
type LoginPayload struct {
	User     string
	Password string
}

// ListPayload is the payload type of the cars service list method.
type ListPayload struct {
	// The car body style.
	Style string
	// JWT used for authentication
	Token string
}

// Car is the result type of the cars service list method.
type Car struct {
	// The make of the car
	Make string
	// The car model
	Model string
	// The car body style
	BodyStyle string
}

// Credentials are invalid
type Unauthorized string

type InvalidScopes string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Credentials are invalid"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// Error returns an error description.
func (e InvalidScopes) Error() string {
	return ""
}

// ErrorName returns "invalid-scopes".
func (e InvalidScopes) ErrorName() string {
	return "invalid-scopes"
}

// NewCar initializes result type Car from viewed result type Car.
func NewCar(vres *carssvcviews.Car) *Car {
	var res *Car
	switch vres.View {
	case "default", "":
		res = newCar(vres.Projected)
	}
	return res
}

// NewViewedCar initializes viewed result type Car from result type Car using
// the given view.
func NewViewedCar(res *Car, view string) *carssvcviews.Car {
	var vres *carssvcviews.Car
	switch view {
	case "default", "":
		p := newCarView(res)
		vres = &carssvcviews.Car{p, "default"}
	}
	return vres
}

// newCar converts projected type Car to service type Car.
func newCar(vres *carssvcviews.CarView) *Car {
	res := &Car{}
	if vres.Make != nil {
		res.Make = *vres.Make
	}
	if vres.Model != nil {
		res.Model = *vres.Model
	}
	if vres.BodyStyle != nil {
		res.BodyStyle = *vres.BodyStyle
	}
	return res
}

// newCarView projects result type Car into projected type CarView using the
// "default" view.
func newCarView(res *Car) *carssvcviews.CarView {
	vres := &carssvcviews.CarView{
		Make:      &res.Make,
		Model:     &res.Model,
		BodyStyle: &res.BodyStyle,
	}
	return vres
}
