package main

// #cgo LDFLAGS: -lm
// #include <stdlib.h>
// #include "mathlib.h"
import "C"
import (
	"fmt"
	"math"
	"unsafe"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Test CGO with Bazel - Structs + External C Library Demo")
	fmt.Println("=======================================================")

	// Test the C add function
	a, b := 10, 5
	result := int(C.add(C.int(a), C.int(b)))
	fmt.Printf("C Function: %d + %d = %d\n", a, b, result)

	// Create C structs
	point1 := C.create_point(C.int(0), C.int(0))
	point2 := C.create_point(C.int(10), C.int(5))

	fmt.Printf("Point 1: (%d, %d)\n", int(point1.x), int(point1.y))
	fmt.Printf("Point 2: (%d, %d)\n", int(point2.x), int(point2.y))

	// Create a rectangle using the points
	rect := C.create_rectangle(point1, point2)

	// Access struct fields
	fmt.Printf("Rectangle: top_left=(%d, %d), bottom_right=(%d, %d)\n",
		int(rect.top_left.x), int(rect.top_left.y),
		int(rect.bottom_right.x), int(rect.bottom_right.y))

	// Calculate area and perimeter
	area := int(C.get_rectangle_area(rect))
	perimeter := int(C.get_rectangle_perimeter(rect))

	fmt.Printf("Rectangle Area: %d\n", area)
	fmt.Printf("Rectangle Perimeter: %d\n", perimeter)

	// Create a circle using math library functions
	center := C.create_point(C.int(5), C.int(5))
	circle := C.create_circle(center, C.double(3.0))

	fmt.Printf("Circle: center=(%d, %d), radius=%.1f\n",
		int(circle.center.x), int(circle.center.y), float64(circle.radius))

	// Calculate circle properties using math library
	circle_area := float64(C.get_circle_area(circle))
	circle_circumference := float64(C.get_circle_circumference(circle))

	fmt.Printf("Circle Area: %.2f\n", circle_area)
	fmt.Printf("Circle Circumference: %.2f\n", circle_circumference)

	// Get points on the circle using trigonometric functions
	point_0 := C.get_point_on_circle(circle, C.double(0.0))           // 0 degrees
	point_90 := C.get_point_on_circle(circle, C.double(math.Pi/2))    // 90 degrees
	point_180 := C.get_point_on_circle(circle, C.double(math.Pi))     // 180 degrees
	point_270 := C.get_point_on_circle(circle, C.double(3*math.Pi/2)) // 270 degrees

	fmt.Printf("Points on circle:\n")
	fmt.Printf("  0°: (%d, %d)\n", int(point_0.x), int(point_0.y))
	fmt.Printf("  90°: (%d, %d)\n", int(point_90.x), int(point_90.y))
	fmt.Printf("  180°: (%d, %d)\n", int(point_180.x), int(point_180.y))
	fmt.Printf("  270°: (%d, %d)\n", int(point_270.x), int(point_270.y))

	// ===== USING EXTERNAL C LIBRARY STRUCT (time.h) =====
	fmt.Printf("\n=== Using struct tm from time.h library ===\n")

	// Get current time using struct tm from time.h
	time_struct := C.get_current_time_struct()

	// Access fields of the external C library struct
	hour := int(C.get_hour_from_time_struct(time_struct))
	minute := int(C.get_minute_from_time_struct(time_struct))
	second := int(C.get_second_from_time_struct(time_struct))

	fmt.Printf("Current time from struct tm: %02d:%02d:%02d\n", hour, minute, second)

	// Format the time struct using C library function
	time_str := C.format_time_struct(time_struct)
	if time_str != nil {
		defer C.free(unsafe.Pointer(time_str))
		fmt.Printf("Formatted time: %s\n", C.GoString(time_str))
	}

	// Access struct tm fields directly (if needed)
	// Note: struct tm is defined in time.h, so we can access it directly
	fmt.Printf("Direct access - Hour: %d, Minute: %d, Second: %d\n",
		int(time_struct.tm_hour), int(time_struct.tm_min), int(time_struct.tm_sec))

	// Generate a UUID
	uuidv6 := uuid.New()
	fmt.Printf("UUIDv6: %s\n", uuidv6.String())

	// Create another circle with different parameters
	center2 := C.create_point(C.int(10), C.int(10))
	circle2 := C.create_circle(center2, C.double(5.0))

	circle2_area := float64(C.get_circle_area(circle2))
	circle2_circumference := float64(C.get_circle_circumference(circle2))

	fmt.Printf("Circle 2: center=(%d, %d), radius=%.1f\n",
		int(circle2.center.x), int(circle2.center.y), float64(circle2.radius))
	fmt.Printf("Circle 2 Area: %.2f, Circumference: %.2f\n", circle2_area, circle2_circumference)

	fmt.Println("\nTest completed successfully!")
}
