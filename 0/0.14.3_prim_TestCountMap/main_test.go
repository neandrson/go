package printer

import "testing"

/*func TestPrintHello(t *testing.T) {
	got := PrintHello("Igor")
	expected := "Hello, Igor!"

	if got != expected {
		t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
	}

	for name := range names {
		if name != "Igor" {
			t.Fatalf(`got %q, want %q`, name, "Igor")
		} else {
			break
		}
	}
}*/

func TestPrintHello(t *testing.T) {
	names["Petr"] = "Petr"
	got := PrintHello("Igor")
	expected := "Hello, Igor!"

	if got != expected {
		t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
	}

	for name := range names {
		if name != "Igor" {
			t.Fatalf(`got %q, want %q`, name, "Igor")
		} else {
			break
		}
	}
}

func TestPrintHelloIvan(t *testing.T) {
	got := PrintHello("Ivan")
	expected := "Hello, Ivan!"

	if got != expected {
		t.Fatalf(`PrintHello("Ivan") = %q, want %q`, got, expected)
	}
}

func TestSum(t *testing.T) {
	// набор тестов
	cases := []struct {
		// имя теста
		name string
		// значения на вход тестируемой функции
		values []int
		// желаемый результат
		want int
	}{
		// тестовые данные № 1
		{
			name:   "positive values",
			values: []int{1, 2, 3},
			want:   6,
		},
		// тестовые данные № 2
		{
			name:   "mixed values",
			values: []int{-1, 2, -3},
			want:   -2,
		},
	}
	// перебор всех тестов
	for _, tc := range cases {
		tc := tc
		// запуск отдельного теста
		t.Run(tc.name, func(t *testing.T) {
			// тестируем функцию Sum
			got := Sum(tc.values...)
			// проверим полученное значение
			if got != tc.want {
				t.Errorf("Sum(%v) = %v; want %v", tc.values, got, tc.want)
			}
		})
	}
}
