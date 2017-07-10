#include <mraa.h>
#include "pwm.h"

#define PWM_PIN      5

mraa_pwm_context       pwmPin;   // guarda o contexto do terminal

void config(void) {
    // 1: Inicia a biblioteca MRAA
    mraa_init();

    // 2: Inicia o terminal D5 para PWM
    pwmPin = mraa_pwm_init(PWM_PIN);

    // 3: Ajusta o período do PWM para 5000us ou 5ms
    mraa_pwm_period_us(pwmPin, 5000);

    // 4: Ativa PWM
    mraa_pwm_enable(pwmPin, 1);
}

void writePWM(float duty) {
    // 5: ajusta o ciclo de trabalho
    mraa_pwm_write(pwmPin, duty);
}

void stopPWM(void) {
    // 6: Para o PWM
    mraa_pwm_enable(pwmPin, 0);
}
