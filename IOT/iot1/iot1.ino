#include <ESP8266WiFi.h>
#include <ESP8266WebServer.h>

// Define the Open WiFi SSID
const char* apSSID = "CTF_Network";  // Open WiFi SSID

// Set up web server
ESP8266WebServer server(80);

// Pin for the built-in LED (most ESP8266 boards use GPIO2 for this)
const int ledPin = 2;

// Function to handle the /admin endpoint
void handleAdmin() {
    Serial.println("🚀 /admin endpoint hit!");

    // First, check if the user is authorized
    if (!server.hasHeader("Authorization") || server.header("Authorization") != "h4x0r") {
        Serial.println("❌ Unauthorized access attempt!");
        server.send(401, "text/plain", "Unauthorized. Please provide the correct Authorization header.");
        return;
    }

    // After authorization, check if there are any parameters to control the LED
    if (server.hasArg("led")) {
        String ledState = server.arg("led");
        if (ledState == "on") {
            digitalWrite(ledPin, LOW);  // Turn LED ON (LOW because LED is usually active LOW)
            Serial.println("🔦 LED ON");
            server.send(200, "text/plain", "LED turned ON!");
            return;
        } else if (ledState == "off") {
            digitalWrite(ledPin, HIGH);  // Turn LED OFF
            Serial.println("🔦 LED OFF");
            server.send(200, "text/plain", "LED turned OFF!");
            return;
        } else {
            server.send(400, "text/plain", "Invalid LED state! Use 'on' or 'off'.");
            return;
        }
    }

    // If no valid LED parameter is provided, send the flag
    server.send(200, "text/plain", "FLAG{4b6f3a7e2c1d9b5f}");
    Serial.println("✅ FLAG SENT!");
}

void setup() {
    Serial.begin(115200);
    Serial.println("\nStarting ESP8266 CTF in Open AP Mode...");

    // Initialize the LED pin as output
    pinMode(ledPin, OUTPUT);
    digitalWrite(ledPin, HIGH);  // Ensure the LED is off at startup

    // Start the WiFi Access Point
    if (!WiFi.softAP(apSSID)) {
        Serial.println("❌ Failed to start AP! Restarting...");
        ESP.restart();
    }
    Serial.println("✅ Open Access Point Started!");

    // Print the AP's IP Address
    Serial.print("📌 AP IP Address: ");
    Serial.println(WiFi.softAPIP());

    // Basic welcome message for the root endpoint
    server.on("/", []() {
        server.send(200, "text/html", "<h2>Welcome to CTF!</h2><p>Find the hidden flag.</p>");
    });

    // Handle the /admin endpoint
    server.on("/admin", handleAdmin);  // Hidden flag endpoint

    // Start the web server
    server.begin();
    Serial.println("🚀 HTTP server started!");
}

void loop() {
    server.handleClient();  // Handle incoming client requests
}
