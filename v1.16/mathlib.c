#include "mathlib.h"
#include <math.h>
#include <time.h>
#include <string.h>
#include <stdlib.h>

int add(int a, int b) {
    return a + b;
}

Point create_point(int x, int y) {
    Point p;
    p.x = x;
    p.y = y;
    return p;
}

Rectangle create_rectangle(Point top_left, Point bottom_right) {
    Rectangle rect;
    rect.top_left = top_left;
    rect.bottom_right = bottom_right;
    return rect;
}

int get_rectangle_area(Rectangle rect) {
    int width = rect.bottom_right.x - rect.top_left.x;
    int height = rect.top_left.y - rect.bottom_right.y;
    // Use absolute values to handle any coordinate order
    if (width < 0) width = -width;
    if (height < 0) height = -height;
    return width * height;
}

int get_rectangle_perimeter(Rectangle rect) {
    int width = rect.bottom_right.x - rect.top_left.x;
    int height = rect.top_left.y - rect.bottom_right.y;
    // Use absolute values to handle any coordinate order
    if (width < 0) width = -width;
    if (height < 0) height = -height;
    return 2 * (width + height);
}

Circle create_circle(Point center, double radius) {
    Circle circle;
    circle.center = center;
    circle.radius = radius;
    return circle;
}

double get_circle_area(Circle circle) {
    return M_PI * pow(circle.radius, 2);
}

double get_circle_circumference(Circle circle) {
    return 2 * M_PI * circle.radius;
}

Point get_point_on_circle(Circle circle, double angle_radians) {
    Point point;
    point.x = (int)(circle.center.x + circle.radius * cos(angle_radians));
    point.y = (int)(circle.center.y + circle.radius * sin(angle_radians));
    return point;
}

// Functions that use external C library structs (time.h)
struct tm* get_current_time_struct(void) {
    time_t now = time(NULL);
    return localtime(&now);
}

int get_hour_from_time_struct(struct tm* time_info) {
    return time_info->tm_hour;
}

int get_minute_from_time_struct(struct tm* time_info) {
    return time_info->tm_min;
}

int get_second_from_time_struct(struct tm* time_info) {
    return time_info->tm_sec;
}

char* format_time_struct(struct tm* time_info) {
    char* buffer = (char*)malloc(64);
    if (buffer == NULL) {
        return NULL;
    }
    
    strftime(buffer, 64, "%H:%M:%S", time_info);
    return buffer;
} 