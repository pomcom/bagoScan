model {
    views {
        systemContext {
            description "An example of a C4 model for a Go application."
            system("Go Application") {
                instances("main")
                instances("commands")
                instances("services")
                instances("core")
                instances("tools")
                instances("nmap")
                instances("testssl")
                instances("config")
            }
        }
    }

    containers {
        container("main") {
            component("main")
            component("monitoring")
            component("logger")
        }
        container("commands") {
            component("scan")
            component("startScan")
        }
        container("services") {
            component("TestRunnerService")
        }
        container("core") {
            component("TestRunner")
            component("Filehandler")
            component("Output")
        }
        container("tools") {
            component("Tool")
        }
        container("nmap") {
            component("Nmap")
        }
        container("testssl") {
            component("Testssl")
        }
        container("config") {
            component("Config")
            component("ConfigHandler")
        }
    }

    relationships {
        relationship("main", "commands", "uses")
        relationship("commands", "services", "uses")
        relationship("services", "core", "uses")
        relationship("core", "tools", "uses")
        relationship("services", "config", "uses")
        relationship("tools", "nmap", "uses")
        relationship("tools", "testssl", "uses")
    }
}

