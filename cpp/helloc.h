#ifndef __HELLOC_H__
#define __HELLOC_H__

#ifdef __cplusplus
extern "C" {
#endif

    typedef void* helloPtr;
    helloPtr helloInit(void);
    void helloFree(helloPtr h);
    void helloPrint(helloPtr h);

#ifdef __cplusplus
}
#endif

#endif
