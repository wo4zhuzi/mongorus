## Use mongorus

### Direct connection

```
package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wo4zhuzi/mongorus"
)

var Logger *logrus.Logger

func main() {
	Logger = logrus.New()

	hooker, err := mongorus.NewMongoHook("127.0.0.1", "test_db", "test_collection")

	if err == nil {
		Logger.Hooks.Add(hooker)
	} else {
		fmt.Print(err)
	}

	Logger.WithFields(logrus.Fields{
		"blockHeight": 1000,
		"txid":        "0x......",
	}).Warn("warn message : .........")
}

```

### Authority connection

```
package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wo4zhuzi/mongorus"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Logger *logrus.Logger

func main() {
	Logger = logrus.New()
	hooker, err := mongorus.NewAuthMongoHook("127.0.0.1:12017", "test_db", "test_collection", options.Credential{
		Username: "test_username",
		Password: "test_password",
	})

	if err == nil {
		Logger.Hooks.Add(hooker)
	} else {
		fmt.Print(err)
	}

	Logger.WithFields(logrus.Fields{
		"blockHeight": 1000,
		"txid":        "0x......",
	}).Warn("warn message : .........")
}

```