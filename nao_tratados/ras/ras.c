#include <mraa.h>

#define RAS_PIN       0        // Terminal do potenciometro A0

mraa_aio_context      rasPin;  // guarda o contexto do potenciometro

void config(void) {
    mraa_init();
    rasPin = mraa_aio_init(RAS_PIN);
}

float readRAS(void) {
    return mraa_aio_read(rasPin);
}

void removeRAS(void) {
    mraa_aio_close(rasPin);
}
