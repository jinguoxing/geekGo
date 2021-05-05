package Week03


import (
	"context"
	"geekGo/Week03/registry"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"github.com/google/uuid"
	"errors"

)



// App is an application components lifecycle manager
type App struct {

	opts options
	ctx context.Context
	cancel func()
	instance *registry.ServiceInstance

}

func New(opts ...Option) *App {

	options := options{
		ctx : context.Background(),
		sigs: []os.Signal{syscall.SIGTERM,syscall.SIGQUIT,syscall.SIGINT},
	}

	if id, err := uuid.NewUUID(); err == nil {
		options.id = id.String()
	}

	for _, o := range opts {
		o(&options)
	}
	ctx, cancel := context.WithCancel(options.ctx)
	return &App{
		opts:     options,
		ctx:      ctx,
		cancel:   cancel,
		instance: buildInstance(options),
	}
}


// Run executes all OnStart hooks registered with the application's Lifecycle.
func (a *App) Run() error {

	g, ctx := errgroup.WithContext(a.ctx)
	for _, srv := range a.opts.servers {
		srv := srv
		g.Go(func() error {
			<-ctx.Done() // wait for stop signal
			return srv.Stop()
		})
		g.Go(func() error {
			return srv.Start()
		})
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, a.opts.sigs...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				a.Stop()
			}
		}
	})
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// Stop gracefully stops the application.
func (a *App) Stop() error {

	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

func buildInstance(o options) *registry.ServiceInstance {
	if len(o.endpoints) == 0 {
		for _, srv := range o.servers {
			if e, err := srv.Endpoint(); err == nil {
				o.endpoints = append(o.endpoints, e)
			}
		}
	}
	return &registry.ServiceInstance{
		ID:        o.id,
		Name:      o.name,
		Version:   o.version,
		Metadata:  o.metadata,
		Endpoints: o.endpoints,
	}
}




