#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#define MAX_STR_SIZE 50

// Basic Struct for our Car object
typedef struct _car {
    int year;
    char make[MAX_STR_SIZE];
    char model[MAX_STR_SIZE];
    // Class like methods for C [ Source: https://quora.com/Can-structs-have-methods ]
    void (*to_string)(struct _car*);
    
} Car;
void print_car(Car* c) {
    printf("Car: %d %s %s", c->year, c->make, c->model);
}
void assign_func(Car* car) {
    car->to_string = print_car;
}
void new(Car* c, int year, char* make, char* model) {
    c->year = year;
    strcpy(c->make, make);
    strcpy(c->model, model);
    assign_func(c);
}

 // This method alone was done with AI assistance as I was unsure of how to work with the pointers in this instance
void trimWhitespace(char* destination, const char* source) {
    const char* start = source;
    const char* end = source + strlen(source) - 1;

    // Find the first non-space character
    while (isspace((unsigned char)*start)) {
        start++;
    }

    // Find the last non-space character
    while (end > start && isspace((unsigned char)*end)) {
        end--;
    }

    // Copy the trimmed part to the destination
    while (start <= end) {
        *destination++ = *start++;
    }

    // Null-terminate the destination
    *destination = '\0';
}

int main() {
    int validation; // Declared for validation

    int year;
    printf("Enter year: ");
    if (scanf("%d", &year) != 1) { 
        printf("Year is invalid!\n");
        return 1;
    }
    getchar();

    char _make[MAX_STR_SIZE];
    printf("Enter Make: ");
    fgets(_make, MAX_STR_SIZE, stdin);
    char make[MAX_STR_SIZE];
    trimWhitespace(make, _make);

    char _model[MAX_STR_SIZE];
    printf("Enter Model: ");
    fgets(_model, MAX_STR_SIZE, stdin);
    char model[MAX_STR_SIZE];
    trimWhitespace(model, _model);

    Car c;
    new(&c, year, make, model);
    c.to_string(&c);

    return 0;
}
