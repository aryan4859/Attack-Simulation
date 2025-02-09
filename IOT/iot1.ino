#include <ESP8266WiFi.h>
#include <ESP8266WebServer.h>
// hiden_access 4b6f3a7e2c1d9b5f
const char* apSSID = "CTF_Network";  // Open WiFi SSID

ESP8266WebServer server(80);

void handleAdmin() {
    Serial.println("🚀 /admin endpoint hit!");

    int numHeaders = server.headers();
    Serial.print("📡 Total Headers: ");
    Serial.println(numHeaders);

    for (int i = 0; i < numHeaders; i++) {
        Serial.print("🔍 Header ");
        Serial.print(server.headerName(i));
        Serial.print(": ");
        Serial.println(server.header(i));
    }

    if (server.hasHeader("Authorization")) {
        String userAgent = server.header("Authorization");
        Serial.print("📡 Received Authorization: ");
        Serial.println(userAgent);

        if (userAgent.equals("h4x0r")) {  // Strict comparison
            server.send(200, "text/plain", "FLAG{redacted}");
            Serial.println("✅ FLAG SENT!");
            return;
        } else {
            Serial.println("❌ Incorrect Authorization!");
        }
    } else {
        Serial.println("⚠️ No User-Agent header found!");
    }

    server.send(401, "text/plain", "Unauthorized. Try harder!");
}


void setup() {
    Serial.begin(115200);
    Serial.println("\nStarting ESP8266 CTF in Open AP Mode...");

    if (!WiFi.softAP(apSSID)) {
        Serial.println("❌ Failed to start AP! Restarting...");
        ESP.restart();
    }
    Serial.println("✅ Open Access Point Started!");

    Serial.print("📌 AP IP Address: ");
    Serial.println(WiFi.softAPIP());  // Print AP's IP

    server.on("/", []() {
        server.send(200, "text/html", "<h2>Welcome to CTF!</h2><p>Find the hidden flag.</p>");
    });

    server.on("/admin", handleAdmin); // Hidden flag endpoint

    server.begin();
    Serial.println("🚀 HTTP server started!");
}

void loop() {
    server.handleClient();
}
