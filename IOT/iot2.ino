#include <ESP8266WiFi.h>
#include <ESP8266WebServer.h>

ESP8266WebServer server(80);

const char* apSSID = "HiddenPathCTF";  // Name of the Wi-Fi AP
const char* apPass = "ctfchallenge";   // Password for the AP

const char* correctUser = "admin";
const char* correctPass = "password123";  // Not the real secret!

// Hidden path flag (in mixed hex+alphabet)
const char* flag = "FLAG{2a5c9e6f4b0d8d2a}";

// Authenticate function (misleading, doesn't work)
void handleAuth() {
    if (server.hasArg("user") && server.hasArg("pass")) {
        String user = server.arg("user");
        String pass = server.arg("pass");

        // Normal login system (doesn't reveal the flag)
        if (user == correctUser && pass == correctPass) {
            server.send(200, "text/plain", "Welcome to the system!");
            return;
        }
    }

    // Custom Header Check for hidden flag access
    if (server.header("X-Hidden-Access") == "Unlock") {
        server.send(200, "text/plain", flag);  // The real flag is here!
        return;
    }

    // Wrong attempt
    server.send(401, "text/plain", "Unauthorized Access");
}

// Debug endpoint for analysis (to mislead)
void handleDebug() {
    String data = "This is a debug page with no flag.\n";
    data += "Try altering the headers or the user-agent.";
    server.send(200, "text/plain", data);
}

void setup() {
    Serial.begin(115200);
    Serial.println("Starting 'The Hidden Path' CTF challenge...");

    // Set up ESP8266 as an Access Point
    WiFi.softAP(apSSID, apPass);

    // Get and print the ESP8266 AP IP Address
    IPAddress apIP = WiFi.softAPIP();
    Serial.print("AP IP Address: ");
    Serial.println(apIP);

    // Handle routes
    server.on("/admin", handleAuth);  // Flag hidden here via custom header
    server.on("/debug", handleDebug); // Debug endpoint (misleading)

    server.begin();
    Serial.println("HTTP server started!");
}

void loop() {
    server.handleClient();
}
