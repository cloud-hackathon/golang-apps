// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "chat/app/chatroom"
	controllers "chat/app/controllers"
	tests "chat/tests"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers1 "github.com/revel/modules/testrunner/app/controllers"
	websocket "golang.org/x/net/websocket"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.WebSocket)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Room",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					15: []string{ 
						"user",
					},
				},
			},
			&revel.MethodType{
				Name: "RoomSocket",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "ws", Type: reflect.TypeOf((**websocket.Conn)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Application)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					12: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "EnterDemo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "demo", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.LongPolling)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Room",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					15: []string{ 
						"user",
					},
				},
			},
			&revel.MethodType{
				Name: "Say",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "message", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "WaitMessages",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "lastReceived", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Leave",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Refresh)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Room",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					28: []string{ 
						"user",
						"events",
					},
				},
			},
			&revel.MethodType{
				Name: "Say",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "message", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Leave",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					72: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					125: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"chat/app/controllers.Application.EnterDemo": { 
			16: "user",
			17: "demo",
		},
	}
	testing.TestSuites = []interface{}{ 
		(*tests.ApplicationTest)(nil),
	}

	revel.Run(*port)
}
