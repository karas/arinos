package arinos

import (
	"fmt"
	"net/http"
	"net/http/fcgi"
	"net/url"
	"strings"
	"time"
)

type Arinos struct {
	LocalHost bool
	Options   *Options
	Router    *Router
}

type Router struct {
	Tree *node
}

type node struct {
	children     []*node
	component    string
	isNamedParam bool
	methods      map[string]Handle
}

type Options struct {
	Port int
}

type Option func(*Options)

type Handle func(http.ResponseWriter, *http.Request, url.Values)

func New(isLocalhost bool, options ...Option) (arinos *Arinos) {
	args := &Options{
		Port: 8000,
	}
	for _, option := range options {
		option(args)
	}
	router := NewRouter()
	return &Arinos{
		LocalHost: isLocalhost,
		Options:   args,
		Router:    router,
	}
}

func Port(portNumber int) Option {
	return func(args *Options) {
		args.Port = portNumber
	}
}

// Mock
func (arinos *Arinos) StartServe() error {
	if arinos.LocalHost {
		portString := fmt.Sprintf(":%d", arinos.Options.Port)
		server := &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			Addr:         portString,
			Handler:      arinos.Router,
		}
		return server.ListenAndServe()
	}
	return fcgi.Serve(nil, arinos.Router)
}

func NewRouter() *Router {
	node := node{
		component:    "/",
		isNamedParam: false,
		methods:      make(map[string]Handle),
	}
	return &Router{Tree: &node}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.Form
	node, _ := r.Tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	if handler := node.methods[req.Method]; handler != nil {
		handler(w, req, params)
	}
}

func (n *node) Add(method, path string, handler Handle) {
	if path == "" {
		panic("Path cannot be empty")
	}
	if path[0] != '/' {
		path = "/" + path
	}
	components := strings.Split(path, "/")[1:]
	for count := len(components); count > 0; count-- {
		aNode, component := n.traverse(components, nil)
		if aNode.component == component && count == 1 {
			aNode.methods[method] = handler
			return
		}
		newNode := node{
			component:    component,
			isNamedParam: false,
			methods:      make(map[string]Handle),
		}

		if len(component) > 0 && component[0] == ':' {
			newNode.isNamedParam = true
		}
		if count == 1 {
			newNode.methods[method] = handler
		}
		aNode.children = append(aNode.children, &newNode)
	}
}

func (n *node) traverse(components []string, params url.Values) (*node, string) {
	component := components[0]
	if len(n.children) > 0 {
		for _, child := range n.children {
			if component == child.component || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params.Add(child.component[1:], component)
				}
				next := components[1:]
				if len(next) > 0 {
					return child.traverse(next, params)
				} else {
					return child, component
				}
			}
		}
	}
	return n, component
}
