#ifndef MATHLIB_H
#define MATHLIB_H

#include <time.h>

// Simple math operations
int add(int a, int b);

// Struct for a 2D point
typedef struct {
    int x;
    int y;
} Point;

// Struct for a rectangle
typedef struct {
    Point top_left;
    Point bottom_right;
} Rectangle;

// Struct for a circle with trigonometric calculations
typedef struct {
    Point center;
    double radius;
} Circle;

// Functions to work with structs
Point create_point(int x, int y);
Rectangle create_rectangle(Point top_left, Point bottom_right);
int get_rectangle_area(Rectangle rect);
int get_rectangle_perimeter(Rectangle rect);

// Circle functions that use math library
Circle create_circle(Point center, double radius);
double get_circle_area(Circle circle);
double get_circle_circumference(Circle circle);
Point get_point_on_circle(Circle circle, double angle_radians);

// Functions that use external C library structs (time.h)
struct tm* get_current_time_struct(void);
int get_hour_from_time_struct(struct tm* time_info);
int get_minute_from_time_struct(struct tm* time_info);
int get_second_from_time_struct(struct tm* time_info);
char* format_time_struct(struct tm* time_info);

#endif // MATHLIB_H 