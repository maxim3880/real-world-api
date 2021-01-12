package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gofiber/fiber/v2"
)

type ctxKey int

const (
	UserDataKey ctxKey = iota
)

func OapiValidatorFromYamlFile(path string) (fiber.Handler, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %s", path, err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s as Swagger YAML: %s",
			path, err)
	}
	return OapiRequestValidator(swagger), nil
}

// Create a validator from a swagger object.
func OapiRequestValidator(swagger *openapi3.Swagger) fiber.Handler {
	return OapiRequestValidatorWithOptions(swagger, nil)
}

// Options to customize request validation. These are passed through to
// openapi3filter.
type Options struct {
	Options      openapi3filter.Options
	ParamDecoder openapi3filter.ContentParameterDecoder
	UserData     interface{}
	Skipper      *bool
}

// Create a validator from a swagger object, with validation options
func OapiRequestValidatorWithOptions(swagger *openapi3.Swagger, options *Options) fiber.Handler {
	router := openapi3filter.NewRouter().WithSwagger(swagger)
	skipper := getSkipperFromOptions(options)
	return func(ctx *fiber.Ctx) error {
		if skipper {
			return ctx.Next()
		}

		err := ValidateRequestFromContext(ctx, router, options)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			if errBuild := ctx.JSON(err); errBuild != nil {
				//nolint:errcheck
				ctx.Next()
				return errBuild
			}
			return nil
		}
		return ctx.Next()
	}
}

// This function is called from the middleware above and actually does the work
// of validating a request.
func ValidateRequestFromContext(ctx *fiber.Ctx, router *openapi3filter.Router, options *Options) error {
	rC := &http.Request{}
	rC.Clone(ctx.Context())
	rC.Body = ioutil.NopCloser(bytes.NewReader(ctx.Body()))
	rC.Header = http.Header{}
	ctx.Context().Request.Header.VisitAll(func(key, value []byte) {
		rC.Header.Add(string(key), string(value))
	})

	method := string(ctx.Context().Method())
	urlRoute, _ := url.Parse(string(ctx.Context().Path()))

	route, pathParams, err := router.FindRoute(method, urlRoute)

	// We failed to find a matching route for the request.
	if err != nil {
		switch e := err.(type) {
		case *openapi3filter.RouteError:
			// We've got a bad request, the path requested doesn't match
			// either server, or path, or something.

			return fiber.NewError(http.StatusBadRequest, e.Reason)
		default:
			// This should never happen today, but if our upstream code changes,
			// we don't want to crash the server, so handle the unexpected error.
			return fiber.NewError(http.StatusInternalServerError,
				fmt.Sprintf("error validating route: %s", err.Error()))
		}
	}

	queryArgs := url.Values{}
	ctx.Context().QueryArgs().VisitAll(func(key, value []byte) {
		queryArgs.Add(string(key), string(value))
	})

	validationInput := &openapi3filter.RequestValidationInput{
		Request:     rC,
		PathParams:  pathParams,
		Route:       route,
		QueryParams: queryArgs,
	}

	// Pass the Echo context into the request validator, so that any callbacks
	// which it invokes make it available.
	requestContext := ctx.Context()

	if options != nil {
		validationInput.Options = &options.Options
		validationInput.ParamDecoder = options.ParamDecoder
		requestContext.SetUserValue(fmt.Sprint(UserDataKey), options.UserData)
	}

	err = openapi3filter.ValidateRequest(requestContext, validationInput)
	if err != nil {
		switch e := err.(type) {
		case *openapi3filter.RequestError:
			// We've got a bad request
			// Split up the verbose error by lines and return the first one
			// openapi errors seem to be multi-line with a decent message on the first
			errorLines := strings.SplitN(e.Error(), "\n", 2)
			return fiber.NewError(http.StatusBadRequest, errorLines[0])
		case *openapi3filter.SecurityRequirementsError:
			for _, err := range e.Errors {
				httpErr, ok := err.(*fiber.Error)
				if ok {
					return httpErr
				}
			}
			return fiber.NewError(http.StatusBadRequest, fmt.Sprintf("%s %s", e.Error(), err))
		default:
			// This should never happen today, but if our upstream code changes,
			// we don't want to crash the server, so handle the unexpected error.
			return fiber.NewError(http.StatusInternalServerError, fmt.Sprintf("error validating request: %s", err))
		}
	}
	return nil
}

func GetUserData(c context.Context) interface{} {
	return c.Value(UserDataKey)
}

// attempt to get the skipper from the options whether it is set or not
func getSkipperFromOptions(options *Options) bool {
	if options == nil {
		return false
	}

	if options.Skipper == nil {
		return false
	}

	return *options.Skipper
}
