#include <mraa.h>
#include "button.h"

#define LED_PIN    5       // Terminal do LED
#define BTN_PIN    7       // Terminal do botão

mraa_gpio_context ledPin;  // guarda o contexto do led
mraa_gpio_context btnPin;  // guarda o contexto do botão

void config(void) {
    mraa_init();    // Inicia a biblioteca MRAA
    ledPin = mraa_gpio_init(LED_PIN);
    btnPin = mraa_gpio_init(BTN_PIN);

    // Define que o terminal do LED é de saída
    mraa_gpio_dir(ledPin, MRAA_GPIO_OUT);

    // Ajusta o estado do LED para iniciar apagado
    mraa_gpio_write(ledPin, 0);
}

int readButton(void) {
    return (mraa_gpio_read(btnPin));
}

void writeLED(int v) {
    mraa_gpio_write(ledPin, v);
}
