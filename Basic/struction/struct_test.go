package struction

import "testing"

func assertCorrectFloat(t *testing.T, got, expected float64) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %f, while got %f", expected, got)
	}

}

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	assertCorrectFloat(t, got, want)
}

func TestRectanglePerimerter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := RectanglePerimeter(rectangle)
	want := 40.0
	assertCorrectFloat(t, got, want)
}

func TestRectangleArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := RectangleArea(rectangle)
	want := 100.0
	assertCorrectFloat(t, got, want)
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		got := Area(6.0, 12.0)
		want := 72.0
		assertCorrectFloat(t, got, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		assertCorrectFloat(t, got, want)
	})

}

// 使用接口形式定义接口方法  实现接口的参数方法的入参和出参
func TestInterfaceCheckArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertCorrectFloat(t, got, want)
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})

}

// 使用表格驱动测试方法
func TestTableDrivenTestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{12, 6}, want: 72.0},
		{shape: Circle{10}, want: 314.1592653589793},
		{shape: Triangle{10, 12}, want: 60},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assertCorrectFloat(t, got, tt.want)
	}
}

func TestTableDrivenTestAreaWithName(t *testing.T) {
	// 增加断言名称 和结构体优化
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{10, 12}, want: 60},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			assertCorrectFloat(t, got, tt.want)
		})
	}

}
