package main

import (
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/kataras/iris/v12"
)

var (
	// The Kafka brokers to connect to, as a comma separated list.
	brokers = []string{getenv("KAFKA_1", "localhost:9092")}
	// The config which makes our live easier when passing around, it pre-mades a lot of things for us.
	config *sarama.Config
)

func getenv(key string, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return def
}

func init() {
	config = sarama.NewConfig()
	config.ClientID = "iris-example-client"
	config.Version = sarama.V0_11_0_2
	// config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message.
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond
	config.Producer.Retry.Max = 10 // Retry up to 10 times to produce the message.
	config.Producer.Return.Successes = true

	// for SASL/basic plain text authentication: config.Net.SASL.
	// config.Net.SASL.Enable = true
	// config.Net.SASL.Handshake = false
	// config.Net.SASL.User = "myuser"
	// config.Net.SASL.Password = "mypass"

	config.Consumer.Return.Errors = true
}

func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, handleNotFound)

	app.Get("/", docsHandler)
}

func handleNotFound(ctx iris.Context) {
	suggestPaths := ctx.FindClosest(3)
	if len(suggestPaths) == 0 {
		ctx.WriteString("not found")
		return
	}

	ctx.HTML("Did you main? <ul>")

	for _, s := range suggestPaths {
		ctx.HTML(`<li><a href="%s">%s</a></li>`, s, s)
	}
	ctx.HTML("</ul>")
}

func docsHandler(ctx iris.Context) {
	ctx.ContentType("text/html") // or ctx.HTML(fmt.Sprintf(...))
	ctx.Writef(`<!DOCTYPE html>
	<html>
		<head>
			<style>
				th, td {
					border: 1px solid black;
					padding: 15px;
					text-align: left;
				}
			</style>
		</head>`)
	defer ctx.Writef("</html>")

	ctx.Writef("<body>")
	defer ctx.Writef("</body>")

	ctx.Writef(`
	<table>
		<tr>
			<th>Method</th>
			<th>Path</th>
			<th>Handler</th>
		</tr>
	`)
	defer ctx.Writef(`</table>`)

	registeredRoutes := ctx.Application().GetRoutesReadOnly()
	for _, r := range registeredRoutes {
		if r.Path() == "/" { // don't list the root, current one.
			continue
		}

		ctx.Writef(`
			<tr>
				<td>%s</td>
				<td>%s%s</td>
				<td>%s</td>
			</tr>
		`, r.Method(), ctx.Host(), r.Path(), r.MainHandlerName())
	}
}
