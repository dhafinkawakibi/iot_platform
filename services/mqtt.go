package services

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

// StartMqtt is function for running mqtt
func StartMqtt() {
	dotenvErr := godotenv.Load()

	if dotenvErr != nil {
		fmt.Println(dotenvErr)
		panic("Failed loading dotenv.")
	}

	uri, err := url.Parse(os.Getenv("MQTT_URL"))

	if err != nil {
		log.Fatal(err)
	}

	topic := os.Getenv("MQTT_TOPIC")
	fmt.Println(topic)
	if topic == "" {
		topic = "test"
	}

	go mqttListen(uri, topic)
}

func mqttConnect(clientID string, uri *url.URL) mqtt.Client {
	opts := mqttCreateClientOptions(clientID, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func mqttCreateClientOptions(clientID string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientID)
	return opts
}

func mqttListen(uri *url.URL, topic string) {
	client := mqttConnect("sub", uri)
	client.Subscribe("/api/req/#", 0, func(client mqtt.Client, msg mqtt.Message) {
		mqttPayloadHandler(msg)
	})
}

func mqttPayloadHandler(msg mqtt.Message) {
	topic := strings.Split(msg.Topic(), "/")
	if len(topic) == 4 {
		fmt.Printf("CREATE USER")
		var result map[string]interface{}
		json.Unmarshal([]byte(string(msg.Payload())), &result)
		fmt.Println(result["username"])
	} else if len(topic) == 5 {
		fmt.Printf("DEVICE-INFO")

	} else if len(topic) == 6 {
		fmt.Printf("CRUD DEVICE")

	} else if len(topic) == 7 {
		fmt.Printf("DATA DEVICE")

	}
}
