int count = 0;

void setup() {
    Serial.begin(115200);
}

void loop() {
 Serial.print("count: ");
 Serial.println(count);
 count++;
 delay(1000);
}
