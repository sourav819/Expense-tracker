package routers

import (
	"expense-tracker/pkg/config"
	"expense-tracker/pkg/logger"
	"fmt"
	"net/http"
)

func runServer(app config.AppConfig) {

	server := &http.Server{
		Addr:    ":" + app.Config.Server.Port,
		Handler: app.Router,
	}

	displayService()

	logger.Infof("Serving and listening on Port %s ", server.Addr)
	server.ListenAndServe()
}

// You can generate ASCI art here
// https://patorjk.com/software/taag/#p=display&f=Doom&t=sample
func displayService() {
	fmt.Println(`

	
$$$$$$$$\                                                                   $$$$$$$$\                           $$\                           
$$  _____|                                                                  \__$$  __|                          $$ |                          
$$ |      $$\   $$\  $$$$$$\   $$$$$$\  $$$$$$$\   $$$$$$$\  $$$$$$\           $$ | $$$$$$\  $$$$$$\   $$$$$$$\ $$ |  $$\  $$$$$$\   $$$$$$\  
$$$$$\    \$$\ $$  |$$  __$$\ $$  __$$\ $$  __$$\ $$  _____|$$  __$$\          $$ |$$  __$$\ \____$$\ $$  _____|$$ | $$  |$$  __$$\ $$  __$$\ 
$$  __|    \$$$$  / $$ /  $$ |$$$$$$$$ |$$ |  $$ |\$$$$$$\  $$$$$$$$ |         $$ |$$ |  \__|$$$$$$$ |$$ /      $$$$$$  / $$$$$$$$ |$$ |  \__|
$$ |       $$  $$<  $$ |  $$ |$$   ____|$$ |  $$ | \____$$\ $$   ____|         $$ |$$ |     $$  __$$ |$$ |      $$  _$$<  $$   ____|$$ |      
$$$$$$$$\ $$  /\$$\ $$$$$$$  |\$$$$$$$\ $$ |  $$ |$$$$$$$  |\$$$$$$$\          $$ |$$ |     \$$$$$$$ |\$$$$$$$\ $$ | \$$\ \$$$$$$$\ $$ |      
\________|\__/  \__|$$  ____/  \_______|\__|  \__|\_______/  \_______|         \__|\__|      \_______| \_______|\__|  \__| \_______|\__|      
                    $$ |                                                                                                                      
                    $$ |                                                                                                                      
                    \__|                                                                                                                      

	

   
	`)
}
