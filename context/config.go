package context

type Config struct {
    Server struct {
        Hostname  string
        Host      string
        SockFile  string
        MaxConn   int
        Debug     bool
        PProfHost string
    }

    AlertMail struct {
        Username string
        Password string
        UseSsl   bool
        Subject  string
        Host     string
        Server   string
        Sender   string
        Receiver []string
    }

    Log struct {
        LogRotate string
        LogDir    string
    }

    Sqldb struct {
        Main struct {
            Host     string
            Port     int
            Username string
            Password string
            Dbname   string
            Charset  string
            Local    string
        }
    }

    Market struct {
        Okex struct {
            HttpHost  string
            ApiKey    string
            ApiSecret string
        }
        Bitmex struct {
            HttpHost  string
            WsHost    string
            Wss       string
            ApiKey    string
            ApiSecret string
        }
    }
}
