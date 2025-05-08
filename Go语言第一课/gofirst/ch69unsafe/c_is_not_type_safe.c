
int main() {
    int a = 0x12345678;
    unsigned char *p = (unsigned char *)&a;
    printf("0x%x\n", a);
    *p = 0x23;
    *(p + 1) = 0x45;
    *(p+2) = 0x67;
    *(p+3) = 0x8a;
    printf("0x%x\n", a); // 0x8a674523 (注：在小端字节序系统中输出此值)￼
}