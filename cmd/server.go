package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/pkg/profile"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"untitled/internal/app"
)


// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	SetupAppWorkingDirectory()
}

func SetupAppWorkingDirectory()  {
	/*
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	*/
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		err := os.Setenv("APP_ROOT_PATH", path.Join(path.Dir(filename), ".."))
		if err != nil {
			panic(err)
		}
	} else {
		panic("wrong runtime variable")
	}

}
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	defer profile.Start().Stop()
	//
	router := app.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	log.Fatal(http.ListenAndServe(":8084", c.Handler(router)))
}
