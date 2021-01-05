package main

import (
	"app/yut/manager/authManager"
	"app/yut/manager/configManager"
	"app/yut/manager/mysqlManager"
	"app/yut/manager/userManager"
	"app/yut/service/blogservice"
	"app/yut/service/userservice"
	"app/yut/utils"
	"app/yut/utils/fileutils"
	"context"
	"flag"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"
)


var g_signal = make(chan os.Signal, 1)
var logXml = flag.String("l4g", "../settings/log.xml", "")
var pidfile = flag.String("pidfile", "./yut.pid", "")

func main() {
	defer func() {
		l4g.Info("defer close")
		if err := recover(); err != nil {
			l4g.Close()
			utils.PanicExt("Bug %v", err)
		}
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	l4g.LoadConfiguration(*logXml)
	defer l4g.Close()

	// 权限路由加载
	if err := authManager.GetAuthManager().Load(); err != nil {
		utils.PanicExt(err.Error())
	}
	// 加载mysql
	serverConf, err := configManager.LoadServerConfig()
	if err != nil {
		utils.PanicExt(err.Error())
		return
	}

	mysqlProxy := mysqlManager.GetMysqlProxy()
	if err := mysqlProxy.Init(serverConf.GetMysqlAddr()); err != nil {
		utils.PanicExt(err.Error())
	}
	defer mysqlProxy.Close()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60*time.Second))

	r.Route("/api", func(r chi.Router) {

		r.Use(ApiMiddleware)

		r.Mount("/user", UserRouter())
		r.Mount("/global", GlobalRouter())
		r.Mount("/blog", BlogRouter())
	})

	r.Route("/pub", func(r chi.Router) {

		r.Route("/blog/{blog_id}", func(r chi.Router) {
			r.Use(PubBlogMiddleware)
			r.Post("/", blogservice.GetBlog)
		})

		r.Route("/user", func(r chi.Router) {
			r.Post("/profile", userservice.GetProfile)
			r.Post("/register", userservice.UserRegisterPup)
		})
	})

	httpServer := &http.Server{Addr: serverConf.Http, Handler: r}

	go httpServer.ListenAndServe()
	defer func() {
		httpServer.Shutdown(context.Background())
		fileutils.DeleteFile(*pidfile)
	}()


	l4g.Debug("api server start %s", serverConf.Http)

	p, _ := filepath.Abs(filepath.Dir("./public/"))
	m := http.NewServeMux()
	fs := http.FileServer(http.Dir(p))
	//m.Handle("/", http.StripPrefix("/", fs))
	m.Handle("/", StaticMiddleware("/", fs))

	staticServ := &http.Server{Addr: ":9000", Handler: m}
	l4g.Debug("static server start: %d", 9000)


	go staticServ.ListenAndServe()

	// 写进程id
	pid := fmt.Sprintf("%d", os.Getpid())
	if err := fileutils.FilePutContents(*pidfile, pid); err != nil {
		utils.PanicExt(err.Error())
	}

	listenSignal(context.Background(), httpServer, staticServ)
}

func listenSignal(ctx context.Context, httpSrv *http.Server, httpStaticSrv *http.Server) {
	signal.Notify(g_signal, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)

	select {
	case sig := <- g_signal:
		l4g.Warn("catch signal %s \n", sig.String())
		userManager.GetUsrSessionMgr().OnShutDown()
		httpSrv.Shutdown(ctx)
		httpStaticSrv.Shutdown(ctx)
		fileutils.DeleteFile(*pidfile)
	}
}
