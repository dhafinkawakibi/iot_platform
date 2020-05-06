package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dhafinkawakibi/iot_platform/api/controllers"
	"github.com/dhafinkawakibi/iot_platform/api/models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	RunMqtt()
	server.Run(":8080")
}

func RunMqtt() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
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
	fmt.Println(msg.Topic())
	topic := strings.Split(msg.Topic(), "/")
	if len(topic) == 4 {
		fmt.Printf("CREATE USER")
	} else if len(topic) == 5 {
		fmt.Printf("DEVICE-INFO")
	} else if len(topic) == 6 {
		fmt.Printf("CRUD DEVICE")
		fmt.Println(string(msg.Payload()))
		CreateDevice(msg.Payload())
	} else if len(topic) == 7 {
		fmt.Printf("DATA DEVICE")

	}
}

func CreateDevice(data []byte) {
	device := models.Device{}
	json.Unmarshal(data, &device)

	deviceCreated, err := device.SaveDevice(server.DB)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(deviceCreated)
}
