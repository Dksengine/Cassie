package cc_lib

import "testing"

func TestFromSnakeToCamel(t *testing.T) {
	input := "user_test_class_type"
	output := "userTestClassType"

	if FromSnakeToCamel(input) != output {
		t.Errorf("This is not correct %s", input)
	}

}

func TestFromSnakeToPascal(t *testing.T) {
	input := "this_is_another_example"
	output := "ThisIsAnotherExample"

	if FromSnakeToPascal(input) != output {
		t.Errorf("This is not correct %s", input)
	}

}

func TestFromCamelOrPascalToSnake(t *testing.T) {
	input := "CarpeDiemBrother"
	output := "carpe_diem_brother"

	if FromCamelOrPascalToSnake(input) != output {
		t.Errorf("This is not correct %s", input)
	}

}

func TestFromCamelToPascal(t *testing.T) {
	input := "armaVirumqueCano"
	output := "ArmaVirumqueCano"

	if FromCamelToPascal(input) != output {
		t.Errorf("This is not correct %s", input)
	}

	input2 := "come"
	output2 := "Come"

	if FromCamelToPascal(input2) != output2 {
		t.Errorf("This is not correct %s", input2)
	}
}

func TestFromPascalToCamel(t *testing.T) {
	input := "NelMezzoDelCammin"
	output := "nelMezzoDelCammin"

	if FromPascalToCamel(input) != output {
		t.Errorf("This is not correct %s", input)
	}

	input2 := "Test"
	output2 := "test"

	if FromPascalToCamel(input2) != output2 {
		t.Errorf("This is not correct %s", input2)
	}
}
